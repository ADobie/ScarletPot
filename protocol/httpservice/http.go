package httpservice

import (
	"fmt"
	"net/http"
	"scarletpot/report"
	"scarletpot/utils/conf"
	"scarletpot/utils/log"
	"strings"
)

var lang, addr, msg string

func httpHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认是不会解析的
	//fmt.Println(r.Form)
	var arg string
	for k, v := range r.Form {
		//fmt.Println("key:", k)
		//fmt.Println("val:", strings.Join(v, ""))
		arg += k + "=" + strings.Join(v, "") + "&"
	}
	//fmt.Println(r.Method + " " + r.URL.Path + "?" + arg)
	info := r.Method + " " + r.URL.Path + "?" + arg
	fmt.Println(info)
	remoteAddr := strings.Split(r.RemoteAddr, ":")

	go report.Do("HTTP", remoteAddr[0], "", info)
	//fmt.Print(r.URL.Path + " ")
	// 记录攻击者ip
	//fmt.Println(r.RemoteAddr)
	//fmt.Println(remoteAddr[0])
	fmt.Fprintf(w, msg)
}

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
