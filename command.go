package gobeanstalk

type Command interface {
	GetBytes() []byte
}
