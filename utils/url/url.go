package url

import (
	"net/url"
)

func UrlEncode(str string) string {
	escapeUrl := url.QueryEscape(str)
	return escapeUrl
}

func UrlDecode(str string) string {
	enEscapeUrl, _ := url.QueryUnescape(str)
	return enEscapeUrl
}
