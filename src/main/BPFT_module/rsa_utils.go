package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func GenerateRsaKeyPair(bits int) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Println("GenerateKey err : ", err)
	}
	x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//新建一个私钥文件
	privateKeyFile, err := os.Create("private.pem")
	if err != nil {
		fmt.Println("CreateFile err: ", err)
	}
	defer privateKeyFile.Close()
	privateKeyBlock := pem.Block{Type: "RSA Private Key", Bytes: x509PrivateKey}
	//将块编码进文件
	pem.Encode(privateKeyFile, &privateKeyBlock)

	publicKey := privateKey.PublicKey
	x509PublicKey := x509.MarshalPKCS1PublicKey(&publicKey)
	//新建一个公钥文件
	publicKeyFile, err := os.Create("public.pem")
	if err != nil {
		fmt.Println("CreateFile err:", err)
	}
	defer publicKeyFile.Close()
	publicKeyBlock := pem.Block{Type: "RSA Public Key", Bytes: x509PublicKey}
	//将公钥的编码写进文件
	pem.Encode(publicKeyFile, &publicKeyBlock)
}

func GetPrivateKey() *rsa.PrivateKey {
	file, err := os.Open("private.pem")
	if err != nil {
		panic(err)
	}
	fileinfo, _ := file.Stat()
	buf := make([]byte, fileinfo.Size())
	file.Read(buf)
	block, _ := pem.Decode(buf)
	privatekey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	return privatekey
}

func GetPublicKey() *rsa.PublicKey {
	file, err := os.Open("public.pem")
	if err != nil {
		panic(err)
	}
	fileinfo, _ := file.Stat()
	buf := make([]byte, fileinfo.Size())
	file.Read(buf)
	block, _ := pem.Decode(buf)
	publickey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	return publickey
}

func DigitalSignature() {

}

func main() {
	//调用生成密钥对
	GenerateRsaKeyPair(2048)
}
