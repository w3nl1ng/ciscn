package judge

import (
	"regexp"
)

func TlsRedisSsl(result map[string]interface{}) bool {
	var buff []byte
	buff, _ = result["banner.byte"].([]byte)
	ok, err := regexp.Match(`(^-ERR(.*)command|^-(.*).Redis)`, buff)
	if err != nil {
		return false
	}
	if ok {
		result["protocol"] = "redis-ssl"
		return true
	}
	return false
}
