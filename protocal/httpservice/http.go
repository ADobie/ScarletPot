package httpservice

import (
	"fmt"
	"net/http"
	"scarletpot/utils/conf"
	"scarletpot/utils/log"
)

var lang, addr, msg string

func httpHandler(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm() // 解析参数，默认是不会解析的
	//fmt.Println(r.Form)
	fmt.Print(r.URL.Path + " ")
	//fmt.Println("scheme", r.URL.Scheme)
	// 记录攻击者ip
	fmt.Println(r.RemoteAddr)
	fmt.Fprintf(w, msg) // 这个写入到 w 的是输出到客户端的
}

func Start() {
	lang = conf.GetUserConfig().Base.SystemLanguage
	addr = conf.GetBaseConfig().Http.Addr
	msg = conf.GetBaseConfig().Http.Msg
	http.HandleFunc("/", httpHandler)     // 设置访问的路由
	err := http.ListenAndServe(addr, nil) // 设置监听的端口
	if err != nil {
		log.Err(lang, "ListenAndServe", err)
	}
}
