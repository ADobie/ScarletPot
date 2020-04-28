package log

import (
	"encoding/json"
	"log"
	"os"
	"scarletpot/utils/conf"
	"scarletpot/utils/request"
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

var apiLogUrl = "http://" + conf.GetUserConfig().Panel.PanelAddr + "/api/log"

// 日志信息上报格式
type Log struct {
	Level       string `json:"level"`
	AccessToken string `json:"accessToken"`
}

func DoLogs(level string) {
	info := Log{
		Level:       level,
		AccessToken: "from base.config",
	}
	logInfo, _ := json.Marshal(info)

	resp, _ := request.PostJson(apiLogUrl, logInfo)
	defer resp.Body.Close()
}
