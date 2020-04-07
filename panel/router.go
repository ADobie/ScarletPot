package panel

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) initRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(s.successJSON("Scarlet Pot"))

	})
	r.POST("/api/report", func(c *gin.Context) {
		c.JSON(s.reportHandler(c))
	})

	r.POST("/api/token/gen", func(c *gin.Context) {
		c.JSON(s.genApiToken(c))
	})

	r.GET("/api/token/sec")
	return r
}
