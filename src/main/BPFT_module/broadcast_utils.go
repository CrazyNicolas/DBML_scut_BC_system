package main

/**
定义一些方法：
	1.获取所有节点的地址 GetRegisterDir()  ---IO
	2.Broadcast2Others() --根据参数动态调用 rpcx.broadcast()
	3.修改注册表 Change_RegisterDir() ----让新主节点的key变成primary 让旧主节点的key变为replica
*/
func GetRegisterDir() [][][]string {

}
