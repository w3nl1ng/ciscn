package message

import (
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"os"
)

func NewIcmpEchoMessage() *icmp.Message {
	return &icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: []byte(""),
		},
	}
}
