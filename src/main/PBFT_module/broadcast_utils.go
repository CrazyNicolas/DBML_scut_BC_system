package PBFT_module

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
	"os"
	"strconv"
	"strings"
)

/**
定义一些方法：
	1.获取所有节点的地址 GetRegisterDir()  ---IO
	2.Broadcast2Others() --根据参数动态调用 rpcx.broadcast()
	3.修改注册表 Change_RegisterDir() ----让新主节点的key变成primary 让旧主节点的key变为replica
*/
//func GetRegisterDir() [][][]string {
//
//}

/**
@author：江声
注册文件放在configure/register,dat
这个函数通过读取注册文件，返回一个map，键为节点号，值为节点ip地址及端口
e.g 1 127.0.0.1:9999
*/
func GetRegisterDir() map[int]string {
	regFile, err := os.Open("src/main/configure/register.dat")
	if err != nil {
		fmt.Println("打开注册中心文件失败", err)
	}
	defer regFile.Close()

	data, err := ReadFileLine("src/main/configure/register.dat")
	if err != nil {
		fmt.Println("读取注册中心文件失败", err)
	}

	//addrNum := len(data)
	addrs := make(map[int]string)

	for _, l := range data {
		metainfo := strings.Split(l[:len(l)-1], " ")
		repNum, _ := strconv.Atoi(metainfo[0])
		repAddr := metainfo[1]
		addrs[repNum] = repAddr
	}

	return addrs
}

/**
@author: 江声
TODO 里面有许多细节有待确定，比如是否需要广播给自己？
	等孙铭泽确定好服务名再修改
*/
func (rep *Replica) Broadcast2Others(servicePath, serviceMethod string, m interface{}) {
	addrs := GetRegisterDir()
	var serverMap []*client.KVPair
	for _, v := range addrs {
		addr := "tcp@" + v
		serverMap = append(serverMap, &client.KVPair{Key: addr})
	}
	d := client.NewMultipleServersDiscovery(serverMap)
	xclient := client.NewXClient(servicePath, client.Failover, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()

	switch serviceMethod {
	case "preprepare":
		msg := m.(Prepreprare_Msg)
		reply := new(interface{})
		//TODO 这里主节点不用给自己广播
		err := xclient.Broadcast(context.Background(), serviceMethod, msg, reply)
		if err != nil {
			fmt.Println("广播preprepare失败", err)
		}
	case "prepare":
		msg := m.(Prepare_Msg)
		reply := new(interface{})
		//TODO 需不需要发给自己？
		err := xclient.Broadcast(context.Background(), serviceMethod, msg, reply)
		if err != nil {
			fmt.Println("广播prepare失败", err)
		}
	case "commit":
		msg := m.(Commit_Msg)
		reply := new(interface{})
		//TODO 应该不用给自己广播吧？
		err := xclient.Broadcast(context.Background(), serviceMethod, msg, reply)
		if err != nil {
			fmt.Println("广播commit失败", err)
		}
	default:
	}
}

/**
@author:江声
TODO 在这里写注册服务的过程
*/
func CreateServer(serviceName, addr, meta string) {
	s := server.NewServer()
	s.RegisterName(serviceName, new(Replica), meta)
	s.Serve("tcp", addr)
}
