package scanner

import (
	// "fmt"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/panjf2000/ants/v2"
)

var Mu_device sync.Mutex
var TempDeviceInfo []string

func insertToDeviceInfo(deviceInfo string) {
	Mu_device.Lock()
	TempDeviceInfo = append(TempDeviceInfo, deviceInfo)
	Mu_device.Unlock()
}

func scanDevice(i interface{}) {
	ipSubnet, ok := i.(string)
	if !ok {
		log.Printf("scanner/deviceScan: can not convert type(%T) to type(string)\n", i)
		return
	}
	log.Printf("scanner/deviceScan: begin scanner %s\n", ipSubnet)
	args := []string{"-O", ipSubnet}
	output := Run(args)
	deviceInfo := findDeviceInfo(string(output))
	insertToDeviceInfo(deviceInfo)
	log.Printf("scanner/deviceScan: finish scanning %s\n", ipSubnet)
}

//此函数扫描存活ip的设备信息
func (sc *Scanner) deviceScan() {
	ipListAll := sc.LiveIP

	var wg sync.WaitGroup
	p, err := ants.NewPoolWithFunc(10, func(i interface{}) {
		scanDevice(i)
		wg.Done()
	})
	if err != nil {
		log.Printf("scanner/deviceScan: %v\n", err)
	}
	defer p.Release()

	for _, ipSubnet := range ipListAll {
		wg.Add(1)
		_ = p.Invoke(ipSubnet)
	}
	wg.Wait()

	for i, v := range sc.LiveIP{
		var liveIpInfo LiveIPInfo
		liveIpInfo.DeviceInfo = TempDeviceInfo[i]
		liveIpInfo.HoneyPot = sc.ScanResult[v].HoneyPot
		liveIpInfo.TimeStamp = sc.ScanResult[v].TimeStamp
		liveIpInfo.Services = sc.ScanResult[v].Services
		fmt.Println(sc.ScanResult[v].Services)
		sc.ScanResult[v] = liveIpInfo
		// sc.ScanResult[v].DeviceInfo = TempDeviceInfo[i]
	}

}

func findDeviceInfo(out string) string {
	Flag := "Device type"
	lines := strings.Split(out, "\n")

	var lineNum int
	for i, v := range lines {
		if strings.Contains(v, Flag){
			lineNum = i
			break
		}
	}

	// fmt.Println(lines[1])
	// fmt.Println(lineNum)

	if lineNum != 0 {
		return lines[lineNum]
	} else {
		return ""
	}
}