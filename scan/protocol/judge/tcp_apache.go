package judge

import "regexp"

func TcpApache(result map[string]interface{}) bool {
	var buff []byte
	buff, _ = result["banner.byte"].([]byte)
	ok, err := regexp.Match(`(?i)(apache|Server: Apache)`, buff)
	if err != nil {
		return false
	}
	if ok {
		result["protocol"] = "mysql"
		return true
	}
	return false
}
