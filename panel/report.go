package panel

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// TODO: 处理蜜罐上报信息 并存入数据库
func (s *Service) reportHandler(c *gin.Context) (int, interface{}) {
	var data SpInfo
	err := c.BindJSON(&data)
	if err != nil {
		return s.errJSON(500, 10000, "JSON解析失败")
	}
	fmt.Println("type: " + data.Type + " attackIp: " + data.AttackIP + " info: " + data.Info)

	spInfos := SpInfo{
		AttackIP:    data.AttackIP,
		ClientIP:    c.ClientIP(),
		AccessToken: data.AccessToken,
		Type:        data.Type,
		WebApp:      data.WebApp,
		Info:        data.Info,
	}
	s.Mysql.Create(&spInfos)
	return s.successJSON("上报成功")
}
