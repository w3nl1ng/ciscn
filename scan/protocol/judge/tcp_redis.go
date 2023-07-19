package judge

import (
	"regexp"
)

func TcpRedis(result map[string]interface{}) bool {
	var buff []byte
	buff, _ = result["banner.byte"].([]byte)
	ok, err := regexp.Match(`(^-ERR(.*)command|^-DENIED.Redis)`, buff)
	if err != nil {
		return false
	}
	if ok {
		result["protocol"] = "redis"
		return true
	}
	return false
}
