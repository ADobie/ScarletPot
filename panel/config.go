package panel

import (
	"scarletpot/utils/conf"
)

func (s *Service) initConfig() {
	s.BaseConf = conf.GetBaseConfig()
	s.UserConf = conf.GetUserConfig()
}
