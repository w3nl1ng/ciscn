package probe

import (
	"fmt"
	"testing"
)

func TestIsKippoHoneyPot(t *testing.T) {
	//targetAddress := "127.0.0.1:22"
	//
	//go func(targetAddress string) {
	//	listener, err := net.Listen("tcp", targetAddress)
	//	defer listener.Close()
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		fmt.Println("无法接受连接:", err)
	//		return
	//	}
	//	defer conn.Close()
	//
	//	conn.Write([]byte("SSH-2.0-OpenSSH_8.8"))
	//	buffer := make([]byte, 1024)
	//
	//	n, err := conn.Read(buffer)
	//	if err != nil {
	//		fmt.Println("无法读取数据:", err)
	//		return
	//	}
	//
	//	// 将接收到的数据转换成字符串并打印
	//	receivedData := string(buffer[:n])
	//	fmt.Println("对方发送的消息:", receivedData)
	//	conn.Write([]byte("Protocol major versions differ"))
	//	//conn.Write([]byte("bad version 1.9"))
	//	fmt.Println("写入 Protocol major versions differ ")
	//	//fmt.Println("写入 bad version 1.9 ")
	//
	//}(targetAddress)

	//fmt.Println(IsKippoHoneyPot(targetAddress))
	//fmt.Println(IsKippoHoneyPot("134.122.46.170:22"))
	//fmt.Println(IsKippoHoneyPot("172.22.122.59:22"))
	//fmt.Println(IsKippoHoneyPot("185.139.228.48:2222"))
	fmt.Println(IsKippoHoneyPot("88.102.211.102:22"))
}
