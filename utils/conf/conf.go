package conf

import (
	"github.com/BurntSushi/toml"
	"log"
)

type BaseConfig struct {
	Logs struct {
		ErrorLog string `toml:"errorLog"`
		InfoLog  string `toml:"infoLog"`
		DebugLog string `toml:"debugLog"`
	} `toml:"logs"`

	API struct {
		APIKeyR     string `toml:"apiKeyR"`
		APIQSecretR string `toml:"apiQSecretR"`
		APIKeyQ     string `toml:"apiKeyQ"`
		APISecretQ  string `toml:"apiSecretQ"`
	} `toml:"api"`
	Mysql struct {
		File string `toml:"file"`
		Addr string `toml:"addr"`
	} `toml:"mysql"`
	SSH struct {
		Addr     string `toml:"addr"`
		Prefix   string `toml:"prefix"`
		User     string `toml:"user"`
		Password string `toml:"password"`
	} `toml:"ssh"`
	Http struct {
		Addr string `toml:"addr"`
		Msg  string `toml:"msg"`
	} `toml:"http"`
	Redis struct {
		Addr string `toml:"addr"`
	} `toml:"redis"`
	Web struct {
		Jsonp string `toml:"jsonp"`
	} `toml:"web"`
}

type UserConfig struct {
	Lang struct {
		Lang string `toml:"lang"`
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

func GetBaseConfig() BaseConfig {
	var cfg BaseConfig
	var cfgPath string = "conf/base.config.toml"
	if _, err := toml.DecodeFile(cfgPath, &cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
