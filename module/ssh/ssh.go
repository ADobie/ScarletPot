package ssh

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"io/ioutil"
	"net/http"
	"scarletpot/report"
	"scarletpot/utils/conf"
	ipinfo "scarletpot/utils/ip"
	"scarletpot/utils/log"
	"scarletpot/utils/request"
	"strings"
)

var lang string

// TODO：使用api获取命令执行结果的时候会因为延迟的原因导致无限上报bug
// TODO: SSH服务可使用高交互蜜罐 docker
var country, city, region string

func Start() {
	lang = conf.GetUserConfig().Lang.Lang
	ssh.ListenAndServe(conf.GetBaseConfig().SSH.Addr, func(s ssh.Session) {
		term := terminal.NewTerminal(s, conf.GetBaseConfig().SSH.Prefix+" ")
		arr := strings.Split(s.RemoteAddr().String(), ":")
		ip := arr[0]
		country, city, region = ipinfo.GetPos(ip)

		report.Do("SSH", ip, "", "建立链接", country, city, region, 1)
		for {
			line, err := term.ReadLine()
			if line == "exit" {
				break
			}
			if err != nil {
				log.Err("zh-CN", "", err)
				break
			}
			if strings.Contains(line, "cd") {
				_, err := io.WriteString(s, line+": Permission denied\n")
				if err != nil {
					log.Err(lang, " ", err)
				}
			}
			//output := getResultFromApi(line).Output
			output := "error\n"

			// 上报ssh蜜罐信息
			go report.Do("SSH", arr[0], "", line, country, city, region, 1)

			_, err = io.WriteString(s, output)
			if err != nil {
				log.Err(lang, " ", err)
			}
		}
	},
		ssh.PasswordAuth(func(s ssh.Context, passwd string) bool {
			info := s.User() + " " + passwd
			arr := strings.Split(s.RemoteAddr().String(), ":")
			log.Info("zh-CN", arr[0]+" 正在尝试连接")
			report.Do("SSH", arr[0], "", info, country, city, region, 0)

			username := conf.GetBaseConfig().SSH.User
			password := conf.GetBaseConfig().SSH.Password

			if username == s.User() && password == passwd {
				report.Do("SSH", arr[0], "", "密码正确 已进入ssh", country, city, region, 1)
				return true
			}
			return false
		}),
	)
}

type CmdRes struct {
	Output string `json:"output"`
	Langid string `json:"langid"`
	Code   string `json:"code"`
	Errors string `json:"errors"`
	Time   string `json:"time"`
}

// 通过接口获取命令结果 暂时不用
func getResultFromApi(cmd string) CmdRes {
	var cmdRes CmdRes
	encodeString := base64.StdEncoding.EncodeToString([]byte("#!/bin/bash\n" + cmd))
	// 接口可能会失效，需要及时检查更新
	res, err := request.PostData("https://runcode-api2-ng.dooccn.com/compile2", "language=11&code="+encodeString+"&stdin=123%0Ahaha2%0A")
	client := &http.Client{}
	res.Header.Set("Referer", "https://www.dooccn.com/shell/")
	resp, err := client.Do(res)
	if err != nil {
		log.Err(lang, " ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Err(lang, " ", err)
	}

	err = json.Unmarshal(body, &cmdRes)
	if err != nil {
		log.Err(lang, " ", err)
	}

	return cmdRes
}
