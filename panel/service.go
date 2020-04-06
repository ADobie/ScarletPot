package panel

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"scarletpot/utils/conf"
)

type Service struct {
	BaseConf conf.BaseConfig
	UserConf conf.UserConfig
	Router   *gin.Engine
	Mysql    *gorm.DB
}

func (s *Service) init() {
	s.initConfig()
	s.initMysql()
	s.Router = s.initRouter()
	panic(s.Router.Run(s.UserConf.Panel.PanelAddr))
}
