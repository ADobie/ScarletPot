package panel

import (
	"github.com/gin-gonic/gin"
)

// TODO: 处理蜜罐上报信息 并存入数据库
var data SpInfo

func (s *Service) reportHandler(c *gin.Context) (int, interface{}) {
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
		Valid:       1,
		Invalid:     1,
	}
	if s.checkIfExist(data.AttackIP, data.Type, data.AccessToken) {
		s.updateInfo(spInfos)
		s.wsSend(s.dataInfo())
		//log.DoLogs("success")
		return s.successJSON("")
	}

	s.insertFirst(spInfos)
	//log.DoLogs("success")
	return s.successJSON("")
}

// 检查是否已经存在该攻击者记录
func (s *Service) checkIfExist(attackIp string, attackType string, token string) bool {
	var datax SpInfo
	res := s.Db.Where(map[string]interface{}{"attack_ip": attackIp, "type": attackType, "access_token": token}).Find(&datax).RowsAffected

	// 攻击者单次攻击记录存在则返回true 否则返回false
	if res > 0 {
		return true
	} else {
		return false
	}
}

// 首次插入
func (s *Service) insertFirst(info SpInfo) {
	s.Db.Create(&info)
}

// 更新数据
func (s *Service) updateInfo(info SpInfo) {
	var oldInfo SpInfo
	// 旧数据拼接
	s.Db.Where(map[string]interface{}{"attack_ip": info.AttackIP, "type": info.Type}).Find(&oldInfo)
	if data.Valid == 1 {
		//fmt.Println(data.Valid)
		s.Db.Model(&info).Where("attack_ip = ? AND type = ?", info.AttackIP, info.Type).Update("valid", oldInfo.Valid+1)
	} else {
		s.Db.Model(&info).Where("attack_ip = ? AND type = ?", info.AttackIP, info.Type).Update("invalid", oldInfo.Invalid+1)
	}
	s.Db.Model(&info).Where("attack_ip = ? AND type = ?", info.AttackIP, info.Type).Update("info", oldInfo.Info+"^^"+info.Info)
	// 更新攻击次数
	s.Db.Model(&info).Where("attack_ip = ? AND type = ?", info.AttackIP, info.Type).Update("count", oldInfo.Count+1)
}
