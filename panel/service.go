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
	Mysql    *gorm.DB
	Ws       *websocket.Conn
}

func (s *Service) init() {
	s.initConfig()
	s.initMysql()
	s.Router = s.initRouter()
	//s.getValidAttack()
	//s.dataInfo()
	//s.genApiToken()
	s.getDayAverageCount()
	panic(s.Router.Run(s.UserConf.Panel.PanelAddr))
}
