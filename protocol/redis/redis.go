package redis

import (
	"bufio"
	"net"
	"scarletpot/report"
	"scarletpot/utils/log"
	"scarletpot/utils/pool"
	"strconv"
	"strings"
	"time"
)

var redData map[string]string
var ip string

func Start() {
	redData = make(map[string]string)

	//建立socket，监听端口
	netListen, _ := net.Listen("tcp", "0.0.0.0:63790")

	defer netListen.Close()

	wg, poolX := pool.New(10)
	defer poolX.Release()

	for {
		wg.Add(1)
		poolX.Submit(func() {
			time.Sleep(time.Second * 2)

			conn, err := netListen.Accept()

			if err != nil {
				log.Err("zh-CN", "Redis 连接失败", err)
			}

			ip = strings.Split(conn.RemoteAddr().String(), ":")[0]
			report.Do("Redis", ip, "", "建立链接")

			log.Info("zh-CN", "Redis "+ip+" 已经连接")

			go handleConnection(conn)

			wg.Done()
		})
	}
}

func handleConnection(conn net.Conn) {
	for {
		str := parseResp(conn)

		switch value := str.(type) {
		case string:
			if len(value) == 0 {
				goto end
			}

			conn.Write([]byte(value))
		case []string:
			if value[0] == "SET" || value[0] == "set" {
				// 模拟 redis set

				// 捕获数组越界异常
				defer func() {
					err := recover()
					if err != nil {
						log.Err("zh-CN", "", err)
					}
				}()

				key := string(value[1])
				val := string(value[2])
				redData[key] = val
				go report.Do("Redis", ip, "", value[0]+" "+value[1]+" "+value[2])

				conn.Write([]byte("+OK\r\n"))
			} else if value[0] == "GET" || value[0] == "get" {
				err := func() {
					// 模拟 redis get
					key := string(value[1])
					val := string(redData[key])

					valLen := strconv.Itoa(len(val))
					str := "$" + valLen + "\r\n" + val + "\r\n"
					go report.Do("Redis", ip, "", value[0]+" "+value[1])
					conn.Write([]byte(str))
				}
				if err != nil {
					conn.Write([]byte("+OK\r\n"))
				}
			} else {
				err := func() {
					go report.Do("Redis", ip, "", value[0]+" "+value[1])
				}
				if err != nil {
					go report.Do("Redis", ip, "", value[0])
				}
				conn.Write([]byte("+OK\r\n"))
			}
			break
		default:

		}
	}
end:
	conn.Close()
}

// 解析 Redis 协议
func parseResp(conn net.Conn) interface{} {
	r := bufio.NewReader(conn)
	line, err := r.ReadString('\n')
	if err != nil {
		return ""
	}

	cmdType := string(line[0])
	cmdTxt := strings.Trim(string(line[1:]), "\r\n")

	switch cmdType {
	case "*":
		count, _ := strconv.Atoi(cmdTxt)
		var data []string
		for i := 0; i < count; i++ {
			line, _ := r.ReadString('\n')
			cmd_txt := strings.Trim(string(line[1:]), "\r\n")
			c, _ := strconv.Atoi(cmd_txt)
			length := c + 2
			str := ""
			for length > 0 {
				block, _ := r.Peek(length)
				if length != len(block) {

				}
				r.Discard(length)
				str += string(block)
				length -= len(block)
			}

			data = append(data, strings.Trim(str, "\r\n"))
		}
		return data
	default:
		return cmdTxt
	}
}
