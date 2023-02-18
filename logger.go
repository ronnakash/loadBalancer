package main

import (
	"fmt"
	"time"
)



type Logger struct {
	log 	bool
}

func (l *Logger) Print(output string) {
	if l.log {
		fmt.Printf("[%s] %s\n", l.CurrTimeString(), output)
	}
}

func (l *Logger) PrintError(errMessage string) {
	if l.log {
		fmt.Printf("[%s] ERROR: %s\n", l.CurrTimeString(), errMessage)
	}
}

func (l *Logger) SetLogging(log bool) {
	l.log = log
}


func (l *Logger) CurrTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func NewLogger(enableLogging bool) *Logger {
	logger := &Logger{
		log: 	enableLogging,
	}
	return logger
}
