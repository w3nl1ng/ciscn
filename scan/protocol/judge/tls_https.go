package judge

import (
	"fmt"
	"net/url"
	"regexp"
)

func TlsHTTPS(result map[string]interface{}, Args map[string]interface{}) bool {
	var buff []byte
	buff, _ = result["banner.byte"].([]byte)
	ok, err := regexp.Match(`^HTTP/\d.\d \d*`, buff)
	if err != nil {
		return false
	}
	if ok {
		result["protocol"] = "https"
		httpResult, httpErr := httpIdentifyResult(result, Args)
		if httpErr != nil {
			result["banner.string"] = "None"
			return true
		}
		result["banner.string"] = httpResult["http.title"].(string)
		u, err := url.Parse(httpResult["http.target"].(string))
		if err != nil {
			result["path"] = ""
		} else {
			result["path"] = u.Path
		}
		r := httpResult["http.result"].(string)
		c := fmt.Sprintf("[%s]", httpResult["http.code"].(string))
		if len(r) != 0 {
			result["identify.bool"] = true
			result["identify.string"] = fmt.Sprintf("%s %s", c, r)
			result["note"] = httpResult["http.target"].(string)
			return true
		} else {
			result["identify.bool"] = true
			result["identify.string"] = c
			result["note"] = httpResult["http.target"].(string)
			return true
		}
	}
	return false
}
