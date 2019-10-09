package PBFT_module

import "crypto/rsa"

/**
message 文件的注释
@author mingze Sun
*/

/*
	request请求的内容，也就是所谓的m
*/
type request struct {
	operation string         //操作o
	timeStamp int64          //时间戳t
	publicKey *rsa.PublicKey //客户端标识c
}

/*
	远程调用getRequest方法所发送和接收的参数
*/
type Request_Msg struct {
	request          //request本身
	signature []byte //签名
}

/*
	新建一个Request_Msg请求的方法
*/
func NewRequest(op string, t int64, pub *rsa.PublicKey, pri *rsa.PrivateKey) Request_Msg {
	req := request{op, t, pub}
	sig := DigitalSignature(req, pri)
	return Request_Msg{req, sig}
}

/*
	preprepare中信息的一部分，也就是需要用私钥加密的一部分
*/
type preprepare struct {
	v      int32  //视图编号v
	n      int32  //主节点分配的序号n
	digest []byte //消息摘要d
}


/**
江声：用匿名成员代替了之前的preprepare preprepare，感觉这样更加清晰一点
=======
/*
	Pre-prepare（）消息体结构，也就是调用getPreprepare（）方法所发送和接收的参数

*/
type Prepreprare_Msg struct {
	preprepare        //preprepare本身
	request           //request本身
	signature  []byte //签名
}

/**

 */
func NewPreprepare(v, n int32, req request, pri *rsa.PrivateKey) Prepreprare_Msg {
	prepre := preprepare{v, n, Digest(req)}

/*
	新发起一个Preprepare—Msg（）请求的方法
*/
func NewPreprepare(n, v int32, req request, pri *rsa.PrivateKey) Prepreprare_Msg {
	prepre := preprepare{n, v, Digest(req)}

	sig := DigitalSignature(prepre, pri)
	return Prepreprare_Msg{prepre, req, sig}
}

/*
	prepare信息内容
*/
type prepare struct {
	v      int32  //视图编号
	n      int32  //主节点分配的序号
	digest []byte //消息摘要
	i      int32  //当前节点编号
}

/*
	Prepare（）消息体结构，也就是调用getPrepare（）方法所发送和接收的参数
*/
type Prepare_Msg struct {
	prepare          //prepare本身
	signature []byte //签名
}

/*
	新建一个Prepare-Msg（）请求的方法
*/
func NewPrepare(n, v, i int32, d []byte, pri *rsa.PrivateKey) Prepare_Msg {
	pre := prepare{
		v:      v,
		n:      n,
		digest: d,
		i:      i,
	}
	sig := DigitalSignature(pre, pri)
	return Prepare_Msg{pre, sig}
}

/*
	commit消息内容
*/
type commit struct {
	v      int32  //视图编号
	n      int32  //主节点分配的序号
	digest []byte //消息摘要
	i      int32  //当前节点编号
}

/*
	Commit消息体结构，也就是调用getCommit（）方法所发送和接收的参数
*/
type Commit_Msg struct {
	commit           //commit本身
	siganture []byte //签名
}

/*
	新建一个Commit请求方法，返回Commit——Msg（）
*/
func NewCommit(n, v, i int32, d []byte, pri *rsa.PrivateKey) Commit_Msg {
	comm := commit{
		v:      v,
		n:      n,
		digest: d,
		i:      i,
	}
	sig := DigitalSignature(comm, pri)
	return Commit_Msg{comm, sig}
}

/*
	由于Reply消息的特殊性，不需要私钥进行签名认证，所以没有消息内容，只有Reply（）消息结构体
*/
type Reply_Msg struct {
	v         int32          //当前视图编号v
	timeStamp int64          //时间戳t
	publicKey *rsa.PublicKey //客户端标识
	i         int32          //节点编号
	r         string         //操作结果
}

/*
	新建一个Reply（）消息所需的方法，返回Reply——Msg
*/
func NewReply(v int32, t int64, pub *rsa.PublicKey, i int32, res string) Reply_Msg {
	return Reply_Msg{
		v:         v,
		timeStamp: t,
		publicKey: pub,
		i:         i,
		r:         res,
	}
}
