package main

import (
	"fmt"
	"os"
	"scarletpot/install"
	"scarletpot/panel"
	"scarletpot/protocol/mysql"
	"scarletpot/protocol/proxy"
	"scarletpot/protocol/redis"
	"scarletpot/protocol/ssh"
	"scarletpot/protocol/web"
	"scarletpot/utils/help"
	"scarletpot/utils/pool"
)

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		help.Info()
	} else {
		if args[1] == "help" {
			help.Info()
		} else if args[1] == "install" {
			install.Install()
		} else if args[1] == "version" {
			fmt.Println("ScarletPot v0.1 2020.4.19\nBy Annevi")
		} else if args[1] == "run" {
			wg, poolX := pool.New(6)
			wg.Add(6)
			poolX.Submit(func() {
				go mysql.Start()
				go ssh.Start()
				go proxy.Start()
				go panel.Start()
				go redis.Start()
				go web.Start()
			})
			wg.Wait()
		} else {
			help.Info()
		}
	}
}
