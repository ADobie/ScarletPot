package panel

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) logsHandler(c *gin.Context) (int, interface{}) {
	var data SpLog
	err := c.BindJSON(&data)
	if err != nil {
		return s.errJSON(500, 10000, "JSON解析失败")
	}

	spLogs := SpLog{
		Level:       data.Level,
		ClientIP:    c.ClientIP(),
		AccessToken: data.AccessToken,
		Count:       1,
	}

	if s.checkIfLogExist(data.Level, data.AccessToken) {
		s.updateLog(spLogs)
		s.wsSend(s.dataInfo())
		//log.DoLogs("report")
		return s.successJSON("")
	}
	s.insertLogFirst(spLogs)
	s.wsSend(s.dataInfo())
	//log.DoLogs("report")

	//s.updateLog(spLogs)
	return s.successJSON("")
}

func (s *Service) insertLogFirst(info SpLog) {
	s.Mysql.Create(&info)
}

func (s *Service) updateLog(info SpLog) {
	var oldlogs SpLog
	// 旧数据拼接
	s.Mysql.Where(map[string]interface{}{"level": info.Level, "access_token": info.AccessToken}).Find(&oldlogs)
	// 更新攻击次数
	s.Mysql.Model(&info).Where("level = ? AND access_token = ?", info.Level, info.AccessToken).Update("count", oldlogs.Count+1)
}

func (s *Service) checkIfLogExist(level string, token string) bool {
	var data SpLog
	res := s.Mysql.Where(map[string]interface{}{"level": level, "access_token": token}).Find(&data).RowsAffected
	// 攻击者单次攻击记录存在则返回true 否则返回false
	if res > 0 {
		return true
	} else {
		return false
	}
}

//func (s *Service) insertReportFirst(info SpLog) {
//	s.Mysql.Create(&info)
//}
//
//func (s *Service) insertReportCount(accessToken string) {
//	var dataLog SpLog
//	var oldData SpLog
//	r := s.Mysql.Where(map[string]interface{}{"level": "report", "access_token": accessToken}).Find(&dataLog)
//	fmt.Println(r.RowsAffected)
//	if r.RowsAffected > 0 {
//		s.Mysql.Where(map[string]interface{}{"level": "report", "access_token": accessToken}).Find(&oldData)
//		s.Mysql.Model(&dataLog).Where("level = ? AND access_token = ?", "report", accessToken).Update("count", oldData.Count+1)
//	} else {
//		s.Mysql.Create(&dataLog)
//	}
//}
