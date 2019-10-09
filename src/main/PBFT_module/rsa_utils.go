package PBFT_module

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
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

func GetPrivateKey(path string) *rsa.PrivateKey {
	file, err := os.Open(path)
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

func GetPublicKey(path string) *rsa.PublicKey {
	file, err := os.Open(path)
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

//这个方法是节点的成员方法，用来根据具体的节点来获取私钥，比GetPrivateKey好用
//TODO 节点的密钥对存储方式有待确定
func (rep *Replica) GetPrivateKey() *rsa.PrivateKey {
	return nil
}

//TODO 获取节点公钥
func (rep *Replica) GetPublicKey() *rsa.PublicKey {
	return nil
}

func DigitalSignature(message interface{}, privatekey *rsa.PrivateKey) []byte {
	//从密钥文件里面把私钥拿到

	//hash := sha256.New()
	//hash.Write(message)
	//digest := hash.Sum(nil)

	//生成message的消息摘要
	digest := Digest(message)
	digital_signature, err := rsa.SignPKCS1v15(rand.Reader, privatekey, crypto.SHA256, digest)
	if err != nil {
		fmt.Println("Signing err: ", err)
		panic(err)
	}
	return digital_signature
}

func Digest(message interface{}) []byte {
	msg, _ := json.Marshal(message)
	hash := sha256.New()
	hash.Write(msg)
	digest := hash.Sum(nil)
	return digest
}

/**
江声
修改了参数定义，siganature是该消息的数字签名，pub为发送者的公钥，message为任意类型的消息
*/
func Verify_ds(signature []byte, pub *rsa.PublicKey, message interface{}) bool {
	hash := sha256.New()
	//把message序列化
	data, _ := json.Marshal(message)
	hash.Write(data)
	//生成摘要待查验
	digest := hash.Sum(nil)
	//将摘要和数字签名公钥解析后的对比
	err := rsa.VerifyPKCS1v15(pub, crypto.SHA256, digest, signature)
	return err == nil
}
