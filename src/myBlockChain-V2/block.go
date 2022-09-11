package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
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
	//	数据
	Data []byte
}

// 创始块
func GenesisBlock() Block {
	return NewBlock("第一个块", []byte{})
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
		data = append(data, block.Data...)*/

	data := bytes.Join([][]byte{
		Uint64ToByte(block.Version),
		block.PreHash,
		block.MerkleRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}, []byte{})

	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

// 可视化输出
func (block *Block) toString() {
	fmt.Println("版本：", block.Version)
	fmt.Println("前区块hash：", hex.EncodeToString(block.PreHash))
	fmt.Println("merkel根：", hex.EncodeToString(block.MerkleRoot))
	fmt.Println("时间戳：", time.Unix(int64(block.TimeStamp), 0))
	fmt.Println("随机数：", block.Nonce)
	fmt.Println("当前区块hash：", hex.EncodeToString(block.Hash))
	fmt.Println("区块数据：", string(block.Data))
}

func (block *Block) toByte() []byte {
	marshal, err := json.Marshal(block)
	if err != nil {
		log.Panic(err)
	}
	return marshal
}

func (block *Block) fromByte(data []byte) {
	err := json.Unmarshal(data, block)
	if err != nil {
		log.Panic(err)
	}
}

// NewBlock 创建区块
func NewBlock(data string, preBlockHash []byte) Block {
	block := Block{
		Version:    00,
		PreHash:    preBlockHash,
		MerkleRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{},
		Data:       []byte(data),
	}

	//block.calcHash()
	proofOfWork := NewProofOfWork(block)
	hash, nonce := proofOfWork.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}

func Uint64ToByte(num uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, num)
	return b
}
