package probe

import (
	"fmt"
	"net"
	"testing"
	"time"
)

//
//func TestIsGlastopfHoneyPot(t *testing.T) {
//	//fmt.Println(IsGlastopfHoneyPot("http://49.232.214.202:8082/"))
//	fmt.Println(IsGlastopfHoneyPot("https://zhuanlan.zhihu.com/p/342849212"))
//}

func TestIsGlastopfHoneyPot(t *testing.T) {
	// 创建TCP连接
	conn, err := net.DialTimeout("tcp", "88.102.211.102:22", 10*time.Second)
	//conn, err := net.Dial("tcp", "88.102.211.102:22")
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer conn.Close()

	// 接收banner
	banner := make([]byte, 1024)
	_, err = conn.Read(banner)
	if err != nil {
		fmt.Println("无法接收数据:", err)
		return
	}
	//fmt.Println(string(banner[:n]))

	// 发送SSH协议版本信息
	sshVersion := "SSH-2.1-OpenSSH_5.9p1\r\n"
	_, err = conn.Write([]byte(sshVersion))
	if err != nil {
		fmt.Println("发送失败:", err)
		return
	}

	// 接收服务器响应
	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		fmt.Println("无法接收数据:", err)
		return
	}
	fmt.Println(string(response[:n]))
}
