package main

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
func main() {

}
