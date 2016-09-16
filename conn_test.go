package gobeanstalk

import (
	"fmt"
	"testing"
)

func TestNewConnection(t *testing.T) {
	addr := "localhost:11300"  // define server address
	newConn := NewConnection(addr)   // create new connection
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
