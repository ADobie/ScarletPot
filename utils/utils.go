package utils

import (
	"fmt"
	"os"
	"time"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func Sleep(Sec time.Duration) {
	time.Sleep(Sec * time.Second)
}

// 命令行是否 选项 默认为n
func YnSelect() bool {
	var selection string
	_, err := fmt.Scan(&selection)
	if err != nil {
		panic(err)
	}
	if selection == "y" || selection == "Y" {
		return true
	} else if selection == "N" || selection == "n" {
		return false
	}
	return false
}
