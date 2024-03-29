package install

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	db "scarletpot/db/mysql"
	"scarletpot/utils"
	"scarletpot/utils/color"
	"scarletpot/utils/conf"
	"scarletpot/utils/i18n"
	"scarletpot/utils/log"
)

var config conf.UserConfig
var lang string

func Init() {
	if err := baseInstall(); err == nil {
		log.Info(lang, "install.base_success")
		//fmt.Println(color.Magenta("基础设置成功"))
	} else {
		log.Err(lang, "install.base_error", err)
		panic(err)
	}

	// Check `conf` folder exist
	if !utils.IsExist("conf") {
		err := os.Mkdir("conf", os.ModePerm)
		if err != nil {
			log.Err(lang, "install.config_mkdir_fail")
		}
	}

	// 检查配置文件是否已存在
	if !utils.IsExist("conf/user.config.toml") {
		// Create user.config.toml
		_, err := os.Create("conf/user.config.toml")
		if err != nil {
			log.Err(lang, "install.config_create_fail")
		}
	} else {
		warning()
		if utils.YnSelect() {
			log.Info(lang, "install.begin_install")
			_, err := os.Create("conf/user.config.toml")
			if err != nil {
				log.Err(lang, "install.config_create_fail")
			}
			utils.Sleep(1)
		} else {
			os.Exit(1)
		}
	}

}

func baseInstall() error {
	i18n.Print("zh-CN", "install.base_language")
	var langI int
	if err := utils.InputInt(&langI); err != nil {
		return err
	}
	switch langI {
	case 1:
		config.Lang.Lang = "zh-CN"
		lang = "zh-CN"
	case 2:
		config.Lang.Lang = "en-US"
		lang = "en-US"
	}

	return nil
}

func panelInstall() error {
	//lang := conf.GetUserConfig().Base.SystemLanguage
	i18n.Print(lang, "install.panel_port")

	var addr string
	if err := utils.InputStr(&addr); err != nil {
		return err
	}
	config.Panel.PanelAddr = addr
	return nil
}

func databaseInstall() error {
	//lang := conf.GetUserConfig().Base.SystemLanguage

	var (
		dbType int
		dbHost string
		dbUser string
		dbPass string
		dbName string
	)
	i18n.Print(lang, "install.db_type")
	if err := utils.InputInt(&dbType); err != nil {
		//panic(err)
		return err
	}
	i18n.Print(lang, "install.db_host")
	if err := utils.InputStr(&dbHost); err != nil {
		return err
	}
	i18n.Print(lang, "install.db_user")
	if err := utils.InputStr(&dbUser); err != nil {
		return err
	}
	i18n.Print(lang, "install.db_pass")
	if err := utils.InputStr(&dbPass); err != nil {
		return err
	}
	i18n.Print(lang, "install.db_name")
	if err := utils.InputStr(&dbName); err != nil {
		return err
	}
	switch dbType {
	case 1:
		config.Database.DbType = "mysql"
	case 2:
		config.Database.DbType = "sqlite"
	}
	config.Database.DbHost = dbHost
	config.Database.DbUser = dbUser
	config.Database.DbPass = dbPass
	config.Database.DbName = dbName
	if dbType == 1 {
		res, err := db.CheckMysql(dbUser, dbPass, dbHost, dbName)
		if res {
			log.Info(lang, "install.db_connect_suc")
			//TODO: 创建、初始化数据表
		} else {
			log.Err(lang, "install.db_connect_fail", err)
			_ = databaseInstall()
		}
	} else if dbType == 2  {

	}

	//fmt.Println("数据库信息：", config.Database.DbType, config.Database.DbHost, config.Database.DbUser, config.Database.DbPass, config.Database.DbName)
	return nil
}

func warning() {
	log.Warn(lang, "install.config_overwrite_warning")
	log.Warn(lang, "install.if_continue")
}

func Install() {
	banner := `

███████╗  ██████╗  █████╗  ██████╗  ██╗     ███████╗ ████████╗   ██████╗  ██████╗ ████████╗
██╔════╝ ██╔════╝ ██╔══██╗ ██╔══██╗ ██║     ██╔════╝ ╚══██╔══╝   ██╔══██╗██╔═══██╗╚══██╔══╝
███████╗ ██║      ███████║ ██████╔╝ ██║     █████╗      ██║  *** ██████╔╝██║   ██║   ██║   
╚════██║ ██║      ██╔══██║ ██╔══██╗ ██║     ██╔══╝      ██║  *** ██╔═══╝ ██║   ██║   ██║   
███████║ ╚██████╗ ██║  ██║ ██║  ██║ ███████╗███████╗    ██║  *** ██║     ╚██████╔╝   ██║   
╚══════╝  ╚═════╝ ╚═╝  ╚═╝ ╚═╝  ╚═╝ ╚══════╝╚══════╝    ╚═╝  *** ╚═╝      ╚═════╝    ╚═╝
`
	fmt.Println(color.Yellow(banner))
	fmt.Println(color.Blue("------------------------------------- ABOUT ---------------------------------------"))
	fmt.Println(color.Blue("|"), color.White(" 		  author: Annevi						 "), color.Blue("|"))
	fmt.Println(color.Blue("|"), color.White(" 		  github: https://github.com/ScarletWaf/ScarletPot	         "), color.Blue("|"))

	fmt.Println(color.Blue("-----------------------------------------------------------------------------------\n"))
	fmt.Println(color.Red("----------------------------- Scarlet Pot installer -------------------------------\n"))
	//fmt.Println(color.Green("		            ++++++++ 基础设置 ++++++++"))

	//fmt.Println(color.Green("		            ++++++++ 管理面板 ++++++++"))
	//lang := conf.GetUserConfig().Base.SystemLanguage
	Init()

	if err := panelInstall(); err == nil {
		log.Info(lang, "install.panel_success")
	} else {
		log.Err(lang, "install.panel_error", err)
		panic(err)
		return
	}
	//fmt.Println(color.Green("		            ++++++++ 数据库面板 +++++++"))

	if err := databaseInstall(); err == nil {
		log.Info(lang, "install.db_success")
	} else {
		log.Err(lang, "install.db_error ", err)
		panic(err)
	}

	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(config); err != nil {
		//fmt.Println(err)
		panic(err)
	}
	err := ioutil.WriteFile("conf/user.config.toml", []byte(buf.String()), 0777)
	if err != nil {
		panic(err)
	}
	log.Info(lang, "install.finished")

}
