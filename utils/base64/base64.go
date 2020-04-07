package base64

import (
	"encoding/base64"
	"scarletpot/utils/conf"
	"scarletpot/utils/log"
)

func Base64Decode(data string) string {
	decodeBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Err(conf.GetUserConfig().Lang.Lang, "", err)
	}
	return string(decodeBytes)
}

func Base64Encode(data string) string {
	encodeString := base64.StdEncoding.EncodeToString([]byte(data))
	return string(encodeString)
}
