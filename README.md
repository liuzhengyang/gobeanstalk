# This is a [beanstalkd](https://github.com/kr/beanstalkd) client for Go.
[中文文档] (./README-zh_CN.md)
Now it's a quite simple implementation. I'll keep improving and refactoring it.

# Introduction
[beanstalkd](https://github.com/kr/beanstalkd) is a fast, general-purpose work queue.
Difference with crontab jobs:
Contab job run with specified period or at some point. But beanstalk can run with a delayed time.
Some use scenarios:
* Check whether user finish the order in 5 minutes.
* Start a process in one minutes.

# How to use

## Mac
### Install And Start Beanstalk server
```
git clone https://github.com/kr/beanstalkd
cd beanstalkd
make
./beanstalkd
```

# Examples

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