// logger is sample package
package logger

type logger interface {
	Print()
}

type message struct {
	messageMap map[string]string
}

type loggerImpl struct {
	label   string
	id      string
	content string
}

func (l *loggerImpl) Print() {
}

func NewLoggerImpl() *loggerImpl {
	return &loggerImpl{}
}
