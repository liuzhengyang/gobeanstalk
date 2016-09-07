## This is [beanstalkd](https://github.com/kr/beanstalkd) client for Go.
Now it's a quite simple implementation. I'll keep improving and refactoring it.

## Examples

```
    addr := "localhost:11300"
	conn, e := net.Dial("tcp", addr)
	defer conn.Close()
	if e != nil {
		panic("connect error")
	}
	newConn, _ := NewConn(conn, addr)
	go Put(newConn, "hello", "test2", 3)
	Listen(newConn, "test2")
```