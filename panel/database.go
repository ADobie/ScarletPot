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
| 		蜜罐数据库结构   |
+----------------------+
| sp_admin             |
| sp_info              |
| sp_user              |
+----------------------+

- 2020.4.24 数据库结构更新
[sp_info]
err次数  err_count (用于统计上报成功量)
攻击ip地区 ip_area
是否在境内 is_ip_inner

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
	gorm.Model
	//ID        uint      `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string    `gorm:"column:name;type:varchar(255)"`
	Pass      string    `gorm:"column:pass;type:varchar(255)"`
	Token     string    `gorm:"column:token;type:varchar(255)"`
	LastLogin time.Time `gorm:"column:last_login;type:datetime"`
	LastIP    string    `gorm:"column:last_ip;type:varchar(255)"`
	//CreatedAt time.Time
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
	gorm.Model
	Type        string `json:"type" gorm:"column:type;type:varchar(255)"`
	WebApp      string `json:"webApp" gorm:"column:web_app;type:varchar(255)"`
	Info        string `json:"detail" gorm:"column:info;type:longtext"`
	AttackIP    string `json:"attackIp" gorm:"column:attack_ip;type:varchar(255)"`
	Country     string `gorm:"column:country;type:varchar(255)"`
	City        string `gorm:"column:city;type:varchar(255)"`
	Region      string `json:"region" gorm:"column:region"`
	ClientIP    string `gorm:"column:client_ip;type:varchar(255)"`
	AccessToken string `json:"accessToken" gorm:"column:access_token;type:varchar(255)"`
	Count       uint   `gorm:"column:count;type:int;default:0"`
	Valid       uint   `json:"valid" gorm:"column:valid;type:int;default:0"`
	Invalid     uint   `gorm:"column:invalid;type:int;default:0"`
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
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT"`
	Username  string    `gorm:"column:username;type:varchar(255)"`
	Password  string    `gorm:"column:password;type:varchar(255)"`
	APIId     string    `gorm:"column:api_id;type:varchar(255)"`
	APISecret string    `gorm:"column:api_secret;type:varchar(255)"`
	LastLogin time.Time `gorm:"column:last_login;type:datetime"`
	CreatedAt time.Time
}

/*
	日志表 记录上报情况
*/
type SpLog struct {
	gorm.Model
	Level       string `json:"level" gorm:"column:level;type:varchar(255)"`
	ClientIP    string `gorm:"column:client_ip;type:varchar(255)"`
	Info        string `gorm:"column:info;type:text"`
	AccessToken string `json:"accessToken" gorm:"column:access_token;type:varchar(255)"`
	Count       uint   `gorm:"column:count;type:int;default:0"`
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
	s.Mysql.AutoMigrate(&SpAdmin{}, &SpInfo{}, &SpUser{}, &SpLog{})
}
