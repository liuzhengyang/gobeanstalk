package client

import (
	"strconv"
	"fmt"
	"net"
	"bufio"
	"strings"
	"beanstalk-go/client/command"
)

type Conn struct {
	conn      net.Conn
	addr      string
	bufReader *bufio.Reader
	bufWriter *bufio.Writer
}

func NewConnection(hostAndIp string) *Conn {
	conn, _ := net.Dial("tcp", hostAndIp)
	//defer conn.Close()
	c := new(Conn)
	c.conn = conn
	c.addr = hostAndIp
	c.bufReader = bufio.NewReader(conn)
	c.bufWriter = bufio.NewWriter(conn)

	return c
}

func (this *Conn) Close() {
	this.conn.Close()
}

// haven't deal with fail conditions
func sendAndGetOneLine(conn *Conn, command string) string {
	conn.bufWriter.Write([]byte(command))
	conn.bufWriter.Flush()
	line, _, _ := conn.bufReader.ReadLine()
	res := string(line)
	fmt.Printf("res : %s\n", res)
	return res
}

func (this *Conn) PutWithTube(body string, tube string, delay int) (int, error) {
	this.Use(tube)
	return this.Put(body, delay)
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
	sendAndGetOneLine(this, command)
}

func (this *Conn) Watch(tube string) {
	command := fmt.Sprintf("watch %s\r\n", tube)
	sendAndGetOneLine(this, command)
}

func (this *Conn) Put(body string, delay int) (int, error) {
	command := client.NewPut(1, delay, 100, []byte(body))
	this.bufWriter.Write(command.GetBytes())
	this.bufWriter.Flush()
	line, _, _ := this.bufReader.ReadLine()
	fmt.Printf("Put answer %s\n", line)
	token := strings.Split(string(line), " ")
	fmt.Println(string(line))
	jobId, _ := strconv.Atoi(token[1])
	return jobId, nil
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

func (this *Conn) Listen(tube string, fun func(body string) bool) {
	listenConnection, _ := net.Dial("tcp", this.addr)
	newConn, _ := NewConn(listenConnection, this.addr)
	newConn.Use(tube)
	newConn.Watch(tube)
	for {
		id, data := Reserve(newConn)
		fmt.Printf("Receive %s\n", data)
		success := fun(data)
		fmt.Printf("Deal Result %s\n", success)
		newConn.deleteMessage(id)
	}
}

