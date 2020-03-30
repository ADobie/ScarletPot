package conf

import (
	"gopkg.in/ini.v1"
	"log"
	"scarletpot/utils/color"
	"scarletpot/utils/i18n"
)

var cfg *ini.File

func init() {
	c, err := ini.Load("conf/config.ini")
	if err != nil {
		log.Println(color.Red(i18n.I18nStr("zh-CN", "install.config_load_fail")))
	}
	c.BlockMode = false
	cfg = c
}

func Get(node string, key string) string {
	val := cfg.Section(node).Key(key).String()
	return val
}

func GetInt(node string, key string) int {
	val, _ := cfg.Section(node).Key(key).Int()
	return val
}
