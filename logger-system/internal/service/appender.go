package service

import "fmt"

type Appender interface {
	Write(message string)
}

type consoleAppender struct {
}

// Write implements [Appender].
func (c consoleAppender) Write(message string) {
	fmt.Println("Console: ", message)
}

type fileAppender struct {
	path string
}

// Write implements [Appender].
func (f fileAppender) Write(message string) {
	fmt.Println("File: ", message)
}

func NewConsoleAppender() Appender {
	return consoleAppender{}
}

func NewFileAppender(path string) Appender {
	return fileAppender{path: path}
}
