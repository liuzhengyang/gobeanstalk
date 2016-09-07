package main

import (
	"fmt"
	"beanstalk-go/client"
	"net"
	"bufio"
	"strings"
	"strconv"
)

type Conn struct {
	conn      net.Conn
	addr      string
	bufReader *bufio.Reader
	bufWriter *bufio.Writer
}

func NewConn(conn net.Conn, addr string) (*Conn, error) {
	c := new(Conn)
	c.conn = conn
	c.addr = addr
	c.bufReader = bufio.NewReader(conn)
	c.bufWriter = bufio.NewWriter(conn)
	return c, nil
}

func (this *Conn) Use(tube string) {
	command := fmt.Sprintf("use %s\r\n", tube)
	this.bufWriter.Write([]byte(command))
	this.bufWriter.Flush()
	line, _, _ := this.bufReader.ReadLine()
	fmt.Println(string(line))
}

func (this *Conn) Watch(tube string) {
	command := fmt.Sprintf("watch %s\r\n", tube)
	this.bufWriter.Write([]byte(command))
	this.bufWriter.Flush()
	line, _, _ := this.bufReader.ReadLine()
	fmt.Println(string(line))
}

func (this *Conn) Put(body string, delay int) {
	command := client.NewPut(1, delay, 100, []byte(body))
	this.bufWriter.Write(command.GetBytes())
	this.bufWriter.Flush()
	line, _, _ := this.bufReader.ReadLine()
	fmt.Println(string(line))
}

func Put(conn *Conn, body string, tube string, delay int) {
	conn.Use(tube)
	conn.Put(body, delay)
}

func Reserve(conn *Conn) (int, string) {
	command := []byte("reserve\r\n")
	conn.bufWriter.Write(command)
	conn.bufWriter.Flush()
	line, _, _ := conn.bufReader.ReadLine()
	dataline, _, _ := conn.bufReader.ReadLine()
	tokens := strings.Split(string(line), " ")
	idstr := tokens[1]
	id, _ := strconv.Atoi(idstr)
	fmt.Printf("Reserve %s\n", string(line))
	fmt.Printf("Reserve %s\n", string(dataline))
	return id, string(dataline)
}

func (this *Conn) deleteMessage(id int) {
	commandStr := fmt.Sprintf("delete %d\r\n", id)
	command := []byte(commandStr)
	this.bufWriter.Write(command)
	this.bufWriter.Flush()
	line, _, _ := this.bufReader.ReadLine()
	fmt.Printf("delete %s\n", string(line))
}

func Listen(conn *Conn, tube string) {
	listenConnection, _ := net.Dial("tcp", conn.addr)
	newConn, _ := NewConn(listenConnection, conn.addr)
	newConn.Use(tube)
	newConn.Watch(tube)
	for {
		id, data := Reserve(newConn)
		fmt.Printf("Receive %s\n", data)
		newConn.deleteMessage(id)
	}
}

func main() {
	addr := "localhost:11300"
	conn, e := net.Dial("tcp", addr)
	defer conn.Close()
	if e != nil {
		panic("connect error")
	}
	newConn, _ := NewConn(conn, addr)
	go Put(newConn, "hello", "test2", 3)
	Listen(newConn, "test2")
}
