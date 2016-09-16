# [go-beanstalk]( 是[beanstalkd](https://github.com/kr/beanstalkd) 的GO语言的一个客户端.
项目还在开发中,欢迎大家提意见

# 介绍
[beanstalkd](https://github.com/kr/beanstalkd)是一个快速的、有各种用途的延迟队列
和定时任务的不同点:
定时任务以一定的周期或者在某个特定的时间运行。beanstalk可以在延迟一段时间执行。
一些使用场景:
* 用户下单5分钟后检查用户是否完成了支付
* 一分钟后开始一个新的程序

# 如何使用

## Mac&Linux
### 安装并启动beantalkd服务器
```
git clone https://github.com/kr/beanstalkd
cd beanstalkd
make
./beanstalkd
```

# 使用示例
```
go get github.com/liuzhengyang/gobeanstalk
```

create a test.go file
```
package main

import (
	"fmt"
	"github.com/liuzhengyang/gobeanstalk"
)

func main() {
	addr := "localhost:11300"  // define server address
	newConn := gobeanstalk.NewConnection(addr)   // create new connection
	channel := make(chan int)   // create int channel
	putFunc := func() {
		// define a function which put some message to one tube
		id, _ := newConn.PutWithTube("hello", "test2", 1)
		channel <- id
	}
	go putFunc()   // run previous function in a go-routine
	id := <-channel  // wait until we finish putting
	fmt.Printf("Receive from channel message of another goroutine %d\n", id)
	listenChannel := make(chan string)  // make a listen channel for receiving results
	dealFunc := func(body string) bool {
		// define a function to deal with tube messages
		fmt.Printf("receive %s\n", body)
		listenChannel <- body
		return true
	}
	go newConn.Listen("test2", dealFunc)  // run deal function in a specified go-routing
	body := <-listenChannel     // wait our message
	fmt.Printf("Listen once %s\n", body)
	newConn.Close()   // Close connection
}

```

And run this 
```
go run test.go
```