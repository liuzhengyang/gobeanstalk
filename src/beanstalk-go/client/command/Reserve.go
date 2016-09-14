package client

import "fmt"

type Reserve struct {
	timeout int
}

func ReserveWithTimeOutBytes(timeout int) *Reserve{
	reserve := new(Reserve)
	reserve.timeout = timeout
	return reserve;
}

func ReserveBytes() *Reserve {
	reserve := new(Reserve)
	return reserve
}

func (this *Reserve) GetBytes() []byte {
	var command string
	if this.timeout > 0 {
		command = "reserve\r\n"
	} else {
		command = fmt.Sprintf("reserve-with-timeout %d\r\n", this.timeout)
	}
	return []byte(command)
}