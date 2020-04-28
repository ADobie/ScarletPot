package ipinfo

import (
	"github.com/ipipdotnet/ipdb-go"
	"log"
)

// 调用ipip本地数据库查询ip位置
func GetPos(ip string) (string, string, string) {
	db, err := ipdb.NewCity("db/ipdb/ipip.ipdb")
	if err != nil {
		log.Fatal(err)
	}

	//db.Reload("/path/to/city.ipv4.ipdb")
	info, _ := db.FindMap(ip, "CN") // return map[string]string
	return info["country_name"], info["city_name"], info["region_name"]
}
