package beanstalk

type Command interface {
	GetBytes() []byte
}
