package logger

import (
	"io"
)

type dummyLogger struct {
}

func GetDummyLogger(w io.Writer) *dummyLogger {
	return &dummyLogger{}
}

// Debug outputs a debug log message
func (l *dummyLogger) Debug(message string) {
}

// Debugf outputs a formatted debug log message
func (l *dummyLogger) Debugf(message string, vars ...interface{}) {
}

// Info outputs an information log message
func (l *dummyLogger) Info(message string) {
}

// Infof outputs a formatted information log message
func (l *dummyLogger) Infof(message string, vars ...interface{}) {
}

// Warn outputs a warning log message
func (l *dummyLogger) Warn(message string) {
}

// Warnf outputs a formatted warning log message
func (l *dummyLogger) Warnf(message string, vars ...interface{}) {
}

// Error outputs an error log message
func (l *dummyLogger) Error(message string) {
}

// Errorf outputs a formatted error log message
func (l *dummyLogger) Errorf(message string, vars ...interface{}) {
}
