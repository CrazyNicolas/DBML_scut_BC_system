package PBFT_module

import (
	"context"
	"crypto/rsa"
	"fmt"
	"math/rand"
	"time"
)

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
	//msg := NewRequest(operation, time.Now().UnixNano(), pub, pri)
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

// @author：Zeyuan Ma  寻找下一个@author同名标记 之间的都是同一个作者写的
type Req struct {
}

func RandomRequest() interface{} {
	/**
	这个函数是作为main函数里 req_generator 生成器函数的参数使用 以随机的睡眠时间生成随机的请求
	*/
	time.Sleep(time.Duration(rand.Int31n(6)) * time.Second) //随机睡眠 0 - 6 s
	return &Req{}

}

func Test() {
	//@Test：以下是对于pipeline的一个小型演示  结束标签请往下寻找@Test
	//生成器函数 主要的是done用来控制goroutine在主线程结束的时候将generator里面的goroutine关闭防止泄露
	//而rStream就是生成的管道 这里把所有权交给了generator只暴露出单向管道给外面的使用者读取
	generator := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
		rStream := make(chan interface{})
		go func() {
			defer close(rStream)
			for {
				select {
				case <-done:
					return
				case rStream <- fn():
				}
			}
		}()
		return rStream
	}
	//类型断言函数 如果确定要使用空接口函数 那么一定不能少了这个环节
	//在新开的协程中 以flowdata的形式将前一个generator中产生的只读chan中的空接口类型向下转型
	toReq := func(done <-chan interface{}, values <-chan interface{}) <-chan *Req {
		rStream := make(chan *Req)
		go func() {
			defer close(rStream)
			for v := range values {
				select {
				case <-done:
					return
				case rStream <- v.(*Req):
				}
			}
		}()
		return rStream
	}
	//取出函数 是生成器函数的必备
	take := func(done <-chan interface{}, reqStream <-chan *Req, num int) <-chan *Req {
		resultStream := make(chan *Req)
		go func() {
			defer close(resultStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case resultStream <- <-reqStream:
				}
			}
		}()
		return resultStream
	}
	//下面我们来把pipieline组装起来调用:
	done := make(chan interface{})
	defer close(done)
	start := time.Now()
	pipeline := take(done, toReq(done, generator(done, RandomRequest)), 10)
	for req := range pipeline {
		fmt.Printf("经过了 %v 的时间产生了一个请求: %v\n", time.Since(start), req)
		start = time.Now()
	}
	fmt.Println("10个请求通过流水线作业生成完了！")

	//@Test
	/**
	以下的部分就是比较麻烦的reply的部分了 大致顺序如下
		收到reply ---->  验证reply -----> 统计reply
		与之并行的有: 检验超时 ----> 广播 request
	*/

}

//@author: zeyuan Ma
