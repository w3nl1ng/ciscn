package probe

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func IsGlastopfHoneyPot(target string) bool {
	request, err := http.NewRequest("GET", target, nil)
	if err != nil {
		fmt.Println("创建请求对象时发生错误:", err)
		return false
	}
	Host := strings.TrimSuffix(target, "/")

	// 设置请求头字段
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.0; Win64; x64; rv:106.0) Gecko/20100101 Firefox/106.0") // 设置自定义的 User-Agent
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	request.Header.Set("Host", strings.Split(Host, "//")[1])
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("无法发送 HTTP 请求:", err)
		return false
	}
	defer resp.Body.Close()

	// 读取响应内容
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("无法读取响应内容:", err)
		return false
	}

	// 将响应内容转换为字符串
	bodyStr := string(bodyBytes)

	// 检查响应内容是否包含特定的 HTML 字符串
	if strings.Contains(bodyStr, "<h2>Blog Comments</h2>") && strings.Contains(bodyStr, "Please post your comments for the blog") {
		return true
	}

	return false
}
