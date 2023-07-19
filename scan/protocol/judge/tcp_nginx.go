package judge

import "regexp"

func TcpNginx(result map[string]interface{}) bool {
	var buff []byte
	buff, _ = result["banner.byte"].([]byte)
	ok, err := regexp.Match(`(?i)(nginx|Server: nginx)`, buff)
	if err != nil {
		return false
	}
	if ok {
		result["protocol"] = "mysql"
		return true
	}
	return false
}
