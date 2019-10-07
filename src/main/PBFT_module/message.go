package main

import "crypto/rsa"

type request struct {
	operation string         //操作o
	timeStamp int64          //时间戳t
	publicKey *rsa.PublicKey //客户端标识c
}

type Request_Msg struct {
	request   request //request本身
	signature []byte  //签名
}

func NewRequest(op string, t int64, pub *rsa.PublicKey, pri *rsa.PrivateKey) Request_Msg {
	req := request{op, t, pub}
	sig := DigitalSignature(req, pri)
	return Request_Msg{req, sig}
}

type preprepare struct {
	n      int32  //主节点分配的序号n
	v      int32  //视图编号v
	digest []byte //消息摘要d
}

/**

 */
type Prepreprare_Msg struct {
	preprepare preprepare //preprepare本身
	request    request    //request本身
	signature  []byte     //签名
}

/**

 */
func NewPreprepare(n, v int32, req request, pri *rsa.PrivateKey) Prepreprare_Msg {
	prepre := preprepare{n, v, Digest(req)}
	sig := DigitalSignature(prepre, pri)
	return Prepreprare_Msg{prepre, req, sig}
}

type prepare struct {
	v      int32  //视图编号
	n      int32  //主节点分配的序号
	digest []byte //消息摘要
	i      int32  //当前节点编号
}

type Prepare_Msg struct {
	prepare   prepare //prepare本身
	signature []byte  //签名
}

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

type commit struct {
	v      int32  //视图编号
	n      int32  //主节点分配的序号
	digest []byte //消息摘要
	i      int32  //当前节点编号
}

type Commit_Msg struct {
	commit    commit //commit本身
	siganture []byte //签名
}

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

type Reply_Msg struct {
	v         int32          //当前视图编号v
	timeStamp int64          //时间戳t
	publicKey *rsa.PublicKey //客户端标识
	i         int32          //节点编号
	r         string         //操作结果
}

func NewReply(v int32, t int64, pub *rsa.PublicKey, i int32, res string) Reply_Msg {
	return Reply_Msg{
		v:         v,
		timeStamp: t,
		publicKey: pub,
		i:         i,
		r:         res,
	}
}
