package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

/*
golang封装的ecdsa目前只有用私钥加密，公钥做校验，没有解密环节；所以目前可以应用于数字签名
所以不能用作非对称加解密
如果一定要用椭圆曲线加密，可以考虑使用以太坊加密库（github.com/ethereum/go-ethereum/crypto/ecies）
*/

func main() {
	//generateRsaKey(1024)
	//generateEccKey()
	src := []byte{1, 2, 3}
	src = RSAPublicEncrypt(src, "public.pem")
	fmt.Println(src)
	src = RSAPrivateDecrypt(src, "private.pem")
	fmt.Println(src)

}

// 生成秘钥对
func generateRsaKey(keySize int) {
	//获取随机私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		panic(err)
	}
	//	通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	derText := x509.MarshalPKCS1PrivateKey(privateKey)
	//	要组织一个pem.Block(base64编码)
	block := pem.Block{
		Type:  "rsa private key",
		Bytes: derText,
	}
	//	创建文件
	privateKeyFile, err := os.Create("private.pem")
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
	derText = x509.MarshalPKCS1PublicKey(&publicKey)
	//	要组织一个pem.Block(base64编码)
	block = pem.Block{
		Type:  "rsa public key",
		Bytes: derText,
	}
	//	创建文件
	publicKeyFile, err := os.Create("public.pem")
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

// 公钥加密
func RSAPublicEncrypt(src []byte, fileName string) []byte {
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
	//	加密
	result, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, src)
	if err != nil {
		panic(err)
	}
	return result
}

// 私钥解密
func RSAPrivateDecrypt(src []byte, fileName string) []byte {
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
		return nil
	}
	//	获取私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//	解密
	result, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)
	if err != nil {
		panic(err)
	}
	return result
}

// 生成椭圆曲线秘钥对
func generateEccKey() {
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
