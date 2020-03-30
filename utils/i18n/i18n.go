package i18n

import (
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
