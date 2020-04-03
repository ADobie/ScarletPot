package request

import (
	"bytes"
	"net/http"
	"strings"
)

func PostData(url string, data string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, err
}

func PostJson(url string, data []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	return resp, err
}
