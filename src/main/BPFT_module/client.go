package main

import (
	"crypto/rsa"
	"crypto/sha256"
	"golang.org/x/net/html/atom"
	"time"
)

/**
 client 有以下几个流程：
	1.首先调用Request() 向 primary发起请求
	2.统计收到的Reply() 的return value 并加和其个数
 还需要有针对各种方法建立一些结构体 相应的要有args
*/

type Request_Args struct {
	Operation string
	Timestamp int64
	//客户端标识
	Publickey *rsa.PublicKey
	Msg       string
	digest    [32]byte
}

type Request_Reply struct {
}

func Request(operation string, msg string) {
	args := &Request_Args{operation,
		time.Now().Unix(),
		GetPublicKey("public.pem"),
		msg,
		sha256.Sum256([]byte(msg))}

}
func main() {
	//首先获取注册表找到primary地址
	d := GetRegisterDir()
	//找到primary
	//构建一个
}
