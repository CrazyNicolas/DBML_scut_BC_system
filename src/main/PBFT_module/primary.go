package PBFT_module

import (
	"context"
	"crypto/rsa"
	"database/sql"
	"encoding/json"
	"fmt"
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

/**
主节点结构，继承replica
*/
type Primary struct {
	Replica
	n int32 //分配的序号n
}

/**
1. 接受从client来的Request请求的参数（即为远程服务）
*/
func (pri *Primary) Get_Request(ctx context.Context, args Request_Msg, reply interface{}) error {
	//要拿到客户端的公钥，否则无法验证，这里暂时使用Request_Msg消息内部的publickKey属性
	publickKey := args.publicKey

	//主节点的私钥,暂时用private.pem
	privateKey := GetPrivateKey("private.pem")
	if !Verify_ds(args.signature, publickKey, args.request) {
		fmt.Println("request客户端校验不通过")
	} else {
		pri.Pre_prepare(args, privateKey)
	}
	return nil
}

/**
江声
参数说明：
request 主节点收到的request
private 主节点的私钥
*/
func (pri *Primary) Pre_prepare(request Request_Msg, private *rsa.PrivateKey) {
	// TODO 分配一个编号(这里面编号的逻辑并没有完成)
	args := NewPreprepare(pri.n, pri.viewNumber, request.request, private)
	pri.n++ //分配好一个request消息后n要增加
	//TODO rpxc调用，向其他节点广播

}

/**
2. 接受从其他Replica发来的prepare（）参数（即为远程服务）
*/
func (pri *Primary) Primary_Get_Prepare(ctx context.Context, args *Prepare_Msg, reply *interface{}) error {
	// TODO 这里面写处理Prepare()的逻辑，如果正确的话执行commit()
	//签名是否正确
	if !Verify_ds(args.signature, pri.GetPublicKey(), args.prepare) {
		fmt.Println("prepare消息签名校验不通过")
		return nil
	}
	//验证视图编号是否一致
	if pri.viewNumber != args.v {
		fmt.Println("prepare消息的视图编号与主节点的不一致")
		return nil
	}
	//验证水线
	if args.n > pri.H || args.n < pri.h {
		fmt.Println("prepare消息的n不在水线之内")
		return nil
	}

	//现在已经验证通过，先记入日志
	pri.log(args)

	//TODO 判断prepare阶段是否已经完成，若完成则广播commit
	ok, req := pri.checkPrepared(pri.n)

	return nil
}

/**
江声
节点的私有方法，用来判断是否完成prepared阶段,传的参数n代表需要检查的request请求编号，即主节点给他分配的n
返回值为节点的v和给定的n下的request m
*/
func (rep *Replica) checkPrepared(n int32) (bool, *request) {
	db, err := sql.Open("mysql", db_address)
	if err != nil {
		fmt.Println("连接到数据库失败", err)
	}
	var (
		flag1 bool //是否记录了request m，且记录了m对应的pre-prepare消息，而且对应的n,v一致
		flag2 bool //是否收到了来自2f个节点的prepare
	)
	var (
		tp  string
		v   int32
		_n  int32 //防止重名
		d   []byte
		i   int32
		m   string
		cnt = 0
	)

	//是否记录了request m
	res, _ := db.Query("select * from replica" + ToString(rep.serialNumber) + " where type='PRE-PREPARE'" +
		" and v=" + ToString(rep.viewNumber) + " and n=" + ToString(n) + ";")
	if res.Next() {
		res.Scan(&tp, &v, &_n, &d, &i, &m)
		flag1 = true
	}
	if res.Next() {
		panic("日志中出现了n和v相同的多份记录")
	}

	//检验是否收到了来自2f个节点的prepare信息
	res, _ = db.Query("select * from replica" + ToString(rep.serialNumber) + " where type='PREPARE'" +
		" and v=" + ToString(rep.viewNumber) + " and n=" + ToString(n) + ";")
	for res.Next() {
		cnt++
		res.Scan(&tp, &v, &_n, &d, &i, &m)
		if cnt >= 2*f {
			flag2 = true
			break
		}
	}
	if cnt < 2*f {
		return false, nil
	}
	if flag1 && flag2 {
		msg := request{}
		json.Unmarshal([]byte(m), msg)
		return true, &msg
	}
}

/*
	3. 这里面应该有一个Commit方法，但是这个方法是可以通用的，不写。
*/
func (rep *Primary) Commit(n int32, digest []byte, private *rsa.PrivateKey) {
	// TODO 这里要根据Commit_Args的参数来指定具体参数
	// TODO 这里要广播所有节点发送commit的参数
	args := NewCommit(n, rep.viewNumber, rep.serialNumber, digest, private)
	println(args)
}

/**
4. 接受从其他节点发来的Commit（）参数（即为远程服务）
*/
func (t *Primary) Primary_Get_Commit(ctx context.Context, args *Commit_Msg, reply *interface{}) error {
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
//func (t *Primary) Get_CheckPoint(ctx context.Context, args *CheckPoint_Args, reply *CheckPoint_Reply) {
//
//}
