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
		Count:       1,
	}

	if s.checkIfExist(data.AttackIP, data.Type, data.AccessToken) {
		s.updateInfo(spInfos)
		s.wsSend(s.getAttackCount())
		return s.successJSON("info数据更新成功")
	}
	s.insertFirst(spInfos)
	s.updateInfo(spInfos)
	return s.successJSON("数据上报成功")
}

func (s *Service) checkIfExist(attackIp string, attackType string, token string) bool {
	var data SpInfo
	res := s.Mysql.Where(map[string]interface{}{"attack_ip": attackIp, "type": attackType, "access_token": token}).Find(&data).RowsAffected

	// 攻击者单次攻击记录存在则返回true 否则返回false
	if res > 0 {
		return true
	} else {
		return false
	}
}

func (s *Service) insertFirst(info SpInfo) {
	s.Mysql.Create(&info)
}

func (s *Service) updateInfo(info SpInfo) {
	var oldInfo SpInfo
	// 旧数据拼接
	s.Mysql.Where(map[string]interface{}{"attack_ip": info.AttackIP, "type": info.Type}).Find(&oldInfo)
	s.Mysql.Model(&info).Where("attack_ip = ? AND type = ?", info.AttackIP, info.Type).Update("info", oldInfo.Info+"^^"+info.Info)
	// 更新攻击次数
	s.Mysql.Model(&info).Where("attack_ip = ? AND type = ?", info.AttackIP, info.Type).Update("count", oldInfo.Count+1)

}
