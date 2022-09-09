package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func main() {
	test := []byte{1, 2, 3}
	//test = paddingLastGroup(test, 8)
	//fmt.Println(test)
	//test = unPaddingLastGroup(test)
	//fmt.Println(test)

	test = desEncrypt(test, []byte("12345678"))
	fmt.Println(test)
	test = desDecrypt(test, []byte("12345678"))
	fmt.Println(test)

	test = aesEncrypt(test, []byte("12345678abcdefgh"))
	fmt.Println(test)
	test = aesDecrypt(test, []byte("12345678abcdefgh"))
	fmt.Println(test)
}

// 填充函数
func paddingLastGroup(plainText []byte, blockSize int) []byte {
	//	求出最后一个组中剩余的字节数
	padNum := blockSize - (len(plainText) % blockSize)
	//	创建新的切片
	char := []byte{byte(padNum)}
	newPlain := bytes.Repeat(char, padNum)
	//将填充数组追加在原数组后面
	plainText = append(plainText, newPlain...)
	return plainText
}

// 删除填充
func unPaddingLastGroup(plainText []byte) []byte {
	//获取最后一个字节
	length := len(plainText)
	lastChar := plainText[length-1]
	num := int(lastChar)
	return plainText[:length-num]
}

func desEncrypt(plainText, key []byte) []byte {
	//获取一个des密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//明文填充
	plainText = paddingLastGroup(plainText, block.BlockSize())
	//	创建cbc分组模式
	iv := []byte("12345678")
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(plainText, plainText)
	return plainText
}

func desDecrypt(plainText, key []byte) []byte {
	//获取一个des密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//	创建cbc分组模式
	iv := []byte("12345678")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(plainText, plainText)
	//去除填充
	plainText = unPaddingLastGroup(plainText)

	return plainText
}

func aesEncrypt(plainText, key []byte) []byte {
	//获取一个aes密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//	创建ctr分组模式
	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plainText, plainText)
	return plainText
}

func aesDecrypt(plainText, key []byte) []byte {
	//获取一个aes密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//	创建ctr分组模式
	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plainText, plainText)
	return plainText
}
