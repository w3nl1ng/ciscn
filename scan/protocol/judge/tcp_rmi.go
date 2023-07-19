package judge

import (
	"bytes"
	Conn "ciscn/scan/protocol/conn"
	"ciscn/scan/protocol/parse"
	"encoding/hex"
)

func TcpRMI(result map[string]interface{}, Args map[string]interface{}) bool {
	timeout := Args["FlagTimeout"].(int)
	host := result["host"].(string)
	port := result["port"].(int)

	conn, err := Conn.ConnTcp(host, port, timeout)
	if err != nil {
		return false
	}

	msg := "\x4a\x52\x4d\x49\x00\x02\x4b"
	_, err = conn.Write([]byte(msg))
	if err != nil {
		return false
	}

	reply := make([]byte, 256)
	_, _ = conn.Read(reply)
	if conn != nil {
		_ = conn.Close()
	}

	var buffer [256]byte
	if bytes.Equal(reply[:], buffer[:]) {
		return false
	} else if hex.EncodeToString(reply[0:1]) != "4e" {
		return false
	}
	result["protocol"] = "rmi"
	result["banner.string"] = parse.ByteToStringParse1(reply)
	result["banner.byte"] = reply
	return true
}
