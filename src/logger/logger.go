package logger

import (
	"log"
	"os"
)

func InitLogger() *log.Logger {
	logFileName := "TgJobNotifier.log"
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	return log.New(logFile, "", log.Ldate|log.Ltime)
}
