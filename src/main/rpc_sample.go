package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

type Transanction struct {
	Tokens    int64  `json:"t"`
	Address   string `json:"add"`
	Timestamp int64  `json:"tt"`
}

func (t *Transanction) Save(req Transanction, res *bool) error {
	*res = true
	conn, err := redis.Dial("tcp", ":6379")
	defer conn.Close()
	if err != nil {
		fmt.Println("redis conn err: ", err)
	}
	_, err = conn.Do("hmset", req.Address, "tokens", req.Tokens, "address", req.Address)
	if err != nil {
		fmt.Println("redis save err :", err)
		*res = false
	}
	return nil
}

//func (t *Transanction) Change(req *Transanction , res *Transanction){
//	req.address += "Changed!!!"
//	res = req
//}

func Create(tokens int64, address string) *Transanction {
	transaction := &Transanction{tokens, address, time.Now().Unix()}
	return transaction
}

func Test(t int64, port string, target string) {
	transaction := Create(t, port)
	rpc.Register(new(Transanction))
	listener, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		fmt.Println("Listen err: ", err)
		return
	}

	fmt.Println("start to connect")

	//监听函数并注册方法
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Accept err: ", err)
				continue
			}
			go func(c net.Conn) {
				fmt.Println("a New Connection comimg in: ", c)
				jsonrpc.ServeConn(c)
			}(conn)
		}
	}()

	for {
		//调用Call方法对别人进行调用
		conn, err := jsonrpc.Dial("tcp", target)
		if err != nil {
			fmt.Println("Dial err: ", err)
		}
		var res bool
		err = conn.Call("Transanction.Save", transaction, &res)
		if err != nil {
			fmt.Println("Call err: ", err)
		}
		fmt.Println("response from "+target+": ", res)
		time.Sleep(time.Second * 1)
	}

}

func main() {
	go Test(100, "10000", "127.0.0.1:10001")
	go Test(200, "10001", "127.0.0.1:10000")

}
