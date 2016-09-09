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
		id, _ := newConn.PutWithTube("hello", "test2", 1)
		channel <- id
	}
	go putFunc()
	id := <-channel
	fmt.Printf("Receive from channel message of another goroutine %d\n", id)
	listenChannel := make(chan string)
	dealFunc := func(body string) bool{
		fmt.Printf("receive %s\n", body)
		listenChannel <- body
		return true
	}
	go newConn.Listen("test2", dealFunc)
	body := <- listenChannel
	fmt.Printf("Listen once %s\n", body)
	newConn.Close()
}
