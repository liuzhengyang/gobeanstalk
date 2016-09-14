# go-beanstalk 是[beanstalkd](https://github.com/kr/beanstalkd) 的GO语言的一个客户端.
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
package main

import (
	"fmt"
	"beanstalk-go/client"
)

func main() {
	addr := "localhost:11300"
	newConn := client.NewConnection(addr)
	channel := make(chan int)
	putFunc := func() {
		id, _ := newConn.PutWithTube("hello", "test2", 3)
		channel <- id
	}
	go putFunc()
	id := <-channel
	fmt.Printf("Receive from channel message of another goroutine %d", id)
	dealFunc := func(body string) bool{
		fmt.Printf("receive %s", body)
		return true
	}
	newConn.Listen("test2", dealFunc)
}

```