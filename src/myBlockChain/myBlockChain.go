package main

import (
	"fmt"
)

func main() {
	//block := NewBlock("给我50个币", []byte{})
	//block.toString()
	chain := NewBlockChain()
	chain.AddBlock("第二个块")
	for i, block := range chain.Blocks {
		fmt.Println("-----------------------------------------------------")
		fmt.Println("区块高度：", i)
		block.toString()
	}

}
