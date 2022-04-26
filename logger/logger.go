package logger

import "os"

var (
	Logger LoggerInterface
)

type LoggerInterface interface {
	Debug(message string)
	Debugf(message string, vars ...interface{})
	Info(message string)
	Infof(message string, vars ...interface{})
	Warn(message string)
	Warnf(message string, vars ...interface{})
	Error(message string)
	Errorf(message string, vars ...interface{})
}

func init() {
	Logger = GetDefaultLogger(os.Stdout)
}
