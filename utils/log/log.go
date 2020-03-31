package log

import (
	"fmt"
	"scarletpot/utils/color"
	"scarletpot/utils/i18n"
)

// 封装日志 标准输出 + 日志文件写入

func Err(lang string, info string) {
	fmt.Println(color.Red(i18n.I18nStr(lang, "[error] ")), color.Red(i18n.I18nStr(lang, info)))
}

func Warn(lang string, info string) {
	fmt.Println(color.Yellow(i18n.I18nStr(lang, "[warn] ")), color.Yellow(i18n.I18nStr(lang, info)))
}

func Succ(lang string, info string) {
	fmt.Println(color.Green(i18n.I18nStr(lang, "[succ] ")), color.Green(i18n.I18nStr(lang, info)))
}
