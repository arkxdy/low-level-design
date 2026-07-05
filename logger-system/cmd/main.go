package main

import "logger-system/internal/service"

func main() {
	console := service.NewConsoleAppender()
	file := service.NewFileAppender("log.txt")

	logger := service.NewLogger(service.INFO)
	logger.AddAppender(console)
	logger.AddAppender(file)
	logger.Debug("Hello, world!")
	logger.Error("Hello, world!")
	logger.Debug("Hello, world!")
	logger.Warn("Hello, world!")
	logger.Fatal("Hello, world!")
	logger.Panic("Hello, world!")
	logger.Trace("Hello, world!")
}
