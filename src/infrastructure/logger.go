package infrastructure

import "log"

type Logger interface {
	Log(message map[string]string)
}

type ConsoleLogger struct{}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

func (l *ConsoleLogger) Log(message map[string]string) {
	log.Default()
}
