package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	"os"
)

func main1() {
	generateEccKey()
}

// 生成椭圆曲线秘钥对
func generateEccKey2() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	//	使用x509序列化
	derText, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	//	要组织一个pem.Block(base64编码)
	block := pem.Block{
		Type:  "ecc private key",
		Bytes: derText,
	}
	//	创建文件
	privateKeyFile, err := os.Create("eccprivate.pem")
	if err != nil {
		panic(err)
	}
	defer func(privateKeyFile *os.File) {
		err := privateKeyFile.Close()
		if err != nil {
			panic(err)
		}
	}(privateKeyFile)
	//	使用pem编码将数据写入文件
	err = pem.Encode(privateKeyFile, &block)
	if err != nil {
		panic(err)
	}

	//	__________________________________________________________

	//获取公钥
	publicKey := privateKey.PublicKey
	//	通过x509标准将得到的ras公钥序列化为ASN.1 的 DER编码字符串
	derText, err = x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	//	要组织一个pem.Block(base64编码)
	block = pem.Block{
		Type:  "ecc public key",
		Bytes: derText,
	}
	//	创建文件
	publicKeyFile, err := os.Create("eccpublic.pem")
	if err != nil {
		panic(err)
	}
	defer func(publicKeyFile *os.File) {
		err := publicKeyFile.Close()
		if err != nil {
			panic(err)
		}
	}(publicKeyFile)
	//	使用pem编码将数据写入文件
	err = pem.Encode(publicKeyFile, &block)
	if err != nil {
		panic(err)
	}

}

// 创建签名
func createEccSign2(src []byte, fileName string) []byte {
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
	signature := append(r.Bytes(), s.Bytes()...)
	return signature
}

// 验证签名
func verifyEccSign2(src []byte, fileName string, signature []byte) bool {
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
	r.SetBytes(signature[:len(signature)/2])
	s.SetBytes(signature[len(signature)/2:])
	return ecdsa.Verify(publicKey, sum256[:], &r, &s)
}
