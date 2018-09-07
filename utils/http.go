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
		_, err := http.DefaultClient.Get(v)
		if err != nil {
			failCount++
		}
	}
	if failCount > len(cnUrls)/2 {
		return true
	}
	return false
}