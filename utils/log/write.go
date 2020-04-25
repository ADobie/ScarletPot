package log

import (
	"log"
	"os"
)

func WriteLog(level string, content interface{}) {
	var fileName string
	switch level {
	case "error":
		fileName = "error.log"
	// info 包括 warn
	case "info":
		fileName = "info.log"
	case "debug":
		fileName = "debug.log"
	default:
		panic("日志等级错误")
	}

	logFile, _ := os.OpenFile("logs/"+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logger := log.New(logFile, "[Scarlet] ", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Printf("%v", content)

}
