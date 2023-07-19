package judge

import (
	"regexp"
)

func TcpVNC(result map[string]interface{}) bool {
	var buff []byte
	buff, _ = result["banner.byte"].([]byte)
	ok, err := regexp.Match(`^RFB \d`, buff)
	if err != nil {
		return false
	}
	if ok {
		result["protocol"] = "vnc"
		return true
	}
	return false
}
