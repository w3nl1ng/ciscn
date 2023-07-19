package probe

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func IsKippoHoneyPot(target string) bool {

	conn, err := net.DialTimeout("tcp", target, 5*time.Second)
	if err != nil {
		fmt.Println("连接失败:", err)
		return false
	}
	defer conn.Close()

	banner := make([]byte, 1024)
	_, err = conn.Read(banner)
	if err != nil {
		fmt.Println("无法接收数据:", err)
		return false
	}

	sshVersion := "SSH-1.9-OpenSSH_5.9p1\r\n"
	_, err = conn.Write([]byte(sshVersion))
	if err != nil {
		fmt.Println("发送失败:", err)
		return false
	}

	response := make([]byte, 128)
	n, err := conn.Read(response)
	if err != nil {
		fmt.Println("无法接收数据:", err)
		return false
	}
	echoResponse := string(response[:n])
	if strings.Contains(echoResponse, "bad version 1.9") {
		return true
	}
	return false
}
