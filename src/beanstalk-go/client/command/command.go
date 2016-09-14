package client

type Command interface {
	GetBytes() []byte
}
