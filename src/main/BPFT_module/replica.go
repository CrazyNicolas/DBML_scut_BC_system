package main

import "context"

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

type Replica struct {
	serialNumber int
	viewNumber   int
}

/**
1. 接受从Primary来的pre-prepare()参数（即为远程服务）
*/
func (t *Replica) Get_Pre_prepare(ctx context.Context, args *Pre_prepare_Args, reply *Pre_prepare_Reply) error {
	// TODO 对Primary的pre-prepare()请求进行校验
	// （1）查看请求的签名是否正确
	//  (2)

	return nil
}

/**
2. prepare()方法
*/
type Prepare_Args struct {
}

type Prepare_Reply struct {
}

func Prepare() {
	// TODO 这里要根据Prepare_Args的参数来指定具体参数
	// TODO 这里要广播所有节点发送prepare的参数
}

/**
3. 接受从其他Replica发来的prepare（）参数（即为远程服务）
*/
func (t *Replica) Replica_Get_Prepare(ctx context.Context, args *Prepare_Args, reply *Prepare_Reply) error {
	// TODO 这里面写处理Prepare()的逻辑，如果正确的话执行commit()
	return nil
}

/**
4. Commit()方法
*/
type Commit_Args struct {
}

type Commit_Reply struct {
}

func Commit() {
	// TODO 这里要根据Commit_Args的参数来指定具体参数
	// TODO 这里要广播所有节点发送commit的参数
}

/**
5. 接受从其他节点发来的Commit（）参数（即为远程服务）
*/
func (t *Replica) Replica_Get_Commit(ctx context.Context, args *Commit_Args, reply *Commit_Reply) error {
	// TODO 这里面写处理Commit()的逻辑，如果正确的话执行Reply()
	return nil
}

/**
6. Reply()方法
*/
type Reply_Args struct {
}

type Reply_Reply struct {
}

func Reply() {
	// TODO 这里要根据Reply_Args的参数来指定具体参数
	// TODO 这里要广播所有节点发送reply的参数
}

func main() {

}
