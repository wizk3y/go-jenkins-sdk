package logger

import (
	"fmt"
	"html/template"
	"io"
	"sync"
)

const (
	levelError = iota
	levelWarn
	levelInfo
	levelDebug
)

var (
	instance          *defaultLogger
	once              sync.Once
	logRecordTemplate *template.Template
)

type defaultLogger struct {
	mu             sync.Mutex
	output         io.Writer
	verbosityLevel int
}

type LogRecord struct {
	Message string
}

func GetDefaultLogger(w io.Writer) *defaultLogger {
	once.Do(func() {
		var (
			err       error
			logFormat = `{{.Message}}{{EndLine}}`
		)

		// Initialize and parse logging templates
		funcs := template.FuncMap{
			"EndLine": EndLine,
		}
		logRecordTemplate, err = template.New("logFormat").Funcs(funcs).Parse(logFormat)
		if err != nil {
			panic(err)
		}

		instance = &defaultLogger{output: w, verbosityLevel: levelInfo}
	})
	return instance
}

// SetVerbosityLevel sets the logger verbosity level
func (l *defaultLogger) SetVerbosityLevel(level int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.verbosityLevel = level
}

// SetOutput sets the logger output destination
func (l *defaultLogger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.output = w
}

// EndLine returns the a newline escape character
func EndLine() string {
	return "\n"
}

// mustLog logs the message according to the specified level and arguments.
// It panics in case of an error.
func (l *defaultLogger) mustLog(level int, message string, args ...interface{}) {
	if level > l.verbosityLevel {
		return
	}
	// Acquire the lock
	l.mu.Lock()
	defer l.mu.Unlock()

	// Create the logging record and pass into the output
	record := LogRecord{
		Message: fmt.Sprintf(message, args...),
	}

	err := logRecordTemplate.Execute(l.output, record)
	if err != nil {
		panic(err)
	}
}

// Debug outputs a debug log message
func (l *defaultLogger) Debug(message string) {
	l.mustLog(levelDebug, message)
}

// Debugf outputs a formatted debug log message
func (l *defaultLogger) Debugf(message string, vars ...interface{}) {
	l.mustLog(levelDebug, message, vars...)
}

// Info outputs an information log message
func (l *defaultLogger) Info(message string) {
	l.mustLog(levelInfo, message)
}

// Infof outputs a formatted information log message
func (l *defaultLogger) Infof(message string, vars ...interface{}) {
	l.mustLog(levelInfo, message, vars...)
}

// Warn outputs a warning log message
func (l *defaultLogger) Warn(message string) {
	l.mustLog(levelWarn, message)
}

// Warnf outputs a formatted warning log message
func (l *defaultLogger) Warnf(message string, vars ...interface{}) {
	l.mustLog(levelWarn, message, vars...)
}

// Error outputs an error log message
func (l *defaultLogger) Error(message string) {
	l.mustLog(levelError, message)
}

// Errorf outputs a formatted error log message
func (l *defaultLogger) Errorf(message string, vars ...interface{}) {
	l.mustLog(levelError, message, vars...)
}
