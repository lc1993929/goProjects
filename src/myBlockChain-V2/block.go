package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Block struct {
	//版本号
	Version uint64
	//	前区块
	PreHash []byte
	//merkel根
	MerkleRoot []byte
	//时间戳
	TimeStamp uint64
	//难度值
	Difficulty uint64
	//随机数
	Nonce uint64
	//	当前区块
	Hash []byte
	//	交易数据
	Transactions []Transaction
}

const myWorkerName = "Chang"

// 创始块
func GenesisBlock(address string) Block {
	transaction := NewCoinbaseTx(address, myWorkerName)
	transactions := []Transaction{transaction}
	return NewBlock(transactions, []byte{})
}

// 计算hash
func (block *Block) calcHash() {
	/*	data := []byte{}
		data = append(data, Uint64ToByte(block.Version)...)
		data = append(data, block.PreHash...)
		data = append(data, block.MerkleRoot...)
		data = append(data, Uint64ToByte(block.TimeStamp)...)
		data = append(data, Uint64ToByte(block.Difficulty)...)
		data = append(data, Uint64ToByte(block.Nonce)...)
		data = append(data, block.Transactions...)*/

	//计算hash时不需要把交易数据纳入计算，交易数据通过merkel影响最终hash
	data := bytes.Join([][]byte{
		Uint64ToByte(block.Version),
		block.PreHash,
		block.MerkleRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		//block.Transactions,
	}, []byte{})

	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

// NewBlock 创建区块
func NewBlock(transactions []Transaction, preBlockHash []byte) Block {
	block := Block{
		Version:      00,
		PreHash:      preBlockHash,
		MerkleRoot:   []byte{},
		TimeStamp:    uint64(time.Now().Unix()),
		Difficulty:   0,
		Nonce:        0,
		Hash:         []byte{},
		Transactions: transactions,
	}

	block.MerkleRoot = block.MakeMerkelRoot()
	proofOfWork := NewProofOfWork(block)
	hash, nonce := proofOfWork.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}

// MakeMerkelRoot 计算merkel根
func (block *Block) MakeMerkelRoot() []byte {
	var result []byte
	//二叉树太麻烦，先直接做拼接了
	for _, transaction := range block.Transactions {
		result = append(result, transaction.TXID...)
	}
	hash := sha256.Sum256(result)
	return hash[:]
}

func (block *Block) JsonSerialize() []byte {
	marshal, err := json.Marshal(block)
	if err != nil {
		log.Panic(err)
	}
	return marshal
}

func Serialize(block Block) []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	data := buffer.Bytes()
	return data
}

func (block *Block) JsonDeserialize(data []byte) {
	err := json.Unmarshal(data, block)
	if err != nil {
		log.Panic(err)
	}
}

// Deserialize 如果写成方法，在循环调用时，不会重新分配内存。会造成逻辑错误
func Deserialize(data []byte) Block {
	block := Block{}
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return block
}

// 可视化输出
func (block *Block) toString() {
	/*fmt.Println("版本：", block.Version)
	fmt.Println("前区块hash：", hex.EncodeToString(block.PreHash))
	fmt.Println("merkel根：", hex.EncodeToString(block.MerkleRoot))
	fmt.Println("时间戳：", time.Unix(int64(block.TimeStamp), 0))
	fmt.Println("随机数：", block.Nonce)
	fmt.Println("当前区块hash：", hex.EncodeToString(block.Hash))
	fmt.Println("区块数据：", any(json.Marshal(block.Transactions)))*/
	marshal, err := json.MarshalIndent(block, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(marshal))
}

func Uint64ToByte(num uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, num)
	return b
}
