package main

import (
	"ciscn/cmd"
	"ciscn/common"
	"ciscn/scan/live/server"
	promotol "ciscn/scan/protocol"
	"ciscn/scanner"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	cmd.Flag()

	if cmd.Model == "Custom_Scanner" {

		//
		//todo: Because of the tight time, this part was not finished
		//

		ipList := common.ParseIP(cmd.IpFileName)

		thread := 16
		liveAddr := server.StartLiveScannerServer(ipList, thread)
		log.Println(liveAddr)

		args := map[string]interface{}{
			"FlagTimeout": 2,
			"FlagType":    "tcp",
			"FlagMode":    "",
			"FlagUrl":     "",
		}

		for _, target := range liveAddr {
			host, portString := common.SplitAddressPort(target)
			port, err := strconv.Atoi(portString)
			if err != nil {
				log.Fatal("port to int fail")
			}
			resp := promotol.DiscoverTcp(host, port, args)
			fmt.Println(resp)
		}

	} else {
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
		sc.HoneyPotScanPlus()

		fmt.Println(sc.ScanResult)

		cost := time.Since(start)
		fmt.Printf("cost: %v", cost)

		sc.SaveScanResult(time.Now().Format(time.DateOnly) + ".txt")
	}
}
