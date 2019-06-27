package utils

import "net/http"

var (
	cnUrls = []string{
		"http://189.cn/",
		"http://www.10010.com",
		"http://10086.cn",

		"https://oss-cn-hangzhou.aliyuncs.com/",
		"https://oss-cn-shanghai.aliyuncs.com",
		"https://oss-cn-qingdao.aliyuncs.com",
		"https://oss-cn-beijing.aliyuncs.com/",
		"https://oss-cn-shenzhen.aliyuncs.com/",
		"https://oss-cn-chengdu.aliyuncs.com/",
	}
)

func IsGfwed() bool {
	var failCount int
	for _, v := range cnUrls {
		resp, err := http.DefaultClient.Head(v)
		if err != nil {
			failCount++
			continue
		}
		resp.Body.Close()
	}
	if failCount > len(cnUrls)/2 {
		return true
	}
	return false
}
