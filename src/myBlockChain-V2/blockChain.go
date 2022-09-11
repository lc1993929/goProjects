package main

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockChain struct {
	db *bolt.DB
	//尾部区块指针
	tail []byte
}

const dbPath = "myBlockChain.db"
const bucketName = "blockBucket"
const lastHashKey = "lastHashKey"

// 添加区块
func (chain *BlockChain) AddBlock(data string) {
	block := NewBlock(data, chain.tail)
	err := chain.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		err := bucket.Put(block.Hash, block.toByte())
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
	chain.tail = block.Hash

}

func NewBlockChain() BlockChain {
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
			genesisBlock := GenesisBlock()
			err = bucket.Put(genesisBlock.Hash, genesisBlock.toByte())
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
