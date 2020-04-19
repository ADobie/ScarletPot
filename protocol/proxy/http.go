package proxy

import (
	"fmt"
	"net/http"
	"scarletpot/report"
	"scarletpot/utils/conf"
	"scarletpot/utils/log"
	"strings"
)

var lang, addr, msg string

func Start() {
	lang = conf.GetUserConfig().Lang.Lang
	addr = conf.GetBaseConfig().Http.Addr
	msg = conf.GetBaseConfig().Http.Msg

	http.HandleFunc("/", httpHandler)     // 设置访问的路由
	err := http.ListenAndServe(addr, nil) // 设置监听的端口
	if err != nil {
		log.Err(lang, "ListenAndServe", err)
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认是不会解析的
	var arg string
	for k, v := range r.Form {
		arg += k + "=" + strings.Join(v, "") + "&"
	}

	info := r.Method + " " + r.URL.Path + "?" + arg
	fmt.Println(info)
	remoteAddr := strings.Split(r.RemoteAddr, ":")

	go report.Do("HTTP", remoteAddr[0], "", info)
	fmt.Fprintf(w, msg)
}
