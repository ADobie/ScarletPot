package conf

import (
	"github.com/BurntSushi/toml"
	"log"
)

type UserConfig struct {
	Base struct {
		SystemLanguage string `toml:"systemLanguage"`
	} `toml:"base"`
	Panel struct {
		PanelAddr string `toml:"panelAddr"`
	} `toml:"panel"`
	Database struct {
		DbType string `toml:"dbType"`
		DbHost string `toml:"dbHost"`
		DbUser string `toml:"dbUser"`
		DbPass string `toml:"dbPass"`
		DbName string `toml:"dbName"`
	} `toml:"database"`
}

func GetUserConfig() UserConfig {
	var ucg UserConfig
	var ucPath string = "conf/user.config.toml"
	if _, err := toml.DecodeFile(ucPath, &ucg); err != nil {
		log.Fatal(err)
	}
	return ucg
}
