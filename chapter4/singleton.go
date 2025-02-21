package main

import (
	"fmt"
	"sync"
	"time"
)

type Logger struct{}

var loggerInstance *Logger
var once sync.Once

func GetLoggerInstance() *Logger {
	once.Do(func() {
		loggerInstance = &Logger{}
	})
	return loggerInstance
}

func (l *Logger) Output(content string) {
	now := time.Now()
	fmt.Printf("%s: %s\n", now.Format("2006-01-02 15:04:05"), content)
}

type Test struct{}

func main() {
	test1 := &Test{}
	test2 := &Test{}
	fmt.Println("Testクラス: ", test1 == test2)

	logger1 := GetLoggerInstance()
	logger2 := GetLoggerInstance()
	fmt.Println("Loggerクラス: ", logger1 == logger2)

	logger1.Output("logger1のクラス")
	logger2.Output("logger2のクラス")
}
