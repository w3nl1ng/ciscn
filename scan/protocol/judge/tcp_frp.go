package judge

import (
	"bytes"
	Conn "ciscn/scan/protocol/conn"
	"encoding/hex"
	"fmt"
	"strings"
)

func TcpFrp(result map[string]interface{}, Args map[string]interface{}) bool {
	timeout := Args["FlagTimeout"].(int)
	host := result["host"].(string)
	port := result["port"].(int)

	conn, err := Conn.ConnTcp(host, port, timeout)
	if err != nil {
		return false
	}

	msg := "\x00\x01\x00\x01\x00\x00\x00\x01\x00\x00\x00\x00"
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
	} else if hex.EncodeToString(reply[0:12]) != "000100020000000100000000" {
		return false
	}
	result["protocol"] = "frp"
	result["banner.string"] = frpByteToStringParse(reply[0:12])
	result["banner.byte"] = reply
	return true
}

func frpByteToStringParse(p []byte) string {
	var w []string
	var res string
	for i := 0; i < len(p); i++ {
		asciiTo16 := fmt.Sprintf("\\x%s", hex.EncodeToString(p[i:i+1]))
		w = append(w, asciiTo16)
	}
	res = strings.Join(w, "")
	return res
}
