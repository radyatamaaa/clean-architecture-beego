package logger

import (
	"fmt"
	"time"
)

// Logger is project standard interface for logging
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

// L is the global instance of the logger
var L = &StdOutLogger{}

// StdOutLogger logs to standard out
type StdOutLogger struct{}

func (s StdOutLogger) timeNow() string {
	timestr := time.Now().String()[0:22]
	return timestr
}

// Debug logs message at debug level
func (s StdOutLogger) Debug(msg string, args ...interface{}) {
	logMark := fmt.Sprintf("%s [DEBUG] ", s.timeNow())
	fmt.Printf(logMark+msg+"\n", args...)
}

// Info logs message at info level
func (s StdOutLogger) Info(msg string, args ...interface{}) {
	logMark := fmt.Sprintf("%s [INFO] ", s.timeNow())
	fmt.Printf(logMark+msg+"\n", args...)
}

// Warn logs message at warn level
func (s StdOutLogger) Warn(msg string, args ...interface{}) {
	logMark := fmt.Sprintf("%s [WARN] ", s.timeNow())
	fmt.Printf(logMark+msg+"\n", args...)
}

// Error logs message at error level
func (s StdOutLogger) Error(msg string, args ...interface{}) {
	logMark := fmt.Sprintf("%s [ERROR] ", s.timeNow())
	fmt.Printf(logMark+msg+"\n", args...)
}
