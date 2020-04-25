package log

import (
	"log"
	"os"
	"scarletpot/utils/conf"
)

func WriteLog(level string, content interface{}) {
	var fileName string
	switch level {
	case "error":
		fileName = conf.GetBaseConfig().Logs.ErrorLog
	// info 包括 warn
	case "info":
		fileName = conf.GetBaseConfig().Logs.InfoLog
	case "debug":
		fileName = conf.GetBaseConfig().Logs.DebugLog
	default:
		panic("日志等级错误")
	}

	logFile, _ := os.OpenFile("logs/"+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logger := log.New(logFile, "[Scarlet] ", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Printf("%v", content)

}
