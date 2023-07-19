package server

import (
	"fmt"
	"github.com/w3nl1ng/ciscn-build-2023/common"
	"github.com/w3nl1ng/ciscn-build-2023/scan/live/probe"
	"log"
	"testing"
	"time"
)

func TestPortScanner(t *testing.T) {
	common.Config.Socks5Proxy = ""
	target := []string{"211.22.90.152"}
	liveAddr := StartLivedPortScan(target, 16)
	log.Println(liveAddr)
}
func TestConfigScanner(t *testing.T) {
	fmt.Println(common.Config.Socks5Proxy)
}

func TestDetectPortLiveTimes(t *testing.T) {
	startTime := time.Now()
	probe.TcpProbe("49.232.214.202:80")
	rtt := time.Since(startTime)
	fmt.Printf("Round-Trip Time: %v\n", rtt)
}
