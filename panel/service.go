package panel

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	"scarletpot/utils/conf"
)

type Service struct {
	BaseConf conf.BaseConfig
	UserConf conf.UserConfig
	Router   *gin.Engine
	Db    *gorm.DB
	Ws       *websocket.Conn
}

func (s *Service) init() {
	s.initConfig()
	if s.UserConf.Database.DbType == "mysql" {
		s.initMysql()
	} else {
		s.initSqlite()
	}
	s.Router = s.initRouter()
	//s.getValidAttack()
	//s.dataInfo()
	//s.genApiToken()
	s.getServiceCount()
	panic(s.Router.Run(s.UserConf.Panel.PanelAddr))
}
