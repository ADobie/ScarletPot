package panel

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// TODO: 处理蜜罐上报信息 并存入数据库

type Info struct {
	AttackIP    string `json:"attackIp"`
	AccessToken string `json:"accessToken"`
	Type        string `json:"type"`
	WebApp      string `json:"webApp"`
	Detail      string `json:"detail"`
}

func (s *Service) reportHandler(c *gin.Context) {
	var data Info
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	}
	fmt.Println("type: " + data.Type + " attackIp: " + data.AttackIP + " info: " + data.Detail)
}
