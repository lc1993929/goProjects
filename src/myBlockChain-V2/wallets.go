package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"log"
	"os"
)

type Wallets struct {
	//	map[地址]钱包
	WalletMap map[string]*Wallet
}

const walletFile = "wallet.dat"
const initAddress = "1CxcCDwRd59aEtevgb9DPaBfJrabyYrmeE"

func NewWallets() Wallets {
	wallets := Wallets{WalletMap: map[string]*Wallet{}}
	wallets.loadFile()

	return wallets
}

func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := wallet.NewAddress()
	ws.WalletMap[address] = &wallet

	ws.saveToFile()
	return address
}

func (ws *Wallets) saveToFile() {

	var buffer bytes.Buffer

	//注意：此处如果用go1.19.1有bug，p256无法注册成功
	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(*ws)
	if err != nil {
		log.Panic(err)
	}

	err = os.WriteFile(walletFile, buffer.Bytes(), 0600)
	if err != nil {
		log.Panic(err)
	}
}

func (ws *Wallets) loadFile() {
	_, err := os.Stat(walletFile)
	if os.IsNotExist(err) {
		return
	}

	file, err := os.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(file))

	var wallets Wallets

	err = decoder.Decode(&wallets)
	if err != nil {
		log.Panic(err)
	}
	ws.WalletMap = wallets.WalletMap
}

func (ws *Wallets) listAllAddress() []string {
	var addresses []string

	for address := range ws.WalletMap {
		addresses = append(addresses, address)
	}
	return addresses

}
