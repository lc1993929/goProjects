package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	//first([]byte{1, 2, 3})
	//second([]byte{1, 2, 3})
	third("public.pem")

}

func first(src []byte) {
	sum256 := sha256.Sum256(src)
	fmt.Printf("%x", sum256)
	fmt.Println(hex.EncodeToString(sum256[:]))
}

func second(src []byte) {
	hash := md5.New()
	//_, err := io.WriteString(hash, string(src))
	//if err != nil {
	//	panic(err)
	//}
	hash.Write(src)
	sum := hash.Sum(nil)
	fmt.Println(hex.EncodeToString(sum))

}

func third(src string) {
	file, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	hash := sha1.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		panic(err)
	}
	result := hash.Sum(nil)
	fmt.Println(hex.EncodeToString(result))
}
