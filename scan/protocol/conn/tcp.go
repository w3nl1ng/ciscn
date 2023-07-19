package conn

import (
	"net"
	"strconv"
	"time"
)

func ConnTcp(host string, port int, timeout int) (net.Conn, error) {
	target := net.JoinHostPort(host, strconv.Itoa(port))

	var conn net.Conn

	conn, err := net.DialTimeout("tcp", target, time.Duration(timeout)*time.Second)
	if err != nil {
		return nil, err
	}

	err = conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	if err != nil {
		if conn != nil {
			_ = conn.Close()
		}
		return nil, err
	}
	return conn, nil
}
