package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	//block := NewBlock("给我50个币", []byte{})
	//block.toString()
	chain := NewBlockChain()
	chain.AddBlock("第二个块")
	stack := NewStack()
	err := chain.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		//初始化第一个节点
		lastBlock := bucket.Get(chain.tail)
		block := Block{}
		block.fromByte(lastBlock)
		//入栈
		stack.Push(block)
		//循环查找上 一 区块
		for len(block.PreHash) > 0 {
			preBlock := bucket.Get(block.PreHash)
			block.fromByte(preBlock)
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

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

}
