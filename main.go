package main

import (
	"ciscn/cmd"
	"ciscn/scanner"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	cmd.Flag()
	sc := scanner.NewScanner(cmd.IpFileName)
	if sc == nil {
		log.Println("main/main: failed to new a scanner")
		os.Exit(0)
	}

	start := time.Now()

	sc.ScanLiveIP()
	sc.PortScan()
	sc.ServiceScan()
	sc.DeviceScan()

	fmt.Println(sc.ScanResult)

	cost := time.Since(start)
	fmt.Printf("cost: %v", cost)

	sc.SaveScanResult("./out.json")
}
