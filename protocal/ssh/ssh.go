package ssh

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"scarletpot/utils/request"
)

// 虚假 ssh 服务
// 方案一: 通过获取攻击者输入的指定命令
// 方案二：寻找在线linux命令运行接口： https://tool.runoob.com/compile.php 通过接口获取命令结果并返回
// https://runcode-api2-ng.dooccn.com/compile2 [POST] language=11&code=IyEvYmluL2Jhc2gKZWNobyAnaGksVzN4dWUhJw==&stdin=
// https://runcode-api2-ng.dooccn.com/compile2
func Start() {
	getResultFromApi()
	//log.Fatal(ssh.ListenAndServe(":2222", func(s ssh.Session) {
	//	term := terminal.NewTerminal(s, "[root@ubuntu ~]# ")
	//	line := ""
	//	for {
	//		line, _ = term.ReadLine()
	//		if line == "exit" {
	//			break
	//		}
	//
	//		//fileName := "whoami"
	//		output := "test"
	//		io.WriteString(s, output+"\n")
	//	}
	//}))
}

// 通过接口获取命令结果
func getResultFromApi() {
	res, err := request.PostData("https://runcode-api2-ng.dooccn.com/compile2", "language=11&code=IyEvYmluL2Jhc2gKI2VjaG8gaGkKbHM%3D&stdin=123%0Ahaha2%0A")
	client := &http.Client{}
	res.Header.Set("Referer", "https://www.dooccn.com/shell/")
	resp, err := client.Do(res)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
