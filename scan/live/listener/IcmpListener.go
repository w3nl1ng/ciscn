package listener

import (
	"context"
	"golang.org/x/net/icmp"
)

var Conn *icmp.PacketConn

type IcmpListener struct {
	Callback func(listener *IcmpListener)
	RootCtx  context.Context
	Cancel   context.CancelFunc
}

func NewIcmpListener() *IcmpListener {
	icmpListener := &IcmpListener{}
	icmpListener.Callback = IcmpListenerStartFunc

	icmpListener.Callback(icmpListener)

	return icmpListener
}

func IcmpListenerStartFunc(listener *IcmpListener) {

	rootCtx, cancel := context.WithCancel(context.Background())
	listener.RootCtx = rootCtx
	listener.Cancel = cancel

	Conn, _ = icmp.ListenPacket("ip4:icmp", "0.0.0.0")
}
