package main

import (
	"context"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
	"golang.org/x/net/html/atom"
	"time"
)

/**
 client 有以下几个流程：
	1.首先调用Request() 向 primary发起请求
	2.统计收到的Reply() 的return value 并加和其个数
 还需要有针对各种方法建立一些结构体 相应的要有args
*/

type Client struct {
}

type Request_Args struct {
	Operation string
	Timestamp int64
	//客户端标识
	Publickey *rsa.PublicKey
}

type Request_Reply struct {
}

func Request(operation string) {
	args := &Request_Args{operation,
		time.Now().Unix(),
		GetPublicKey("public.pem"),
	}
	// TODO 这里应该首先获得主节点的地址，并调用它的那个接收函数

}

/**
4. 接受从其他节点发来的Commit（）参数（即为远程服务）
*/
func (t *Client) Get_Reply(ctx context.Context, args *Reply_Args, reply *Reply_Reply) error {
	// TODO 这里面写处理Reply()的逻辑
	return nil
}

func main() {
}
