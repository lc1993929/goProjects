package main

import (
	"encoding/hex"
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
const initAddress = "chang"

// AddBlock 添加区块
func (blockChain *BlockChain) AddBlock(transactions []Transaction) {
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
	var UTXOs []TXOutput

	transactions := blockChain.FindUTXOTransactions(address)
	for _, transaction := range transactions {
		for _, output := range transaction.TXOutputs {
			//TODO 此处需要修改为验证hash
			if output.ScriptPubKey == address {
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

func (blockChain *BlockChain) FindNeedUTXOs(fromAddress string, amount float64) (map[string][]uint64, float64) {
	//key为交易id，value为outputs中的索引下标
	UTXOs := make(map[string][]uint64)
	var total float64

	transactions := blockChain.FindUTXOTransactions(fromAddress)
	for _, transaction := range transactions {
		for i, output := range transaction.TXOutputs {
			//TODO 此处需要修改为验证hash
			if output.ScriptPubKey == fromAddress {
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

func (blockChain *BlockChain) FindUTXOTransactions(address string) []Transaction {
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
				//TODO 此处需要修改为验证hash
				if output.ScriptPubKey == address {
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
					//TODO	此处需要修改为验签
					if input.Sig == address {
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
