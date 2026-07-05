package service

import "fmt"

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
	PANIC
	TRACE
)

type LogLevel int

type ILogger interface {
	AddAppender(appender Appender)
	log(level LogLevel, message string)
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
	Panic(message string)
	Trace(message string)
}

type Logger struct {
	appenders []Appender
	minLevel  LogLevel
}

// AddAppender implements [ILogger].
func (l *Logger) AddAppender(appender Appender) {
	l.appenders = append(l.appenders, appender)
}

// log implements [ILogger].
func (l *Logger) log(level LogLevel, message string) {
	if level < l.minLevel {
		return
	}

	for _, appender := range l.appenders {
		appender.Write(fmt.Sprintf("[%s] %s", level, message))
	}
}

func NewLogger(minLevel LogLevel) ILogger {
	return &Logger{minLevel: minLevel, appenders: []Appender{}}
}

func (l *Logger) Debug(message string) {
	l.log(DEBUG, message)
}
func (l *Logger) Info(message string) {
	l.log(INFO, message)
}
func (l *Logger) Warn(message string) {
	l.log(WARN, message)
}
func (l *Logger) Error(message string) {
	l.log(ERROR, message)
}
func (l *Logger) Fatal(message string) {
	l.log(FATAL, message)
}
func (l *Logger) Panic(message string) {
	l.log(PANIC, message)
}
func (l *Logger) Trace(message string) {
	l.log(TRACE, message)
}
