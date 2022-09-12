package main

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"reflect"
)

type BlockChain struct {
	db *bolt.DB
	//尾部区块指针
	tail []byte
}

const dbPath = "myBlockChain.db"
const bucketName = "blockBucket"
const lastHashKey = "lastHashKey"

// AddBlock 添加区块
func (blockChain *BlockChain) AddBlock(transactions []Transaction) {
	// 验签
	for _, transaction := range transactions {
		if !blockChain.VerifyTransaction(&transaction) {
			fmt.Println("验签失败")
			return
		}
	}

	block := NewBlock(transactions, blockChain.tail)

	err := blockChain.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		err := bucket.Put(block.Hash, Serialize(block))
		if err != nil {
			log.Panic(err)
		}
		//更新尾区块
		err = bucket.Put([]byte(lastHashKey), block.Hash)
		if err != nil {
			log.Panic(err)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	//更新尾区块
	blockChain.tail = block.Hash
	fmt.Printf("新增区块成功！hash:%s,nonce:%v \r\n", hex.EncodeToString(block.Hash), block.Nonce)
}

func NewBlockChain(address string) BlockChain {
	//获取数据库db
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	var lastHash []byte

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
			if err != nil {
				log.Panic(err)
			}
			genesisBlock := GenesisBlock(address)
			err = bucket.Put(genesisBlock.Hash, Serialize(genesisBlock))
			if err != nil {
				log.Panic(err)
			}
			err = bucket.Put([]byte(lastHashKey), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte(lastHashKey))
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return BlockChain{
		db: db, tail: lastHash,
	}
}

func (blockChain *BlockChain) FindUTXOs(address string) []TXOutput {
	if !IsValidAddress(address) {
		log.Panic("地址错误，请检查交易地址:", address)
	}

	var UTXOs []TXOutput
	//根据地址获取公钥hash
	publicKeyHash := AddressToPublicKeyHash(address)
	transactions := blockChain.FindUTXOTransactions(publicKeyHash)
	for _, transaction := range transactions {
		for _, output := range transaction.TXOutputs {
			if bytes.Equal(output.PublicKeyHash, publicKeyHash) {
				UTXOs = append(UTXOs, output)
			}
		}
	}

	return UTXOs

}

func (blockChain *BlockChain) toString() {
	stack := NewStack()
	iterator := blockChain.NewIterator()
	for block, have := iterator.Pre(); have; block, have = iterator.Pre() {
		stack.Push(block)
	}
	//出栈打印
	height := 0
	value := stack.Pop()
	for value != nil {
		b := value.(Block)
		fmt.Println("--------------------------------------------------------------")
		fmt.Println("区块高度：", height)
		b.toString()
		value = stack.Pop()
		height++
	}
}

func (blockChain *BlockChain) toStringR() {
	iterator := blockChain.NewIterator()
	for block, have := iterator.Pre(); have; block, have = iterator.Pre() {
		fmt.Println("--------------------------------------------------------------")
		block.toString()
	}
}

func (blockChain *BlockChain) NewIterator() BlockChainIterator {
	return BlockChainIterator{BlockChain: blockChain, Block: Block{}, db: blockChain.db}
}

func (blockChain *BlockChain) FindNeedUTXOs(fromPublicKeyHash []byte, amount float64) (map[string][]uint64, float64) {
	//key为交易id，value为outputs中的索引下标
	UTXOs := make(map[string][]uint64)
	var total float64

	transactions := blockChain.FindUTXOTransactions(fromPublicKeyHash)
	for _, transaction := range transactions {
		for i, output := range transaction.TXOutputs {
			// 验证提供的公钥hash是否与output原本锁定的公钥hash是否相等
			if bytes.Equal(fromPublicKeyHash, output.PublicKeyHash) {
				//	找到自己可用的最少的utxos
				total += output.Value
				indexArray := UTXOs[string(transaction.TXID)]
				indexArray = append(indexArray, uint64(i))
				UTXOs[string(transaction.TXID)] = indexArray
				if total >= amount {
					return UTXOs, total
				}
			}
		}
	}

	fmt.Printf("不满足转账金额，当前总额：%f，目标金额：%f \r\n", total, amount)
	return UTXOs, total
}

type BlockChainIterator struct {
	BlockChain *BlockChain
	Block      Block
	db         *bolt.DB
}

func (it *BlockChainIterator) Pre() (Block, bool) {
	var block Block
	var hash []byte
	if reflect.DeepEqual(it.Block, Block{}) {
		//	获取尾区块
		hash = it.BlockChain.tail
	} else {
		hash = it.Block.PreHash
	}
	if hash == nil {
		return Block{}, false
	}
	err := it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		preBlockBytes := bucket.Get(hash)
		block = Deserialize(preBlockBytes)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	it.Block = block
	return block, true
}

/*
FindUTXOTransactions
根据给到的公钥hash去比较区块中的output中记录的公钥hash，如果匹配相等，那就说明是给的公钥hash所关联的output
*/
func (blockChain *BlockChain) FindUTXOTransactions(myPublicKeyHash []byte) []Transaction {
	var transactions []Transaction
	//记录已经支出的记录。key为交易id，value为outputs中的索引下标
	spentOutputs := make(map[string][]int64)
	iterator := blockChain.NewIterator()
	for {
		block, have := iterator.Pre()
		if !have {
			break
		}
		for _, transaction := range block.Transactions {
			//获取output
		OUTPUT:
			for i, output := range transaction.TXOutputs {
				if bytes.Equal(output.PublicKeyHash, myPublicKeyHash) {
					//验证是否在spentOutPuts中，如果已存在，说明已被使用，直接跳过
					values, have := spentOutputs[string(transaction.TXID)]
					if have {
						for _, value := range values {
							if int64(i) == value {
								continue OUTPUT
							}
						}
					}
					transactions = append(transactions, transaction)
				}
			}
			//如果当前交易是挖矿交易，那不做遍历，直接跳过
			if !transaction.isCoinbase() {
				//	获取input
				for _, input := range transaction.TXInputs {
					//如果input中的公钥hash等于要查找的公钥hash，那就说明此input所对应的output已经使用过了，在后续的查找中需要跳过。此处使用交易的txId和output的下标索引就可以定位到一个指定的output
					if bytes.Equal(HashPublicKey(input.PublicKey), myPublicKeyHash) {
						indexs := spentOutputs[string(input.TXID)]
						indexs = append(indexs, input.Index)
						spentOutputs[string(input.TXID)] = indexs
					}
				}
			}

		}

	}
	return transactions

}

func (blockChain *BlockChain) FindTransactionByTXId(txid []byte) (Transaction, error) {
	iterator := blockChain.NewIterator()
	for block, have := iterator.Pre(); have; block, have = iterator.Pre() {
		for _, transaction := range block.Transactions {
			if bytes.Equal(transaction.TXID, txid) {
				return transaction, nil
			}
		}
	}
	fmt.Println("未找到指定交易.txid:" + string(txid))
	return Transaction{}, errors.New("未找到指定交易.txid:" + string(txid))
}

func (blockChain *BlockChain) SignTransaction(transaction *Transaction, privateKey *ecdsa.PrivateKey) {
	//签名
	prevTXs := make(map[string]Transaction)
	//找到对应的交易
	for _, input := range transaction.TXInputs {
		if !reflect.DeepEqual(prevTXs[string(input.TXID)], Transaction{}) {
			continue
		}
		targetTransaction, err := blockChain.FindTransactionByTXId(input.TXID)
		if err != nil {
			log.Panic(err)
		}

		prevTXs[string(input.TXID)] = targetTransaction
	}

	transaction.Sign(*privateKey, prevTXs)
}

func (blockChain *BlockChain) VerifyTransaction(transaction *Transaction) bool {
	if transaction.isCoinbase() {
		return true
	}
	//签名
	prevTXs := make(map[string]Transaction)
	//找到对应的交易
	for _, input := range transaction.TXInputs {
		if !reflect.DeepEqual(prevTXs[string(input.TXID)], Transaction{}) {
			continue
		}
		targetTransaction, err := blockChain.FindTransactionByTXId(input.TXID)
		if err != nil {
			log.Panic(err)
		}

		prevTXs[string(input.TXID)] = targetTransaction
	}

	return transaction.Verify(prevTXs)
}
