package judge

import (
	"bytes"
	Conn "ciscn/scan/protocol/conn"
	"ciscn/scan/protocol/parse"
	"encoding/hex"
	"fmt"
	"net"
	"regexp"
	"strconv"
)

func TcpRTSP(result map[string]interface{}, Args map[string]interface{}) bool {
	var buff []byte
	buff, _ = result["banner.byte"].([]byte)
	ok, err := regexp.Match(`^RTSP/`, buff)
	if err != nil {
		return false
	}
	if ok {
		result["protocol"] = "rtsp"
		return true
	}

	if rtsp(result, Args) {
		return true
	}
	return false
}

func rtsp(result map[string]interface{}, Args map[string]interface{}) bool {
	timeout := Args["FlagTimeout"].(int)
	host := result["host"].(string)
	port := result["port"].(int)

	address := net.JoinHostPort(host, strconv.Itoa(port))
	conn, err := Conn.ConnTcp(host, port, timeout)
	if err != nil {
		return false
	}

	msg := fmt.Sprintf("OPTIONS rtsp://%s RTSP/1.0\r\nCSeq:1\r\n\r\n", address)
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
	} else if hex.EncodeToString(reply[0:4]) != "52545350" {
		return false
	}
	result["protocol"] = "rtsp"
	result["banner.string"] = parse.ByteToStringParse1(reply)
	result["banner.byte"] = reply
	return true
}
