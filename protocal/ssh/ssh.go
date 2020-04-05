package ssh

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"scarletpot/report"
	"scarletpot/utils/request"
	"strings"
)

// 虚假 ssh 服务
// 寻找在线linux命令运行接口： https://tool.runoob.com/compile.php 通过接口获取命令结果并返回
// https://runcode-api2-ng.dooccn.com/compile2 [POST] language=11&code=IyEvYmluL2Jhc2gKZWNobyAnaGksVzN4dWUhJw==&stdin=
// https://runcode-api2-ng.dooccn.com/compile2
func Start() {
	log.Fatal(ssh.ListenAndServe(":2222", func(s ssh.Session) {
		term := terminal.NewTerminal(s, "[root@ubuntu ~]# ")
		line := ""
		for {
			line, _ = term.ReadLine()
			if line == "exit" {
				break
			}
			if strings.Contains(line, "cd") {
				io.WriteString(s, line+": Permission denied\n")
			}
			output := getResultFromApi(line).Output
			report.ReportSSH("SSH", s.RemoteAddr().String(), "", line)
			io.WriteString(s, output)
		}
	}))
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
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &cmdRes)
	if err != nil {
		panic(err)
	}

	return cmdRes
}
