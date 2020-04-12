package report

import (
	"encoding/json"
	"fmt"
	"scarletpot/utils/log"
	"scarletpot/utils/request"
)

// 各服务统一上报格式
type Info struct {
	AttackIP    string `json:"attackIp"`
	AccessToken string `json:"accessToken"`
	Type        string `json:"type"`
	WebApp      string `json:"webApp"`
	Detail      string `json:"detail"`
}

// 平台回应结果结构

type RetMsg struct {
	Data  string `json:"data"`
	Error string `json:"error"`
	Msg   string `json:"msg"`
}

var a string

const apiUrl = "http://localhost:9000/api/report"

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
	return attackInfo
}

// 只需负责上报数据，具体是插入还是更新 由panel来负责
func Do(atype string, attackIP string, webApp string, detail string) {
	info := buildJson(atype, attackIP, webApp, detail)
	fmt.Println(string(info))
	resp, err := request.PostJson(apiUrl, info)
	defer resp.Body.Close()
	if err != nil {
		log.Err("zh-CN", "与远程上报服务器断开连接")
		// TODO 后期做尝试断线重连
		panic(err)
	}
}

//func UpdateDo(atype string, attackIP string, webApp string, detail string) string {
//
//}
