package panel

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"scarletpot/utils/log"
	"time"
)

/*
+----------------------+
| 		数据库结构	   |
+----------------------+
| sp_admin             |
| sp_info              |
| sp_user              |
+----------------------+
*/

/*
sp_admin 表结构
+-----------+--------------+------+-----+---------+----------------+
| Field     | Type         | Null | Key | Default | Extra          |
+-----------+--------------+------+-----+---------+----------------+
| id        | int(11)      | NO   | PRI | NULL    | auto_increment |
| name      | varchar(255) | YES  |     | NULL    |                |
| pass      | varchar(255) | YES  |     | NULL    |                |
| token     | varchar(255) | YES  |     | NULL    |                |
| lastLogin | datetime     | YES  |     | NULL    |                |
| lastIp    | varchar(255) | YES  |     | NULL    |                |
+-----------+--------------+------+-----+---------+----------------+
*/
type SpAdmin struct {
	ID        int       `gorm:"primary_key;column:id;type:int(11);not null"`
	Name      string    `gorm:"column:name;type:varchar(255)"`
	Pass      string    `gorm:"column:pass;type:varchar(255)"`
	Token     string    `gorm:"column:token;type:varchar(255)"`
	LastLogin time.Time `gorm:"column:lastLogin;type:datetime"`
	LastIP    string    `gorm:"column:lastIp;type:varchar(255)"`
}

/*
+-------------+--------------+------+-----+---------+----------------+
| Field       | Type         | Null | Key | Default | Extra          |
+-------------+--------------+------+-----+---------+----------------+
| id          | int(11)      | NO   | PRI | NULL    | auto_increment |
| type        | varchar(255) | YES  |     | NULL    |                |
| webApp      | varchar(255) | YES  |     | NULL    |                |
| info        | text         | YES  |     | NULL    |                |
| time        | datetime     | YES  |     | NULL    |                |
| attackIp    | varchar(255) | YES  |     | NULL    |                |
| clientIp    | varchar(255) | YES  |     | NULL    |                |
| accessToken | varchar(255) | YES  |     | NULL    |                |
+-------------+--------------+------+-----+---------+----------------+
*/

type SpInfo struct {
	ID          int       `gorm:"primary_key;column:id;type:int(11);not null"`
	Type        string    `gorm:"column:type;type:varchar(255)"`
	WebApp      string    `gorm:"column:webApp;type:varchar(255)"`
	Info        string    `gorm:"column:info;type:text"`
	Time        time.Time `gorm:"column:time;type:datetime"`
	AttackIP    string    `gorm:"column:attackIp;type:varchar(255)"`
	ClientIP    string    `gorm:"column:clientIp;type:varchar(255)"`
	AccessToken string    `gorm:"column:accessToken;type:varchar(255)"`
}

/*
+-----------+--------------+------+-----+---------+-------+
| Field     | Type         | Null | Key | Default | Extra |
+-----------+--------------+------+-----+---------+-------+
| id        | int(11)      | YES  | PRI | NULL    |       |
| username  | varchar(255) | YES  |     | NULL    |       |
| password  | varchar(255) | YES  |     | NULL    |       |
| apiKey    | varchar(255) | YES  |     | NULL    |       |
| apiSecret | varchar(255) | YES  |     | NULL    |       |
| lastLogin | datetime     | YES  |     | NULL    |       |
+-----------+--------------+------+-----+---------+-------+
*/

type SpUser struct {
	ID        int       `gorm:"primary_key;column:id;type:int(11)"`
	Username  string    `gorm:"column:username;type:varchar(255)"`
	Password  string    `gorm:"column:password;type:varchar(255)"`
	APIKey    string    `gorm:"column:apiKey;type:varchar(255)"`
	APISecret string    `gorm:"column:apiSecret;type:varchar(255)"`
	LastLogin time.Time `gorm:"column:lastLogin;type:datetime"`
}

func (s *Service) initMysql() {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local&charset=utf8mb4,utf8",
		s.UserConf.Database.DbUser,
		s.UserConf.Database.DbPass,
		s.UserConf.Database.DbHost,
		s.UserConf.Database.DbName,
	))
	if err != nil {
		log.Err(s.UserConf.Lang.Lang, "", err)
	}

	s.Mysql = db

	// 创建表自动迁移
	s.Mysql.AutoMigrate(&SpAdmin{}, &SpInfo{}, &SpUser{})
}
