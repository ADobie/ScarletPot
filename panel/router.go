package panel

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) initRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Nothing here")
	})
	r.POST("/api/report", func(c *gin.Context) {
		s.reportHandler(c)
	})
	//
	//r.POST("/scannerStart", func(c *gin.Context) {
	//
	//})
	//
	//r.GET("/getLastTime", func(c *gin.Context) {
	//
	//})

	return r
}
