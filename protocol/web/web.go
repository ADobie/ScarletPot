package web

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"net/http"
	"scarletpot/report"
	"scarletpot/utils/base64"
	"scarletpot/utils/conf"
	ipinfo "scarletpot/utils/ip"
	"scarletpot/utils/log"
	"scarletpot/utils/url"
	"strings"
)

var country, city, region string

func Start() {
	initJsonp()
}

func initJsonp() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		r.LoadHTMLGlob("./protocol/web/template/jsonp/index.html")
		c.HTML(http.StatusOK, "index.html", gin.H{})
		ip := strings.Split(c.Request.RemoteAddr, ":")[0]
		country, city, region = ipinfo.GetPos(ip)

		report.Do("Jsonp", ip, "weiboJsonp", "访问", country, city, region, false)

	})

	r.POST("/api/a", func(c *gin.Context) {
		data := c.PostForm("data")
		res := url.UrlDecode(base64.Base64Decode(data))
		username := parseWbName(res)
		ip := strings.Split(c.Request.RemoteAddr, ":")[0]
		if username == "登录" {
			report.Do("Jsonp", ip, "weiboJsonp", "未登录", country, city, region, false)
			return
		}
		report.Do("Jsonp", ip, "weiboJsonp", "已登录: username: "+username, country, city, region, true)
	})
	r.Run(conf.GetBaseConfig().Web.Jsonp)
}

func parseWbName(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Err("zh-CN", "", err)
	}
	// 解析html节点 获取微博用户名
	username := doc.Find("a").Eq(4).Text()
	return username
}
