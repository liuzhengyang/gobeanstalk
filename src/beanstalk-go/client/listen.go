package client

type DealFun interface {
	deal(body string) bool
}