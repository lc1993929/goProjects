package main

type BlockChain struct {
	Blocks []Block
}

// 添加区块
func (chain *BlockChain) AddBlock(data string) {
	block := NewBlock(data, chain.Blocks[len(chain.Blocks)-1].Hash)
	chain.Blocks = append(chain.Blocks, block)

}

func NewBlockChain() BlockChain {
	return BlockChain{
		Blocks: []Block{GenesisBlock()},
	}
}
