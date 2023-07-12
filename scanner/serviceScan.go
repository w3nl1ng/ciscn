package scanner

import (
	"fmt"

	"log"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/panjf2000/ants/v2"
)

var Mu_service sync.Mutex
var TempServiceInfo []string

func insertToserviceInfo(serviceInfo string) {
	Mu_service.Lock()
	TempServiceInfo = append(TempServiceInfo, serviceInfo)
	Mu_service.Unlock()
}

// 此函数扫描存活的端口，确定服务和版本
func scanService(i interface{}) {
	arg, ok := i.(struct {
		port string
		ip   string
	})
	if !ok {
		log.Printf("scanner/scanService: can not convert type(%T) to type(string)\n", i)
		return
	}
	port := arg.port
	ip := arg.ip
	log.Printf("scanner/serviceScan: begin scanner %s\n", ip)
	args := []string{"-sV", "-p", port, ip}
	output := Run(args)
	deviceInfo := findServiceInfo(string(output))
	insertToserviceInfo(deviceInfo)
	log.Printf("scanner/serviceScan: finish scanning %s\n", ip)
}

func (sc *Scanner) serviceScan() {
	ipListAll := sc.LiveIP
	var liveIpInfo LiveIPInfo
	var wg sync.WaitGroup
	p, err := ants.NewPoolWithFunc(10, func(i interface{}) {
		scanService(i)
		wg.Done()
	})
	if err != nil {
		log.Printf("scanner/serviceScan: %v\n", err)
	}
	defer p.Release()

	for _, ipSubnet := range ipListAll {
		liveIpInfo = sc.ScanResult[ipSubnet]
		for _, portList := range liveIpInfo.Services {
			wg.Add(1)
			var arg struct {
				port string
				ip   string
			}
			arg.port = strconv.Itoa(portList.port)
			arg.ip = ipSubnet
			_ = p.Invoke(arg)
		}
	}
	wg.Wait()

	var count int
	for i, v := range sc.LiveIP {
		var liveIpInfo LiveIPInfo
		liveIpInfo.DeviceInfo = sc.ScanResult[v].DeviceInfo
		liveIpInfo.HoneyPot = sc.ScanResult[v].HoneyPot
		liveIpInfo.TimeStamp = sc.ScanResult[v].TimeStamp
		for ii, _ := range liveIpInfo.Services {
			liveIpInfo.Services[ii].ServiceApp = TempServiceInfo[count]
			liveIpInfo.Service[ii].Protocol = sc.ScanResult[v].Services[ii].Protocol
			liveIpInfo.Service[ii].Port = sc.ScanResult[v].Services[ii].Port
			count++
		}
		fmt.Println(sc.ScanResult[v].Services)
		sc.ScanResult[v] = liveIpInfo
		// sc.ScanResult[v].DeviceInfo = TempDeviceInfo[i]
	}
}

func findServiceInfo(out string) string {
	Flag := "VERSION"

	var lines []string
	var lineIndex int
	var index int
	var sep string
	var runeLine []rune
	var serviceStr []rune

	if runtime.GOOS == "windows" {
		sep = "\r\n"
	} else {
		sep = "\n"
	}

	lines = strings.Split(string(out), sep)

	for i, v := range lines {
		if strings.Contains(v, Flag) {
			index = strings.Index(v, Flag)
			lineIndex = i
			break
		}
	}

	if lineIndex != 0 && index != -1 {
		runeLine = []rune(lines[lineIndex+1])
		serviceStr = append(serviceStr, runeLine[index:]...)
		return string(serviceStr)
	} else {
		return ""
	}
}
