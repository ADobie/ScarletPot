package report

import (
	"encoding/json"
	"fmt"
	"scarletpot/utils/log"
	"scarletpot/utils/request"
)

// 各服务统一上报格式
type Info struct {
	//ID          int    `json:"id"`
	AttackIP    string `json:"attackIp"`
	AccessToken string `json:"accessToken"`
	Type        string `json:"type"`
	WebApp      string `json:"webApp"`
	Detail      string `json:"detail"`
}

var a string

func buildJson(atype string, attackIP string, webApp string, detail string) []byte {
	info := Info{
		AttackIP:    attackIP,
		AccessToken: "from base.config",
		Type:        atype,
		WebApp:      webApp,
		Detail:      detail,
	}
	attackInfo, err := json.Marshal(info)
	if err != nil {
		log.Err("zh-CN", "json解析失败")
		panic(err)
	}
	//fmt.Println(attackInfo)
	return attackInfo
}

func ReportMysql(atype string, attackIP string, webApp string, detail string) {
	info := buildJson(atype, attackIP, webApp, detail)
	fmt.Println(string(info))
	_, err := request.PostJson("http://47.99.241.73:1234", info)
	if err != nil {
		log.Err("zh-CN", "与远程服务器断开连接")
		panic(err)
	}
}
