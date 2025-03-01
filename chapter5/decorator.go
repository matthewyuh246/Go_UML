package main

import (
	"fmt"
	"time"
)

type Component interface {
	GetLogMessage(msg string) string
}

type Logger struct{}

func (l Logger) GetLogMessage(msg string) string {
	return msg
}

type TimestampDecorator struct {
	component Component
}

func (td TimestampDecorator) GetLogMessage(msg string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return td.component.GetLogMessage(fmt.Sprintf("%s %s", timestamp, msg))
}

type LogLevelDecorator struct {
	component Component
	logLevel  string
}

func (ld LogLevelDecorator) GetLogMessage(msg string) string {
	return ld.component.GetLogMessage(fmt.Sprintf("%s %s", ld.logLevel, msg))
}

// func main() {
// 	logger := Logger{}
// 	loglevel := LogLevelDecorator{component: logger, logLevel: "INFO"}
// 	timestamp := TimestampDecorator{component: loglevel}

// 	fmt.Println(logger.GetLogMessage("Design Pattern"))
// 	fmt.Println(loglevel.GetLogMessage("Design Pattern"))
// 	fmt.Println(timestamp.GetLogMessage("Design Pattern"))
// }
