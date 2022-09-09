package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	secretKey := []byte("secret")
	message := []byte{1, 2, 3}
	sign := generateHMAC(message, secretKey)
	isSame := VerifyHMAC(sign, message, secretKey)
	fmt.Println(isSame)
}

// 生成消息验证码
func generateHMAC(src, key []byte) []byte {
	// 1. 创建一个底层采用sha256算法的 hash.Hash 接口
	myHmac := hmac.New(sha256.New, key)
	// 2. 添加测试数据
	myHmac.Write(src)
	// 3. 计算结果
	result := myHmac.Sum(nil)
	return result
}

// 验证消息验证码
func VerifyHMAC(res, src, key []byte) bool {
	// 1. 创建一个底层采用sha256算法的 hash.Hash 接口
	myHmac := hmac.New(sha256.New, key)
	// 2. 添加测试数据
	myHmac.Write(src)
	// 3. 计算结果
	result := myHmac.Sum(nil)
	// 4. 比较结果
	return hmac.Equal(res, result)
}
