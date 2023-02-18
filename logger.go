package main

import (
	"fmt"
	"time"
)

type Logger interface{

	Print(output string)

	PrintError(errMessage string)

	EnableLogging()

	DisableLogging()

	CurrTimeString() string
}

type LoggerImpl struct {
	log 	bool
}

func (l *LoggerImpl) Print(output string) {
	if l.log {
		fmt.Printf("[%s] %s\n", l.CurrTimeString(), output)
	}
}

func (l *LoggerImpl) PrintError(errMessage string) {
	if l.log {
		fmt.Printf("[%s] ERROR: %s\n", l.CurrTimeString(), errMessage)
	}
}

func (l *LoggerImpl) EnableLogging() {
	l.log = true
}

func (l *LoggerImpl) DisableLogging() {
	l.log = false
}

func (l *LoggerImpl) CurrTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}