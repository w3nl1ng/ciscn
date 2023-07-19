package get

import (
	"bytes"
	Conn "ciscn/scan/protocol/conn"
	"fmt"
	"time"
)

// 发起tcp连接，返回tcp连接对象

func TcpProtocol(host string, port int, timeout int) ([]byte, error) {
	conn, err := Conn.ConnTcp(host, port, timeout)
	if err != nil {
		return nil, err
	}
	_ = conn.SetDeadline(time.Now().Add(time.Duration(2) * time.Second))
	reply := make([]byte, 256)
	_, err = conn.Read(reply)

	var buffer [256]byte
	if err == nil && bytes.Equal(reply[:], buffer[:]) == false {
		if conn != nil {
			_ = conn.Close()
		}
		return reply, nil

	}
	conn, err = Conn.ConnTcp(host, port, timeout)
	if err != nil {
		return nil, err
	}
	msg := fmt.Sprintf("GET / HTTP/1.0\r\nHost: %s:%d\r\nUser-Agent: Mozilla/5.0 (Windows NT 6.0; Win64; x64; rv:106.0) Gecko/20100101 Firefox/106.0\r\nAccept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2\r\nAccept: */*\r\n\r\n", host, port)
	_, err = conn.Write([]byte(msg))
	if err != nil {
		return nil, err
	}
	_ = conn.SetDeadline(time.Now().Add(time.Duration(2) * time.Second))
	reply = make([]byte, 512)
	_, _ = conn.Read(reply)
	if conn != nil {
		_ = conn.Close()
	}
	return reply, nil
}
