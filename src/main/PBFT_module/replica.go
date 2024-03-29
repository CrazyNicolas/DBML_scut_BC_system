package PBFT_module

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/**
  副本节点replica有以下几个流程：
	1.接收来自 primary的 参数给 Prepare()
	2.调用Prepare() 向所有其他节点触发 Commit()
	3.接收来自其他节点的参数给Commit()
	4.调用Commit()去触发其他节点的Reply()
	5.收到其他节点的参数给Reply()
	6.调用Reply()返回结果给Client
	7.Timer结构体及其绑定的方法是用来记录主节点行为是否超时
	8.Checkpoint()用来检查一个其他节点发来的请求
	9.View_Change()发起更换主节点操作 是要广播给其他所有relica
	10.New_View() 先接收到足够数量的viewchange消息之后 向其他所有节点广播自己变成了primary
 还需要有针对各种方法建立一些结构体 相应的要有args
*/

/**
Replica应该是一个服务
*/
type Replica struct {
	serialNumber int32
	viewNumber   int32
	h            int32
	H            int32
}

/**
江声 全局变量f表示拜占庭节点数
*/
var f = 1

/**
江声：
创建一个replica对象，并且判断数据库中是否存在他的日志，如果没有则创建，表的命名为replica i
而且，需要在数据库中生成密钥对
*/

func NewReplica(i, v, h, H int32) Replica {
	db, err := sql.Open("mysql", db_address)
	if err != nil {
		fmt.Println("数据库打开失败", err)
	}
	//判断数据表是否存在
	rows, err := db.Query("select * from information_schema.tables where table_name = 'replica" + ToString(i) + "';")
	if err != nil {
		fmt.Println("查询information_schema出错", err)
	}
	if !rows.Next() { //表不存在
		_, err := db.Exec("create table replica" + ToString(i) + " (" +
			"type enum('PRE-PREPARE', 'PREPARE', 'COMMIT')," +
			"v int," +
			"n int," +
			"d varchar(50)," +
			"i int," +
			"m json);")
		if err != nil {
			fmt.Println("数据表创建失败", err)
		}
	}

	//判断数据库中是否存在节点i的密钥对，若没有则生成
	rows, err = db.Query("select * from information_schema.tables where table_name = 'key'" + ToString(i))
	if err != nil {
		fmt.Println("查询information_schema出错", err)
	}

	return Replica{
		serialNumber: i,
		viewNumber:   v,
		h:            h,
		H:            H,
	}

	//TODO: 广播prepare消息
	//return nil
}

/**
2. prepare()方法
*/
//func (rep *Replica) Prepare(n int32, digest []byte, private *rsa.PrivateKey) {
//	args := NewPrepare(n, rep.viewNumber, rep.serialNumber, digest, private)
//	// TODO 这里要广播所有节点发送prepare的参数
//	//TODO 记录到log
//}

/**
3. 接受从其他Replica发来的prepare（）参数（即为远程服务）
*/
func (t *Replica) Replica_Get_Prepare(ctx context.Context, args *Prepare_Msg, reply *interface{}) error {
	// TODO 这里面写处理Prepare()的逻辑，如果正确的话执行commit()
	/*
		这部分
	*/
	return nil
}

/**
4. Commit()方法
*/
//func (rep *Replica) Commit(n int32, digest []byte, private *rsa.PrivateKey) {
//	// TODO 这里要根据Commit_Args的参数来指定具体参数
//	// TODO 这里要广播所有节点发送commit的参数
//	args := NewCommit(n, rep.viewNumber, rep.serialNumber, digest, private)
//}

/**
5. 接受从其他节点发来的Commit（）参数（即为远程服务）
*/
//func (t *Replica) Replica_Get_Commit(ctx context.Context, args *Commit_Msg, reply *interface{}) error {
//	// TODO 这里面写处理Commit()的逻辑，如果正确的话执行Reply()
//	return nil
//
//}

///**
//1. 接受从Primary来的pre-prepare()参数（即为远程服务）
//*/
//func (t *Replica) Get_Pre_prepare(ctx context.Context, args *Prepreprare_Msg, reply *interface{}) error {
//	// 对Primary的pre-prepare()请求进行校验
//	// （1）判断d和m是否一致
//	digest := Digest(args.request)
//	if digest != args.preprepare.digest {
//		return nil
//	}
//	//  (2)判断
//	bytes, _ := json.Marshal(args)
//	if !Verify_ds(args.signature, "public.pem", bytes) {
//		return nil
//	}
//	//(3)查看当前view是否与pre-prepare中的view相同
//	if t.viewNumber != args.preprepare.v {
//		return nil
//	}
//	//(4)TODO：查看当前replica是否接受过一个v,n相同但是d不同的pre-prepare请求
//	//(5)判断水线
//	if !(args.preprepare.n < t.H && args.preprepare.n > t.h) {
//		return nil
//	}
//
//	//TODO: 广播prepare消息
//	return nil
//}
//
///**
//2. prepare()方法
//*/
//func (rep *Replica) Prepare(n int32, digest []byte, private *rsa.PrivateKey) {
//	args := NewPrepare(n, rep.viewNumber, rep.serialNumber, digest, private)
//	// TODO 这里要广播所有节点发送prepare的参数
//	//TODO 记录到log
//}
//
///**
//3. 接受从其他Replica发来的prepare（）参数（即为远程服务）
//*/
//func (t *Replica) Replica_Get_Prepare(ctx context.Context, args *Prepare_Msg, reply *interface{}) error {
//	// TODO 这里面写处理Prepare()的逻辑，如果正确的话执行commit()
//	return nil
//}
//
///**
//4. Commit()方法
//*/
//func (rep *Replica) Commit(n int32, digest []byte, private *rsa.PrivateKey) {
//	// TODO 这里要根据Commit_Args的参数来指定具体参数
//	// TODO 这里要广播所有节点发送commit的参数
//	args := NewCommit(n, rep.viewNumber, rep.serialNumber, digest, private)
//}
//
///**
//5. 接受从其他节点发来的Commit（）参数（即为远程服务）
//*/
//func (t *Replica) Replica_Get_Commit(ctx context.Context, args *Commit_Msg, reply *interface{}) error {
//	// TODO 这里面写处理Commit()的逻辑，如果正确的话执行Reply()
//	return nil
//}
//
///**
//6. Reply()方法
//*/
//func (t *Replica) Reply(time int64, pub *rsa.PublicKey, res string) {
//	// TODO 这里要根据Reply_Args的参数来指定具体参数
//	// TODO 这里要广播所有节点发送reply的参数
//	args := NewReply(t.viewNumber, time, pub, t.serialNumber, res)
//}
