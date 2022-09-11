package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
	"log"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	//公钥存储生成的签名
	PublicKey []byte
}

func NewWallet() Wallet {
	PrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	publicKeyS := PrivateKey.PublicKey
	PublicKey := append(publicKeyS.X.Bytes(), publicKeyS.Y.Bytes()...)
	return Wallet{PrivateKey: PrivateKey, PublicKey: PublicKey}
}

// 生成地址
func (wallet *Wallet) NewAddress() string {
	publicKey := wallet.PublicKey
	//sha256
	hash := sha256.Sum256(publicKey)
	//ripemd160
	ripemd160er := ripemd160.New()
	_, err := ripemd160er.Write(hash[:])
	if err != nil {
		log.Panic(err)
	}

	ripemd160Hash := ripemd160er.Sum(nil)
	//	拼接version
	version := byte(00)
	payload := append([]byte{version}, ripemd160Hash...)
	//	copy
	var payloadCopy []byte
	copy(payloadCopy, payload)
	//	计算checkSum
	checkSum1 := sha256.Sum256(payloadCopy)
	checkSum2 := sha256.Sum256(checkSum1[:])
	checkSum := checkSum2[:4]

	//最终的25bytes数据
	payload = append(payload, checkSum...)
	//	address
	address := base58.Encode(payload)
	return address
}
