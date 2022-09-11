package main

func main() {
	blockChain := NewBlockChain(initAddress)
	cli := CLI{blockChain: &blockChain}
	cli.Run()
	//blockChain.AddBlock("第二个块")
	//blockChain.AddBlock("第3个块")
	//blockChain.toString()
}
