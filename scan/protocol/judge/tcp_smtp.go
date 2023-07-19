package judge

import (
	"regexp"
)

func TcpSMTP(result map[string]interface{}) bool {
	var buff []byte
	buff, _ = result["banner.byte"].([]byte)
	ok, err := regexp.Match(`(^220[ -](.*)ESMTP|^421(.*)Service not available|^554 )`, buff)
	if err != nil {
		return false
	}
	if ok {
		result["protocol"] = "smtp"
		return true
	}
	return false
}
