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
	"scarletpot/utils/log"
	"scarletpot/utils/request"
	"strings"
)

var lang string

// TODO：使用api获取命令执行结果的时候会因为延迟的原因导致无限上报bug
// TODO: SSH服务可使用高交互蜜罐 docker

func Start() {
	lang = conf.GetUserConfig().Lang.Lang
	errT := ssh.ListenAndServe(conf.GetBaseConfig().SSH.Addr, func(s ssh.Session) {
		term := terminal.NewTerminal(s, conf.GetBaseConfig().SSH.Prefix+" ")
		arr := strings.Split(s.RemoteAddr().String(), ":")
		report.Do("SSH", arr[0], "", "建立链接")

		line := ""
		for {
			line, _ = term.ReadLine()
			if line == "exit" {
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
			go report.Do("SSH", arr[0], "", line)

			_, err := io.WriteString(s, output)
			if err != nil {
				log.Err(lang, " ", err)
			}
		}
	})
	if errT != nil {
		log.Err(lang, " ", errT)
	}
}

type CmdRes struct {
	Output string `json:"output"`
	Langid string `json:"langid"`
	Code   string `json:"code"`
	Errors string `json:"errors"`
	Time   string `json:"time"`
}

// 通过接口获取命令结果
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
