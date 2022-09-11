package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

const reward = 12.5

// Transaction 交易
type Transaction struct {
	//交易id
	TXID []byte
	//	交易输入
	TXInputs []TXInput
	//交易输出
	TXOutputs []TXOutput
}

// TXInput 交易输入
type TXInput struct {
	//	引用的交易ID
	TXID []byte
	//	引用的output的索引值
	Index int64
	//	解锁脚本
	Sig string
}

// TXOutput 交易输出
type TXOutput struct {
	//	转账金额
	Value float64
	//	锁定脚本
	ScriptPubKey string
}

// 计算交易ID
func (transaction *Transaction) calcHash() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(transaction)
	if err != nil {
		log.Panic(err)
	}
	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	transaction.TXID = hash[:]
}

// NewCoinbaseTx 挖矿交易
func NewCoinbaseTx(address string, data string) Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	//挖矿只有一个input,sig现在一般是矿池的名称
	inputs = append(inputs, TXInput{[]byte{}, -1, data})
	outputs = append(outputs, TXOutput{reward, address})

	transaction := Transaction{[]byte{}, inputs, outputs}
	transaction.calcHash()
	return transaction
}

// 判断是否是挖矿交易
func (transaciton *Transaction) isCoinbase() bool {
	return len(transaciton.TXInputs) == 1 && len(transaciton.TXInputs[0].TXID) == 0 && transaciton.TXInputs[0].Index == -1
}

// NewNormalTx 创建普通交易
func NewNormalTx(fromAddress, toAddress string, amount float64, blockChain BlockChain) *Transaction {
	utxos, utxosAmount := blockChain.FindNeedUTXOs(fromAddress, amount)
	if utxosAmount < amount {
		fmt.Println("余额不足交易失败")
		return nil
	}

	var inputs []TXInput
	var outputs []TXOutput
	//将output转为input
	for txId, indexs := range utxos {
		for _, index := range indexs {
			input := TXInput{
				TXID:  []byte(txId),
				Index: int64(index),
				Sig:   fromAddress,
			}
			inputs = append(inputs, input)
		}
	}

	output := TXOutput{ScriptPubKey: toAddress, Value: amount}
	outputs = append(outputs, output)
	if utxosAmount > amount {
		//	找零，把多的金额再转给自己一次
		outputs = append(outputs, TXOutput{ScriptPubKey: fromAddress, Value: utxosAmount - amount})
	}

	transaction := Transaction{[]byte{}, inputs, outputs}
	transaction.calcHash()

	return &transaction
}
