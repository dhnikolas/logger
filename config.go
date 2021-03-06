package logger

const (
	ALERT = iota
	ERROR
	LOG
	DEBUG
	TRACE
)

type LoggerConfig struct {
	ServiceName string
	Level       int
	Buffer      int
	Output      []LogDriver
}

type LogDriver interface {
	PutMsg(msg Message) error
	Init() error
}
