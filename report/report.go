package report

import (
	"encoding/json"
	"fmt"
	"scarletpot/utils/conf"
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
	Country     string `json:"country"`
	City        string `json:"city"`
	Valid       uint   `json:"valid"`
	Region      string `json:"region"`
}

// 平台回应结果结构
type RetMsg struct {
	Data  string `json:"data"`
	Error string `json:"error"`
	Msg   string `json:"msg"`
}

var apiUrl = "http://" + conf.GetUserConfig().Panel.PanelAddr + "/api/report"

func buildJson(atype string, attackIP string, webApp string, detail string, country string, city string, region string, valid uint) []byte {
	info := Info{
		AttackIP:    attackIP,
		AccessToken: "from base.config",
		Type:        atype,
		WebApp:      webApp,
		Detail:      detail,
		Country:     country,
		City:        city,
		Region:      region,
		Valid:       valid,
	}
	attackInfo, err := json.Marshal(info)
	if err != nil {
		log.Err("zh-CN", "json解析失败")
		panic(err)
	}
	return attackInfo
}

// 只需负责上报数据，具体是插入还是更新 由panel来负责
func Do(atype string, attackIP string, webApp string, detail string, country string, city string, region string, valid uint) {
	log.DoLogs("report")
	info := buildJson(atype, attackIP, webApp, detail, country, city, region, valid)
	fmt.Println(string(info))
	resp, err := request.PostJson(apiUrl, info)
	defer resp.Body.Close()
	if err != nil {
		log.Err("zh-CN", "", err)
		// TODO 后期做尝试断线重连
		//panic(err)
	}
}
