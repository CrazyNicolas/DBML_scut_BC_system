package main

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
func main() {

}
