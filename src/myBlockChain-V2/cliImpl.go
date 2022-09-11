package main

import "fmt"

func (cli *CLI) GetBalance(address string) {
	utxos := cli.blockChain.FindUTXOs(address)
	total := 0.0
	for _, utxo := range utxos {
		total += utxo.Value
	}
	fmt.Printf("%s	的余额为：%f", address, total)
}

// 转账交易
func (cli *CLI) send(from string, to string, amount float64, miner string, data string) {
	//创建挖矿交易
	coinbaseTx := NewCoinbaseTx(miner, data)
	//	创建普通交易
	normalTx := NewNormalTx(from, to, amount, *cli.blockChain)
	if normalTx == nil {
		return
	}
	//	添加到区块
	cli.blockChain.AddBlock([]Transaction{coinbaseTx, *normalTx})
	fmt.Println("转账成功！")
}

// 挖矿处理
func (cli *CLI) mining(miner string, data string) {
	//创建挖矿交易
	coinbaseTx := NewCoinbaseTx(miner, data)
	//	添加到区块
	cli.blockChain.AddBlock([]Transaction{coinbaseTx})
	fmt.Println("挖矿成功！")
}

// NewWallet 创建钱包
func (cli *CLI) NewWallet() {
	//wallet := NewWallet()
	//address := wallet.NewAddress()
	//fmt.Printf("私钥：%v \r\n", wallet.PrivateKey)
	//fmt.Printf("公钥：%v \r\n", wallet.PublicKey)
	//fmt.Printf("地址：%s \r\n", address)
	wallets := NewWallets()
	address := wallets.CreateWallet()
	fmt.Printf("地址：%s \r\n", address)
}

// ListWalletAddresses 创建钱包
func (cli *CLI) ListWalletAddresses() {
	wallets := NewWallets()
	addresses := wallets.listAllAddress()
	for _, address := range addresses {
		fmt.Printf("地址：%s \r\n", address)
	}
}
