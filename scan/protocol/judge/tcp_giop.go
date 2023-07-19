package judge

import (
	Conn "ciscn/scan/protocol/conn"
	"ciscn/scan/protocol/parse"
	"encoding/hex"
	"strings"
)

func TcpGIOP(result map[string]interface{}, Args map[string]interface{}) bool {
	timeout := Args["FlagTimeout"].(int)
	host := result["host"].(string)
	port := result["port"].(int)

	conn, err := Conn.ConnTcp(host, port, timeout)
	if err != nil {
		return false
	}

	msg := "\x47\x49\x4f\x50\x01\x02\x00\x03\x00\x00\x00\x17\x00\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x0b\x4e\x61\x6d\x65\x53\x65\x72\x76\x69\x63\x65"
	_, err = conn.Write([]byte(msg))
	if err != nil {
		return false
	}

	reply := make([]byte, 256)
	_, _ = conn.Read(reply)
	if conn != nil {
		_ = conn.Close()
	}

	if strings.Contains(hex.EncodeToString(reply[0:4]), "47494f50") == false {
		return false
	}

	result["protocol"] = "giop"
	result["banner.string"] = parse.ByteToStringParse2(reply[0:4])
	result["banner.byte"] = reply
	return true
}
