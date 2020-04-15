package main

import (
	"scarletpot/panel"
	"scarletpot/protocol/httpservice"
	"scarletpot/protocol/mysql"
	"scarletpot/protocol/redis"
	"scarletpot/protocol/ssh"
	"scarletpot/protocol/web"
	"scarletpot/utils/pool"
)

func main() {
	// 引导安装/初始化
	//install.Install()
	//conf.Init()
	wg, poolX := pool.New(6)
	wg.Add(6)
	poolX.Submit(func() {
		go mysql.Start()
		go ssh.Start()
		go httpservice.Start()
		go panel.Start()
		go redis.Start()
		go web.Start()
	})
	wg.Wait()
	//test.Test()
	//a := conf.Get("base")
	//fmt.Println()
}
