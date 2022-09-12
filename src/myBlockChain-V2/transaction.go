package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"log"
	"math/big"
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
	//	引用的交易ID，注意此交易id是上一区块生成output时的交易id
	TXID []byte
	//	引用的output的索引值
	Index int64
	//	解锁脚本
	Signature []byte
	//钱包中的公钥（是公钥，不是公钥hash），注意是对应output中收款方的公钥的hash
	PublicKey []byte
}

// TXOutput 交易输出
type TXOutput struct {
	//	转账金额
	Value float64
	//	锁定脚本，储存公钥生成的hash，注意是收款方的公钥hash
	PublicKeyHash []byte
}

// NewOutput 生成新的OUTPUT
func NewOutput(value float64, address string) TXOutput {
	output := TXOutput{Value: value}
	output.Lock(address)
	return output
}

// Lock 生成锁定脚本（OUTPUT）。根据收款方地址，反向计算收款方的公钥hash
func (output *TXOutput) Lock(address string) {
	publicKeyHash := AddressToPublicKeyHash(address)
	output.PublicKeyHash = publicKeyHash
}

// AddressToPublicKeyHash address转公钥hash
// 收款地址的计算流程见(http://t.zoukankan.com/kumata-p-10477369.html)，也可以见 wallet.NewAddress()
func AddressToPublicKeyHash(address string) []byte {
	//	base58解码
	addressByte := base58.Decode(address)
	//	去除version和checksum
	publicKeyHash := addressByte[1 : len(addressByte)-4]
	return publicKeyHash
}

// 计算交易ID
func (tx *Transaction) calcHash() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TXID = hash[:]
}

/*
挖矿交易：
	挖矿时input的生成不重要，因为都是系统发放，所以可随意
	生成output相当于生成了锁定脚本。通过收款方的地址反向推断出收款方的公钥hash，保存在output中
*/

// NewCoinbaseTx 挖矿交易
func NewCoinbaseTx(address string, data string) Transaction {
	if !IsValidAddress(address) {
		log.Panic("地址错误，请检查交易地址:", address)
	}

	var inputs []TXInput
	var outputs []TXOutput

	//挖矿只有一个input,sig现在一般是矿池的名称
	//挖矿交易的公钥hash为矿工自己随意指定，挖矿交易没有签名，也没有引用的output的txid和index
	inputs = append(inputs, TXInput{[]byte{}, -1, nil, []byte(data)})
	outputs = append(outputs, NewOutput(reward, address))

	transaction := Transaction{[]byte{}, inputs, outputs}
	transaction.calcHash()
	return transaction
}

// 判断是否是挖矿交易
func (tx *Transaction) isCoinbase() bool {
	return len(tx.TXInputs) == 1 && len(tx.TXInputs[0].TXID) == 0 && tx.TXInputs[0].Index == -1
}

// NewNormalTx 创建普通交易
func NewNormalTx(fromAddress, toAddress string, amount float64, blockChain BlockChain) *Transaction {
	if !IsValidAddress(fromAddress) {
		log.Panic("地址错误，请检查交易地址:", fromAddress)
	}
	if !IsValidAddress(toAddress) {
		log.Panic("地址错误，请检查交易地址:", toAddress)
	}

	//获取钱包
	wallets := NewWallets()
	wallet := wallets.WalletMap[fromAddress]
	if wallet == nil {
		//未找到发送方地址的钱包数据
		fmt.Println("未找到发送方地址的钱包数据，交易创建失败")
		return nil
	}
	//获取秘钥对
	publicKey := wallet.PublicKey
	privateKey := wallet.PrivateKey
	//计算公钥hash
	publicKeyHash := HashPublicKey(publicKey)

	utxos, utxosAmount := blockChain.FindNeedUTXOs(publicKeyHash, amount)
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
				TXID:      []byte(txId),
				Index:     int64(index),
				Signature: nil,
				PublicKey: publicKey,
			}
			inputs = append(inputs, input)
		}
	}

	output := NewOutput(amount, toAddress)
	outputs = append(outputs, output)
	if utxosAmount > amount {
		//	找零，把多的金额再转给自己一次
		outputs = append(outputs, NewOutput(utxosAmount-amount, fromAddress))
	}

	transaction := Transaction{[]byte{}, inputs, outputs}
	//计算交易的TXID
	transaction.calcHash()
	//计算input中的signature
	blockChain.SignTransaction(&transaction, privateKey)

	return &transaction
}

// Sign 签名，将私钥和与本次交易有关的交易都拿到进行签名
func (tx *Transaction) Sign(privateKey ecdsa.PrivateKey, prevTXs map[string]Transaction) {
	//挖矿交易不需要做input的签名
	if tx.isCoinbase() {
		return
	}
	transactionCopy := tx.copy()
	for i, input := range transactionCopy.TXInputs {
		preTransaction := prevTXs[string(input.TXID)]
		if len(preTransaction.TXID) == 0 {
			log.Panic("引用了无效交易")
		}
		//将引用的output中的公钥hash赋值给新增input中的公钥字段
		transactionCopy.TXInputs[i].PublicKey = preTransaction.TXOutputs[input.Index].PublicKeyHash
		//根据引用output中的公钥hash和当前交易中的所有output进行hash计算
		transactionCopy.calcHash()
		//置为空以免影响后续其它input的hash计算
		transactionCopy.TXInputs[i].PublicKey = nil
		signDataHash := transactionCopy.TXID
		//	开始签名
		r, s, err := ecdsa.Sign(rand.Reader, &privateKey, signDataHash)
		if err != nil {
			log.Panic("err")
		}
		signature := append(r.Bytes(), s.Bytes()...)
		tx.TXInputs[i].Signature = signature
	}

}

// Verify 验证签名
func (tx *Transaction) Verify(prevTXs map[string]Transaction) bool {
	//挖矿交易不需要做input的验签
	if tx.isCoinbase() {
		return true
	}

	transactionCopy := tx.copy()
	for i, input := range tx.TXInputs {
		preTransaction := prevTXs[string(input.TXID)]
		if len(preTransaction.TXID) == 0 {
			log.Panic("引用了无效交易")
		}
		//将引用的output中的公钥hash赋值给新增input中的公钥字段
		transactionCopy.TXInputs[i].PublicKey = preTransaction.TXOutputs[input.Index].PublicKeyHash
		//根据引用output中的公钥hash和当前交易中的所有output进行hash计算
		transactionCopy.calcHash()
		//置为空以免影响后续其它input的hash计算
		transactionCopy.TXInputs[i].PublicKey = nil
		//得到根据数据计算出的签名前的hash数据
		signDataHash := transactionCopy.TXID
		//获取需要验证的签名数据
		signature := input.Signature
		var r big.Int
		var s big.Int
		r.SetBytes(signature[:len(signature)/2])
		s.SetBytes(signature[len(signature)/2:])
		//获取公钥
		publicKey := input.PublicKey
		var X big.Int
		var Y big.Int
		X.SetBytes(publicKey[:len(publicKey)/2])
		Y.SetBytes(publicKey[len(publicKey)/2:])
		publicKeyOrigin := ecdsa.PublicKey{Curve: elliptic.P256(), X: &X, Y: &Y}

		//	开始验签
		verifySuccess := ecdsa.Verify(&publicKeyOrigin, signDataHash, &r, &s)
		if !verifySuccess {
			fmt.Println("验签失败")
			return false
		}

	}
	return true
}

func (tx *Transaction) copy() Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	for _, input := range tx.TXInputs {
		inputs = append(inputs, TXInput{input.TXID, input.Index, nil, nil})
	}

	for _, output := range tx.TXOutputs {
		outputs = append(outputs, output)
	}

	return Transaction{tx.TXID, inputs, outputs}
}
