package main

import (
	"scarletpot/panel"
	"scarletpot/protocal/httpservice"
	"scarletpot/protocal/mysql"
	"scarletpot/protocal/ssh"
	"scarletpot/utils/pool"
)

func main() {
	// 引导安装/初始化
	//install.Install()
	//conf.Init()
	wg, poolX := pool.New(5)
	wg.Add(2)
	poolX.Submit(func() {
		go mysql.Start()
		go ssh.Start()
		go httpservice.Start()
		go panel.Start()
	})
	wg.Wait()
	//test.Test()
	//a := conf.Get("base")
	//fmt.Println()
}
