package panel

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

/* 控制查询蜜罐数据
考虑使用websocket实现大屏数据展示
1.攻击总数 attackCount 现有数据计算得
2.攻击成功数据（计算攻击失败次数）attackSuccess (attackFailed = attackCount-attackSuccess)
3.境内攻击数  (attackInner)
4.境外攻击数  (attackOuter)
5.攻击区域坐标及次数  (x,y,count)
6.每日各服务攻击数据 (date,service,count)
7.各服务总计攻击数据
8.各服务实时攻击数据列表，返回最新15条
9.所有攻击数据，分页显示，每页20条
{
	"data": {
		"reportCount": 123,
		"errCount": 123,
		"validAttack": 123,
		"totalAttack": 123,
		"dayAttack": 123,
		"sshCount": 123,
		"mysqlCount": 123,
		"telnetCount": 123,
		"proxyCount": 123,
		"redisCount": 123,
		"webCount": 123
	},

	"status": {
		"cpu": 10,
		"mem": 10,
		"disk": 10
	},

	"service": {
		"ssh": true,
		"telnet": true,
		"proxy": true,
		"mysql": true,
		"redis": true,
		"web": true
	},

	"list": {
		"type": "ssh",
		"country": "中国",
		"city": "杭州",
		"ip": "127.0.0.1",
		"infoBrief": "root root...",
		"time": "2020-1-1 11:23:32"
	}
}
*/

// json数据格式
type ScreenInfo struct {
	Data struct {
		ReportCount   uint    `json:"reportCount"`
		ErrCount      int     `json:"errCount"`
		ValidAttack   int     `json:"validCount"`
		InvalidAttack int     `json:"invalidCount"`
		DayAttack     float64 `json:"dayCount"`
		SSHCount      int     `json:"sshCount"`
		MysqlCount    int     `json:"mysqlCount"`
		TelnetCount   int     `json:"telnetCount"`
		ProxyCount    int     `json:"proxyCount"`
		RedisCount    int     `json:"redisCount"`
		WebCount      int     `json:"webCount"`
	} `json:"data"`
	Status struct {
		CPU  int `json:"cpu"`
		Mem  int `json:"mem"`
		Disk int `json:"disk"`
	} `json:"status"`
	Service struct {
		SSH    int `json:"ssh"`
		Telnet int `json:"telnet"`
		Proxy  int `json:"proxy"`
		Mysql  int `json:"mysql"`
		Redis  int `json:"redis"`
		Web    int `json:"web"`
	} `json:"service"`
	List struct {
		Type      string `json:"type"`
		Country   string `json:"country"`
		City      string `json:"city"`
		IP        string `json:"ip"`
		InfoBrief string `json:"infoBrief"`
		Time      string `json:"time"`
	} `json:"list"`
}

var connClient = make(map[*websocket.Conn]bool)

// 去除跨域限制
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// TODO: 首先完成大屏数据的基本展示，控制台优先级放低
func (s *Service) wsHandler(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// 创建 WebSocket 失败
		return
	}
	connClient[ws] = true
	defer ws.Close()
	//fmt.Println(s.dataInfo())
	s.wsSend(s.dataInfo())

	for {
		_, _, err := ws.ReadMessage()
		s.wsSend(s.dataInfo())
		if err != nil {
			// 客户端断开
			connClient[ws] = false
			break
		}
	}
}

func (s *Service) wsSend(data []byte) {
	//fmt.Println(data)
	for k, v := range connClient {
		if v {
			//j := gin.H{"data": data}
			err := k.WriteMessage(1, data)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}
	//time.Sleep(1 * time.Second)
}

// 检查ws连接
func (s *Service) ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		//写入ws数据
		err = ws.WriteMessage(mt, []byte(s.dataInfo()))
		if err != nil {
			break
		}
	}
}

// 获取总上报次数
func (s *Service) getReportCount() uint {
	//rows, err := s.Mysql.Table("sp_infos").Select("sum(count) AS total").Rows()
	//if err != nil {
	//	fmt.Println("报错1 %v", err)
	//}
	//defer rows.Close()
	//if rows.Next() {
	//	total := 0
	//	err := rows.Scan(&total)
	//	if err != nil {
	//		fmt.Println(err) //return 0, err
	//	}
	//	return total
	//}
	//return 0
	var spLogs SpLog
	s.Mysql.Where(map[string]interface{}{"level": "report"}).Find(&spLogs)
	//fmt.Println(spLogs.Count)
	return spLogs.Count
	//s.Mysql.Where()
}

// 获取上报错误次数
func (s *Service) getErrCount() int {
	//rows, err := s.Mysql.Table("sp_logs").Select("sum(count) AS total").Rows()
	//if err != nil {
	//	fmt.Println("报错1 %v", err)
	//}
	//defer rows.Close()
	//if rows.Next() {
	//	total := 0
	//	err := rows.Scan(&total)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	return total
	//}
	var spLogs SpLog
	s.Mysql.Where(map[string]interface{}{"level": "error"}).Find(&spLogs)
	//fmt.Println(spLogs.Count)
	return int(spLogs.Count)
	//return 0
}

// 获取有效攻击数据
func (s *Service) getValidAttack() int {
	rows, err := s.Mysql.Table("sp_infos").Select("sum(valid) AS total").Rows()
	if err != nil {
		fmt.Println("报错1 %v", err)
	}
	defer rows.Close()
	if rows.Next() {
		total := 0
		err := rows.Scan(&total)
		if err != nil {
			fmt.Println(err)
		}
		return total
	}
	return 0
}

// 获取无效攻击次数
func (s *Service) getInvalidAttack() int {
	rows, err := s.Mysql.Table("sp_infos").Select("sum(invalid) AS total").Rows()
	if err != nil {
		fmt.Println("报错1 %v", err)
	}
	defer rows.Close()
	if rows.Next() {
		total := 0
		err := rows.Scan(&total)
		if err != nil {
			fmt.Println(err)
		}
		return total
	}
	return 0
}

func (s *Service) getDayAverageCount() float64 {
	var count int

	/*
		SELECT
			DATE_FORMAT( created_at, '%Y-%m-%d' ) AS date,
			sum( count ) AS count
		FROM
			`sp_infos`
		WHERE
			( `created_at` > '2018-08-30' AND `created_at` < '2020-09-06' )
		GROUP BY
			date
		ORDER BY
			date ASC
	*/
	s.Mysql.Table("sp_infos").
		Select("DATE_FORMAT(created_at, '%Y-%m-%d') AS date,sum( count ) AS day_attack").
		Where("`created_at` > '2020-04-20' AND `created_at` < '2100-09-06'").
		Group("DATE_FORMAT(created_at, '%Y-%m-%d')").
		Order("DATE_FORMAT(created_at, '%Y-%m-%d') asc").Count(&count)

	rows, err := s.Mysql.Table("sp_infos").
		Select("DATE_FORMAT( created_at, '%Y-%m-%d' ) AS date,sum( count ) AS day_attack").
		Where("`created_at` > '2020-04-20' AND `created_at` < '2100-09-06'").
		Group("date").
		Order("date asc").Rows()
	if err != nil {
		panic(err)
	}

	var sum float64
	for rows.Next() {
		var date string
		dayAttack := 0
		err := rows.Scan(&date, &dayAttack)
		sum += float64(dayAttack)
		if err != nil {
			fmt.Println(err)
		}
	}
	average := sum / float64(count)
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", average), 64)
	return value
}

// 各服务受攻击次数
func (s *Service) getServiceCount() map[string]int {
	rows, err := s.Mysql.Table("sp_infos").Select("type as service, sum(count) as total").Group("service").Order("service desc").Rows()
	if err != nil {
		panic(err)
	}
	ret := make(map[string]int, 0)

	for rows.Next() {
		var service string
		var count int
		err := rows.Scan(&service, &count)
		ret[service] = count
		if err != nil {
			panic(err)
		}
	}
	return ret
}

// 构建前端数据json
func (s *Service) dataInfo() []byte {
	var data ScreenInfo
	serviceCount := s.getServiceCount()

	data.Data.ReportCount = s.getReportCount()
	data.Data.ErrCount = s.getErrCount()
	data.Data.ValidAttack = s.getValidAttack()
	data.Data.InvalidAttack = s.getInvalidAttack()
	data.Data.DayAttack = s.getDayAverageCount()
	data.Service.Mysql = serviceCount["MySQL"]
	data.Service.Proxy = serviceCount["HTTP"]
	data.Service.SSH = serviceCount["SSH"]

	jsonInfo, _ := json.Marshal(&data)
	fmt.Println(string(jsonInfo))
	return jsonInfo
}

/*
	SELECT
		type AS service,
		sum( count ) AS total
	FROM
		`sp_infos`
	GROUP BY
		service
*/
