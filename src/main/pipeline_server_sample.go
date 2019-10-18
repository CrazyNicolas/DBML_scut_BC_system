//@author: mzy
package main

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/server"
	"runtime"
	"strconv"
	"sync"
	"time"
)

/**
定义一个判断素数的函数 因为我们这里对req的处理需要一个耗时操作
所以使用这个函数来做到比较耗时
*/
var IsPrime = func(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

/**
这里演示的是一个接收端 将服务通过rpcx暴露出去 然后开启多个发送端以高并发的形式向
这个接收端发送远程调用请求，接收端将通过pipeline构建一系列stage有效的对于请求进行处理
*/
//定义全局变量 relayStream 用以接收高频率过来的请求 考虑到可能是高频率请求所以设置一定缓冲
var relayStream = make(chan interface{}, 10)

//定义rpcx服务
type Service struct {
}

type Args struct {
	Info int
	Addr string
}

type Reply struct {
	Response bool
}

func (s *Service) Relay(c context.Context, req *Args, res *Reply) error {
	//直接将req转发
	relayStream <- req
	res.Response = true
	return nil
}

/**
以下是我们的pipeline的各stage的申明
*/
//这一步是类型断言Stage 这一阶段我们把接口流转换为特定类型流
func Toargs(done <-chan interface{}, values <-chan interface{}) <-chan *Args {
	argStream := make(chan *Args)
	go func() {
		defer close(argStream)
		for v := range values {
			select {
			case <-done:
				return
			case argStream <- v.(*Args):
			}
		}
	}()
	return argStream
}

/**
处理stage是所有的并行的扇出管道要做的任务 他们要负责验证req中节点地址的正确性
并且还要做一个比较耗时的判断素数的任务（模拟pbft中某些耗时操作）
*/
func process(done <-chan interface{}, args <-chan *Args) chan string {
	rStream := make(chan string)
	go func() {
		defer close(rStream)
		for arg := range args {
			//书写对一个arg请求的判断逻辑

			if len(arg.Addr) > 4 { //判断地址是否有效 简单地以长度来判断

				if IsPrime(arg.Info) { //判断请求中的计算任务是否是符合条件的 这是一个耗时操作
					logd := "Received a valid request from " + arg.Addr + " and this number is prime:  " + strconv.Itoa(arg.Info)
					select {
					case <-done:
						return
					case rStream <- logd:
					}
				}
			}
		}
	}()
	return rStream
}

/**
扇入stage十分关键 因为在这个过程中我们会对扇出的数据做验证和求素数的处理然后归一化为一个输出管道
在这里我们需要使用同步组的概念 ，因为同步组可以让我们明确的知道什么时候多条并行管道做完了自己的任务
在这个时间点我们就该关闭结果管道同时将这个关闭信号发给main goroutine用以做后续的工作
*/
func fan_in(done <-chan interface{}, channels ...chan string) chan string {
	resultStream := make(chan string) //处理结果输出的管道
	var wg sync.WaitGroup             //申明一个同步组锁 将扇出操作的多条并行管道的操作同步化
	wg.Add(len(channels))

	//匿名函数用来开启扇出goroutine的流动
	doWork := func(c <-chan string) {
		defer wg.Done()
		for arg := range c {
			select {
			case <-done:
				return
			case resultStream <- arg:
			}
		}
	}
	for _, c := range channels {
		go doWork(c)
	}

	//下面这个协程是为了逃逸出wait的锁区以至于我们可以返回我们的输出管道
	go func() {
		wg.Wait()
		close(resultStream) //当所有协程做完工作以后，我们需要做的是关闭它来通知外部消费者这里的工作结束了
	}()
	return resultStream
}

//最后我们需要一个take函数来及时的截断整个数据流 因为我们并不需要那么多数据 而客户端则一直在发
func take(done <-chan interface{}, responses <-chan string, num int) <-chan string {
	finalStream := make(chan string)
	go func() {
		defer close(finalStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case finalStream <- <-responses:
			}
		}
	}()
	return finalStream
}

/**
接下来我们就可以来组装这个pipeline然后获取结果了并且引入计时来对比效率
*/
func main() {
	done := make(chan interface{})
	defer close(done)
	start := time.Now() //计时起点
	//下面构造我们的扇出管道并行组
	max_channel_num := runtime.NumCPU()              //获取当前可工作cpu数目作为我们来做操作的并行组最大channel数
	channels := make([]chan string, max_channel_num) //申明一个用来放这些管道的slice 这是fan_in函数的关键参数
	for i := 0; i < max_channel_num; i++ {           //构建上述slice
		channels[i] = process(done, Toargs(done, relayStream))
	}
	//组装
	pipeline := take(done, fan_in(done, channels...), 20)

	//在pipeline组装好了以后我们着手把服务器的监听做好 但是这要做在一个goroutine里面 因为Serve()
	//函数是会阻塞进程的 那样的话我们会得到一个死锁！
	//s := server.NewServer()
	//s.RegisterName("Serv" , new(Service) , "")
	//s.Serve("tcp" , ":9090")
	s := server.NewServer()
	go func() {
		s.RegisterName("Serv", new(Service), "")
		s.Serve("tcp", ":9090")
	}()

	fmt.Println("test")

	//接下来输出结果
	for logd := range pipeline {
		fmt.Printf("%v\n", logd)
	}
	fmt.Printf("成功处理所有请求，用时： %v\n", time.Since(start))
	s.Close()
}

//@author: mzy
