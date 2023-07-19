package judge

import (
	"regexp"
)

func TcpMysql(result map[string]interface{}) bool {
	var buff []byte
	buff, _ = result["banner.byte"].([]byte)
	ok, err := regexp.Match(`(mysql_native_password|MySQL server|MariaDB server|mysqladmin flush-hosts)`, buff)
	if err != nil {
		return false
	}
	if ok {
		result["protocol"] = "mysql"
		return true
	}
	return false
}
