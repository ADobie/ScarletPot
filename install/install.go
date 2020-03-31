package install

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"scarletpot/utils"
	"scarletpot/utils/color"
	"scarletpot/utils/log"
)

var config = map[string]interface{}{
	"rpc": map[string]string{
		"status": "val1",
		"addr":   "val2",
	},
}

func init() {
	// Check `conf` folder exist
	if !utils.IsExist("conf") {
		err := os.Mkdir("conf", os.ModePerm)
		if err != nil {
			log.Err("zh-CN", "install.config_mkdir_fail")
		}
	}

	// 检查配置文件是否已存在
	if utils.IsExist("conf/config,ini") {
		warning()
	}

	// Check `config.toml` file exist
	//if !utils.IsExist("conf/config.toml") {
	//	// Create config.toml
	//	_, err := os.Create("conf/config.toml")
	//	if err != nil {
	//		log.Err("zh-CN", "install.config_create_fail")
	//	}
	//} else {
	//	//警告 覆盖配置
	//	warning()
	//	if utils.YnSelect() {
	//		log.Succ("zh-CN", "install.begin_install")
	//		_, err := os.Create("conf/config.toml")
	//		if err != nil {
	//			log.Err("zh-CN", "install.config_create_fail")
	//		}
	//		utils.Sleep(1)
	//	} else {
	//		os.Exit(1)
	//	}
	//}

}

func baseInstall() {

	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(config); err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf.String())
}

func warning() {
	log.Warn("zh-CN", "install.config_overwrite_warning")
	log.Warn("zh-CN", "install.if_continue")
}

func Install() {
	banner := `

███████╗ ██████╗ █████╗ ██████╗ ██╗     ███████╗████████╗   ██████╗  ██████╗ ████████╗
██╔════╝██╔════╝██╔══██╗██╔══██╗██║     ██╔════╝╚══██╔══╝   ██╔══██╗██╔═══██╗╚══██╔══╝
███████╗██║     ███████║██████╔╝██║     █████╗     ██║      ██████╔╝██║   ██║   ██║   
╚════██║██║     ██╔══██║██╔══██╗██║     ██╔══╝     ██║      ██╔═══╝ ██║   ██║   ██║   
███████║╚██████╗██║  ██║██║  ██║███████╗███████╗   ██║      ██║     ╚██████╔╝   ██║   
╚══════╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚══════╝   ╚═╝      ╚═╝      ╚═════╝    ╚═╝
`
	fmt.Println(color.Yellow(banner))
	fmt.Println(color.Blue("------------------------------------- ABOUT ---------------------------------------"))
	fmt.Println(color.Blue("|"), color.White(" 		  author: Annevi						 "), color.Blue("|"))
	fmt.Println(color.Blue("|"), color.White(" 		  github: https://github.com/ScarletWaf/ScarletPot	         "), color.Blue("|"))

	fmt.Println(color.Blue("-----------------------------------------------------------------------------------\n"))
	fmt.Println(color.Red("----------------------------- Scarlet Pot installer -------------------------------"))
	baseInstall()
}
