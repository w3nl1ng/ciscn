package server

import (
	"log"
	"sync"
	"time"
)

var (
	WG sync.WaitGroup
)

func StartLiveScannerServer(ipList []string, thread int) []string {
	startTime := time.Now()
	liveIP := StartLivedIPScan(ipList, thread)
	rtt := time.Since(startTime)
	log.Printf("探活IP耗费时间 : %s", rtt.String())
	log.Println(liveIP)

	startTime = time.Now()
	liveAddr := StartLivedPortScan(liveIP, thread)
	rtt = time.Since(startTime)
	log.Printf("探活端口耗费时间 : %s", rtt.String())
	log.Println(liveAddr)

	return liveAddr
}
