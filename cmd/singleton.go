package main

import (
	"fmt"
	"sync"

	"github.com/sivaram-codfolio/design_patterns/singleton"
)

var logger *singleton.MyLogger
var once sync.Once

func getLoggerInstance() *singleton.MyLogger {
	once.Do(func() {
		fmt.Println("Creating Logger Instance")
		logger = &singleton.MyLogger{}
	})
	fmt.Println("Returning Logger Instance")

	return logger
}

func main() {
	log := getLoggerInstance()

	log.SetLogLevel(1)
	log.Log("This is a log message")

	log = getLoggerInstance()
	log.SetLogLevel(2)
	log.Log("This is a log message")

	log = getLoggerInstance()
	log.SetLogLevel(3)
	log.Log("This is a log message")

	for i := 1; i < 10; i++ {
		go getLoggerInstance()
	}

	fmt.Scanln()
}
