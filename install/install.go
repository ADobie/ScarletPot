package install

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"scarletpot/utils/color"
	"scarletpot/utils/i18n"
)

const configTemplate = `
[base]
baseAddr = "{{.baseAddr}}"                  # 管理后台url 例: 0.0.0.0:9090
SystemLanguage = "{{.SystemLanguage}}"      # 系统语言 [1] zh-CN [2] en-US  

[rpc]
status = 0                                  # rpc状态  0 关闭 1 开启
addr = 

`

var cfg *ini.File

func init() {
	f, err := os.Create("conf/config.ini")
	if err != nil {

	}
	c, err := ini.Load("conf/config.ini")
	if err != nil {
		log.Println(color.Red(i18n.I18nStr("zh-CN", "install.config_load_fail")))
	}
	c.BlockMode = false
	cfg = c
}

func Install() {
	fmt.Println("123")
}
