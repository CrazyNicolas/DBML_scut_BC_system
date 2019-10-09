package PBFT_module

/**
 client 有以下几个流程：
	1.首先调用Request() 向 primary发起请求
	2.统计收到的Reply() 的return value 并加和其个数
 还需要有针对各种方法建立一些结构体 相应的要有args
*/

type Client struct {
}

//func Request(operation string, pub *rsa.PublicKey, pri *rsa.PrivateKey) {
//	// TODO 这里应该首先获得主节点的地址，并调用它的那个接收函数
//	msg := NewRequest(operation, time.Now().UnixNano(), pub, pri)
//
//}

func Request(operation string, pub *rsa.PublicKey, pri *rsa.PrivateKey) {
	// TODO 这里应该首先获得主节点的地址，并调用它的那个接收函数
	msg := NewRequest(operation, time.Now().UnixNano(), pub, pri)
	/**
	此部分要求获得主节点的地址，并要求广播调用其getRequest（）方法
	1.首先要求从数据库里取出类似的Reply（）消息（即为看有没有n，v，r消息）
	@author:Mingze Sun
	*/

}


/**
4. 接受从其他节点发来的Commit（）参数（即为远程服务）
*/

//func (t *Client) Get_Reply(ctx context.Context, args *Reply_Args, reply *Reply_Reply) error {
//	// TODO 这里面写处理Reply()的逻辑
//	return nil
//}

func (t *Client) Get_Reply(ctx context.Context, args Reply_Msg, reply *interface{}) error {
	// TODO 这里面写处理Reply()的逻辑
	/**
	此部分要进行reply回复消息的处理，但是这个Reply消息比较特殊，必须要从数据库取出过往数据，这点暂时还不能完成，先放下
	@author：Mingze Sun
	*/

	return nil
}

func main() {
}

