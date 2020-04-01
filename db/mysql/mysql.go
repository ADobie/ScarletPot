package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	//"scarletpot/utils/conf"
)

func CheckMysql(user string, pass string, host string, name string) bool {
	//_ := conf.GetUserConfig().Database
	_, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local&charset=utf8mb4,utf8",
		user,
		pass,
		host,
		name))
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}
