package judge

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
)

func TcpActiveMQ(result map[string]interface{}) bool {
	var buff []byte
	buff, _ = result["banner.byte"].([]byte)
	ok, err := regexp.Match(`ActiveMQ`, buff)
	if err != nil {
		return false
	}
	if ok {
		ver, err := strconv.ParseUint(hex.EncodeToString(buff[13:17]), 16, 32)
		if err == nil {
			version := fmt.Sprintf("Version:%s", strconv.FormatUint(ver, 10))
			result["identify.bool"] = true
			result["identify.string"] = fmt.Sprintf("[%s]", version)
		}
		result["protocol"] = "activemq"
		return true
	}
	return false
}
