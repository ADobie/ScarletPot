package panel

import (
	"github.com/gin-gonic/gin"
)

// TODO: 处理蜜罐上报信息 并存入数据库
func (s *Service) reportHandler(c *gin.Context) (int, interface{}) {
	var data SpInfo
	err := c.BindJSON(&data)
	if err != nil {
		return s.errJSON(500, 10000, "JSON解析失败")
	}

	spInfos := SpInfo{
		AttackIP:    data.AttackIP,
		ClientIP:    c.ClientIP(),
		AccessToken: data.AccessToken,
		Type:        data.Type,
		WebApp:      data.WebApp,
		Info:        data.Info,
		Count:       1,
		Country:     data.Country,
		City:        data.City,
		Region:      data.Region,
		Valid:       data.Valid,
	}

	if s.checkIfExist(data.AttackIP, data.Type, data.AccessToken) {
		s.updateInfo(spInfos)
		s.wsSend(s.dataInfo())
		return s.successJSON("")
	}

	s.insertFirst(spInfos)
	return s.successJSON("")
}

// 检查是否已经存在该攻击者记录
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

// 首次插入
func (s *Service) insertFirst(info SpInfo) {
	s.Mysql.Create(&info)
}

// 更新数据
func (s *Service) updateInfo(info SpInfo) {
	var oldInfo SpInfo
	// 旧数据拼接
	s.Mysql.Where(map[string]interface{}{"attack_ip": info.AttackIP, "type": info.Type}).Find(&oldInfo)
	s.Mysql.Model(&info).Where("attack_ip = ? AND type = ?", info.AttackIP, info.Type).Update("info", oldInfo.Info+"^^"+info.Info)
	// 更新攻击次数
	s.Mysql.Model(&info).Where("attack_ip = ? AND type = ?", info.AttackIP, info.Type).Update("count", oldInfo.Count+1)
}
