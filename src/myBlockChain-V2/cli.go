package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type CLI struct {
	blockChain *BlockChain
}

const Usage = `
		mining miner data 					"创建挖矿区块"
		printChain							"print all blockChain data"
		printChainR							"print all blockChain data reverse"
		getBalance address					"get balance of this address"
		send from to amount miner data		"由from转amount给to，由miner挖矿，同时写入data"
		newWallet							"创建一个新钱包(ECC秘钥对)"
		listWallet							"输出所有钱包地址"
`

func (cli *CLI) Run() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(Usage)
		return
	}

	cmd := args[1]
	switch cmd {
	case "mining":
		//	添加区块
		fmt.Println("开始挖矿")
		if len(args) != 4 {
			fmt.Println("参数个数错误")
			fmt.Println(Usage)
			return
		}
		miner := args[2]
		data := args[3]
		cli.mining(miner, data)
	case "printChain":
		//	打印
		fmt.Println("打印区块")
		cli.blockChain.toString()

	case "printChainR":
		//	打印
		fmt.Println("逆向打印区块")
		cli.blockChain.toStringR()

	case "getBalance":
		fmt.Println("获取余额")
		if len(args) != 3 {
			fmt.Println("参数个数错误")
			fmt.Println(Usage)
			return
		}
		address := args[2]
		cli.GetBalance(address)
	case "send":
		fmt.Println("开始转账")
		if len(args) != 7 {
			fmt.Println("参数个数错误")
			fmt.Println(Usage)
			return
		}
		from := args[2]
		to := args[3]
		amount, err := strconv.ParseFloat(args[4], 64)
		if err != nil {
			log.Panic(err)
		}
		miner := args[5]
		data := args[6]

		cli.send(from, to, amount, miner, data)
	case "newWallet":
		fmt.Println("开始创建钱包")
		cli.NewWallet()
	case "listWallet":
		fmt.Println("列举所有钱包地址")
		cli.ListWalletAddresses()
	default:
		fmt.Println("无效命令")
		fmt.Println(Usage)
	}
}
