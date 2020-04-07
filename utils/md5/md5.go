package md5

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string) string {
	md5s := []byte(str)
	has := md5.Sum(md5s)
	ext := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return ext
}
