package scanner

import (
	// "fmt"
	"ciscn/config"
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
	// fmt.Println(deviceInfo)
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
	var deviceInfo config.DeviceInfo
	Flag := "Device type"
	lines := strings.Split(out, "\n")
	InFostr := "" 

	var lineNum int
	for i, v := range lines {
		if strings.Contains(v, Flag){
			lineNum = i
			break
		}
	}

	// fmt.Println(lines[1])
	// fmt.Println(lineNum)

	// fmt.Println(out)

	if lineNum != 0 {
		info := lines[lineNum]
		str := strings.Split(info, "|")
		str[0] = strings.Split(str[0], ": ")[1]
		//从前往后匹配
		for _, v := range str{
			switch v{
			case "general purpose":
				// deviceInfo.Name = "general purpose"
				// deviceInfo.Type = "general purpose"
				InFostr = deviceInfo.Type + deviceInfo.Name
				return InFostr
			case "firewall":
				deviceInfo.Name = "pfsense"
				deviceInfo.Type = "firewall"
				InFostr = deviceInfo.Type + deviceInfo.Name
				return InFostr
			case "Webcam":
				deviceInfo.Name = "Hikvision"
				deviceInfo.Type = "Webcam"
				InFostr = deviceInfo.Type + deviceInfo.Name
				return InFostr
			case "switch":
				deviceInfo.Name = "cisco"
				deviceInfo.Type = "switch"
				InFostr = deviceInfo.Type + deviceInfo.Name
				return InFostr
			case "storage-misc":
				deviceInfo.Name = "synology"
				deviceInfo.Type = "Nas"
				InFostr = deviceInfo.Type + deviceInfo.Name
				return InFostr
			}
			// if strings.Contains(v ,"general purpose"){
			// 	return deviceInfo
			// } else if (strings.Contains(v ,"firewall")) {
			// 	deviceInfo.Name = "pfsense"
			// 	deviceInfo.Type = "firewall"
			// 	return deviceInfo
			// } else if (strings.Contains(v ,"Webcam")) {
			// 	deviceInfo.Name = "Hikvision"
			// 	deviceInfo.Type = "Webcam"
			// 	return deviceInfo
			// } else if (strings.Contains(v ,"switch")) {
			// 	deviceInfo.Name = "cisco"
			// 	deviceInfo.Type = "switch"
			// 	return deviceInfo
			// } else if (strings.Contains(v ,"storage-misc")) {
			// 	deviceInfo.Name = "synology"
			// 	deviceInfo.Type = "Nas"
			// 	return deviceInfo
			// }
		}
		return InFostr
	} else {
		return InFostr
	}
}