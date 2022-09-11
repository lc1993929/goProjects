package main

import "math/big"

type ProofOfWork struct {
	block  Block
	target big.Int
}

func NewProofOfWork(block Block) ProofOfWork {
	targetStr := "000100000000000000000000000000000000000000000000000000000000000"
	b := big.Int{}
	b.SetString(targetStr, 16)

	proofOfWork := ProofOfWork{block: block, target: b}
	return proofOfWork

}

// 循环计算符合要求的nonce和对应的hash并输出
func (proofOfWork ProofOfWork) Run() (hash []byte, nonce uint64) {
	nonce = uint64(0)
	//先初始化一下temp
	temp := big.Int{}
	proofOfWork.block.calcHash()
	temp.SetBytes(proofOfWork.block.Hash)
	//计算出一个比目标值小的hash再退出循环
	for temp.Cmp(&proofOfWork.target) > 0 {
		proofOfWork.block.Nonce = nonce
		proofOfWork.block.calcHash()
		temp.SetBytes(proofOfWork.block.Hash)
		nonce = nonce + 1
	}
	return proofOfWork.block.Hash, nonce

}
