package main

import (
	"encoding/json"
	"fmt"
)

type Block struct {
	//版本号
	Version uint64 `json:"version"`
	//	前区块
	PreHash []byte `json:"preHash"`
	//merkel根
	MerkleRoot []byte `json:"merkleRoot"`
	//	交易数据
	Transactions []Transaction
}

type Transaction struct {
	Id   int
	Name string
}

func main() {

	block := Block{Version: 123, PreHash: []byte{1, 2, 3}, MerkleRoot: []byte{1, 2, 3}, Transactions: []Transaction{
		{Id: 123, Name: "test1"}, {Id: 133, Name: "test2"},
	}}
	//marshal, _ := json.Marshal(block)
	marshal, _ := json.MarshalIndent(block, "", "    ")
	fmt.Println(string(marshal))

}
