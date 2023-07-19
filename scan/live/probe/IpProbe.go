package probe

import (
	"ciscn/scan/live/listener"
	"ciscn/scan/live/message"
	"log"
	"net"
	"os"
	"time"
)

func CheckIfIPLive(host string) bool {
	ipAddr, err := net.ResolveIPAddr("ip", host)

	msg := message.NewIcmpEchoMessage()
	msgBytes, err := msg.Marshal(nil)
	if err != nil {
		log.Fatalf("无法构建ICMP消息：%v", err)
	}

	_, err = listener.Conn.WriteTo(msgBytes, ipAddr)
	if err != nil {
		log.Fatalf("无法发送ICMP请求：%v", err)
	}

	err = listener.Conn.SetReadDeadline(time.Now().Add(2000 * time.Millisecond))
	if err != nil {
		log.Fatalf("无法设置读取超时时间：%v", err)
		os.Exit(0)
	}

	response := make([]byte, 100)
	_, _, err = listener.Conn.ReadFrom(response)
	if err != nil {
		return false
	}

	return true
}
