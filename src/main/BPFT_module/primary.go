package main

import (
	"context"
)

/**
  主节点primary 有以下几个流程需要相互配合
	1.从client接收 Pre-prepare() 的参数
	2.调用 Pre-prepare() 去触发 replica的 Prepare()
	3.接收来自其他节点的参数给Commit()
	4.调用Commit()去触发其他节点的Reply()
	5.收到其他节点的参数给Reply()
	6.调用Reply()返回结果给Client
	7.Checkpoint()用来检查一个其他节点发来的请求

  还需要有针对各种方法建立一些结构体 相应的要有args
  还需要一些全局变量来管理request 管理 view视图

*/

type Primary struct {
	serialNumber int
	viewNumber   int
}

/**
1. 接受从client来的Request请求的参数（即为远程服务）
*/
func (t *Primary) Get_Request(ctx context.Context, args Request_Args, reply *Request_Reply) error {
	// TODO 对客户端的请求进行校验
	// 请求消息签名是否正确
	// 若校验成功
	Pre_prepare(args.Msg, args.digest)
	return nil
}

type Pre_prepare_Args struct {
	viewNumber   int
	digest       [32]byte
	Msg          string
	serialNumber int
}

type Pre_prepare_Reply struct {
}

func Pre_prepare(message string, digest [32]byte) {
	// TODO (1) 分配一个编号(这里面编号的逻辑并没有完成)
	args := Pre_prepare_Args{
		viewNumber:   0,
		digest:       digest,
		Msg:          message,
		serialNumber: 0,
	}

	// TODO 这里要广播
}

/**
2. 接受从其他Replica发来的prepare（）参数（即为远程服务）
*/
func (t *Primary) Primary_Get_Prepare(ctx context.Context, args *Prepare_Args, reply *Prepare_Reply) error {
	// TODO 这里面写处理Prepare()的逻辑，如果正确的话执行commit()
	return nil
}

/*
	3. 这里面应该有一个Commit方法，但是这个方法是可以通用的，不写。
*/

/**
4. 接受从其他节点发来的Commit（）参数（即为远程服务）
*/
func (t *Primary) Primary_Get_Commit(ctx context.Context, args *Commit_Args, reply *Commit_Reply) error {
	// TODO 这里面写处理Commit()的逻辑，如果正确的话执行Reply()
	return nil
}

/**
5. 这里面也应该有一个Reply方法，但是这个Reply方法可以通用，不写。
*/

/**
6. 所有节点通用的CheckPoint（）方法及结构体
*/

/**
7. 接受从其他节点传来的CheckPoint（）参数（即为远程服务）
*/
func (t *Primary) Get_CheckPoint(ctx context.Context, args *CheckPoint_Args, reply *CheckPoint_Reply) {

}

func main() {

}
