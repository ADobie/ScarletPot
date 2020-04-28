package panel

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) initRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(s.successJSON("Scarlet"))

	})
	r.POST("/api/report", func(c *gin.Context) {
		c.JSON(s.reportHandler(c))
	})

	r.POST("/api/log", func(c *gin.Context) {
		c.JSON(s.logsHandler(c))
	})

	r.POST("/api/token/gen", func(c *gin.Context) {
		c.JSON(s.genApiToken(c))
	})

	r.GET("/api/token/sec")
	//r.POST("/api/getAttackList", func(c *gin.Context) {
	//	mod := c.Query("mod")
	//	s.dataHandler(mod, c)
	//})

	r.GET("/ping", func(c *gin.Context) {
		s.ping(c)
	})

	r.GET("/ws", func(c *gin.Context) {
		s.wsHandler(c)
	})

	//r.GET("/dataInit", func(c *gin.Context) {
	//	c.String(200, s.dataInfo())
	//})

	return r

	// 获取攻击列表信息
}
