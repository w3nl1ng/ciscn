package judge

import (
	"bytes"
	Conn "ciscn/scan/protocol/conn"
	"ciscn/scan/protocol/parse"
	"encoding/hex"
	"strings"
)

func TcpLDAP(result map[string]interface{}, Args map[string]interface{}) bool {
	timeout := Args["FlagTimeout"].(int)
	host := result["host"].(string)
	port := result["port"].(int)

	conn, err := Conn.ConnTcp(host, port, timeout)
	if err != nil {
		return false
	}

	msg := "\x30\x0c\x02\x01\x01\x60\x07\x02\x01\x03\x04\x00\x80\x00"
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
	}
	if strings.Contains(hex.EncodeToString(reply), "010004000400") == false {
		return false
	}
	result["protocol"] = "ldap"
	result["banner.string"] = parse.ByteToStringParse2(reply[0:16])
	result["banner.byte"] = reply
	return true
}
