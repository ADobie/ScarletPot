package panel

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"time"
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
*/

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

	for {
		_, _, err := ws.ReadMessage()
		s.wsSend(s.getAttackCount())
		if err != nil {
			// 客户端断开
			connClient[ws] = false
			break
		}
	}
}

func (s *Service) wsSend(data string) {
	//var k websocket.Conn
	for k, v := range connClient {
		if v {
			err := k.WriteJSON(data)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}
	time.Sleep(1 * time.Second)
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
		err = ws.WriteMessage(mt, []byte(s.getAttackCount()))
		if err != nil {
			break
		}
	}
}

func (s *Service) getAttackCount() string {
	rows, err := s.Mysql.Table("sp_infos").Select("sum(count) AS total").Rows()
	if err != nil {
		fmt.Println("报错1 %v", err)
	}
	defer rows.Close()
	if rows.Next() {
		total := 0
		err := rows.Scan(&total)
		if err != nil {
			fmt.Println(err) //return 0, err
		}
		return strconv.Itoa(total)
	}
	return ""
}
