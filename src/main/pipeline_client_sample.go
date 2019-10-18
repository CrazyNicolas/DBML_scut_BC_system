package main

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

/**
在这个程序里面我们就模拟几个client以随机的时间间隔发送请求消息
*/
type Args struct {
	Info int
	Addr string
}

type Reply struct {
	Response bool
}

func Run_client(wg *sync.WaitGroup, port int) {
	defer wg.Done()
	//生成一个随机数种子以确保后面每个客户端的随机序列是不一样的
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	addr := "tcp@127.0.0.1:9090"
	d := client.NewPeer2PeerDiscovery(addr, "")
	xcli := client.NewXClient("Serv", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xcli.Close()
	//目前让给客户端按照随机的时间间隔无限发送
	for {

		//构建请求参数
		args := &Args{r.Intn(5000), strconv.Itoa(port)}
		reply := &Reply{}

		//构建call
		_, _ = xcli.Go(context.Background(), "Relay", args, reply, nil)
		//睡眠一个随机时间再发下一个请求
		time.Sleep(time.Duration(r.Int31n(5)) * time.Second)

	}

}
func main() {
	var wg sync.WaitGroup
	client_number := 5 //模拟5个client
	wg.Add(client_number)
	for i := 0; i < 5; i++ {
		go Run_client(&wg, 9998+i)
	}
	wg.Wait()
	fmt.Println("测试过程结束")

}
