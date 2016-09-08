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
	newConn.Listen("test2")
}
