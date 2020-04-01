package i18n

import (
	"fmt"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/yaml"
)

var I *i18n.I18n

func init() {
	I = i18n.New(yaml.New("locals"))
}

func I18n() *i18n.I18n {
	return I
}

func I18nStr(lang string, str string) string {
	return string(I.T(lang, str))
}

func Println(lang string, str string) {
	fmt.Println(I18n().T(lang, str))
}

func Print(lang string, str string) {
	fmt.Print(I18n().T(lang, str))
}
