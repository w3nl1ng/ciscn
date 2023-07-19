package conn

import (
	"crypto/tls"
	"log"
	"net"
	"strconv"
	"time"
)

// 发起Tls连接，返回连接对象
func ConnTls(host string, port int, timeout int) (net.Conn, error) {
	target := net.JoinHostPort(host, strconv.Itoa(port))
	conn, err := tls.DialWithDialer(
		&net.Dialer{Timeout: time.Duration(timeout) * time.Second},
		"tcp",
		target,
		&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Println(err)
		return conn, err
	}
	err = conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	if err != nil {
		log.Println(err)
		return conn, err
	}
	return conn, nil
}
