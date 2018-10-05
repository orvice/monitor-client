package utils

import "net/http"

var (
	cnUrls = []string{
		"http://189.cn/",
		"http://www.10010.com",
		"http://10086.cn",
	}
)

func IsGfwed() bool {
	var failCount int
	for _, v := range cnUrls {
		resp, err := http.DefaultClient.Get(v)
		if err != nil {
			failCount++
		}
		resp.Body.Close()
	}
	if failCount > len(cnUrls)/2 {
		return true
	}
	return false
}
