package judge

import (
	"regexp"
)

func TcpIMAP(result map[string]interface{}) bool {
	var buff []byte
	buff, _ = result["banner.byte"].([]byte)
	ok, err := regexp.Match(`^* OK`, buff)
	if err != nil {
		return false
	}
	if ok {
		result["protocol"] = "imap"
		return true
	}
	return false
}
