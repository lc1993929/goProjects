package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
)

func main() {
	src := []byte("需要传递的消息")
	/*sign := createRsaSign(src, "private.pem")
	fmt.Println(sign)
	success := verifyRsaSign(src, "public.pem", sign)
	fmt.Println(success)*/

	r, s := createEccSign(src, "eccprivate.pem")
	fmt.Println(r)
	fmt.Println(s)
	success := verifyEccSign(src, "eccpublic.pem", r, s)
	fmt.Println(success)

}

// 创建签名
func createRsaSign(src []byte, fileName string) string {
	//计算hash
	sum256 := sha256.Sum256(src)
	//	获取私钥
	//	读取文件
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	fileInfo, err := file.Stat()
	privateKeyText := make([]byte, fileInfo.Size())
	_, err = file.Read(privateKeyText)
	if err != nil {
		panic(err)
	}
	//	通过pem解析为block
	block, _ := pem.Decode(privateKeyText)
	if block == nil {
		panic(nil)
	}
	//	获取私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//	生成签名
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, sum256[:])
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(sign)
}

// 验证签名
func verifyRsaSign(src []byte, fileName string, sign string) bool {
	//计算hash
	sum256 := sha256.Sum256(src)
	//	获取公钥
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	fileInfo, err := file.Stat()
	publicKeyText := make([]byte, fileInfo.Size())
	_, err = file.Read(publicKeyText)
	if err != nil {
		panic(err)
	}
	//	通过pem解析为block
	block, _ := pem.Decode(publicKeyText)
	if block == nil {
		panic(err)
	}
	//	获取公钥
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	signBytes, err := hex.DecodeString(sign)
	if err != nil {
		panic(err)
	}
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, sum256[:], signBytes)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// 创建签名
func createEccSign(src []byte, fileName string) (rText, sText []byte) {
	//计算hash
	sum256 := sha256.Sum256(src)
	//	获取私钥
	//	读取文件
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	fileInfo, err := file.Stat()
	privateKeyText := make([]byte, fileInfo.Size())
	_, err = file.Read(privateKeyText)
	if err != nil {
		panic(err)
	}
	//	通过pem解析为block
	block, _ := pem.Decode(privateKeyText)
	if block == nil {
		panic(nil)
	}
	//	获取私钥
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//	生成签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, sum256[:])
	if err != nil {
		panic(err)
	}
	//对r,s进行数据格式化
	rText, err = r.MarshalText()
	if err != nil {
		panic(err)
	}
	sText, err = s.MarshalText()
	if err != nil {
		panic(err)
	}
	return rText, sText
}

// 验证签名
func verifyEccSign(src []byte, fileName string, rText, sText []byte) bool {
	//计算hash
	sum256 := sha256.Sum256(src)
	//	获取公钥
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	fileInfo, err := file.Stat()
	publicKeyText := make([]byte, fileInfo.Size())
	_, err = file.Read(publicKeyText)
	if err != nil {
		panic(err)
	}
	//	通过pem解析为block
	block, _ := pem.Decode(publicKeyText)
	if block == nil {
		panic(err)
	}
	//	获取公钥
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	publicKey := publicKeyInterface.(*ecdsa.PublicKey)
	if err != nil {
		panic(err)
	}
	var r, s big.Int
	err = r.UnmarshalText(rText)
	if err != nil {
		panic(err)
	}
	err = s.UnmarshalText(sText)
	if err != nil {
		panic(err)
	}
	return ecdsa.Verify(publicKey, sum256[:], &r, &s)
}
