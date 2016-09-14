package beanstalk

import (
	"fmt"
)

type Put struct {
	pri   int
	delay int
	ttr   int
	//bytes int32
	body  []byte
}

func NewPut(pri int, delay int, ttr int, body []byte) *Put {
	return &Put{pri:pri, delay:delay, ttr:ttr, body:body}
}

func (this *Put) GetBytes() []byte {
	str := fmt.Sprintf("put %d %d %d %d\r\n", this.pri, this.delay, this.ttr, len(this.body))
	res := []byte(str)
	res2 := append(res, this.body...)
	res3 := append(res2, []byte("\r\n")...)
	return res3
}
