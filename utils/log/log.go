package log

import (
	"gopkg.in/gookit/color.v1"
	"scarletpot/utils/i18n"
)

// 封装日志 标准输出 + 日志文件写入
func Err(lang string, info string, err ...interface{}) {
	if err != nil {
		color.Red.Println(i18n.I18nStr(lang, "[ERR] "+info), err)
		WriteLog("error", err)
		return
	}
	WriteLog("error", info)
	color.Red.Println(i18n.I18nStr(lang, "[ERR] "+info))
}

func Info(lang string, info string) {
	color.Info.Println(i18n.I18nStr(lang, "[INFO] "), i18n.I18nStr(lang, info))
	WriteLog("info", info)

}

func Warn(lang string, info string) {
	color.Warn.Println(i18n.I18nStr(lang, "[WARN] "), i18n.I18nStr(lang, info))
	WriteLog("info", info)
}
