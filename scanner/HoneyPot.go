package scanner

import (
	"ciscn/scan/HoneyPot/probe"
	"fmt"
)

//此函数扫描存活ip的设备信息

func (sc *Scanner) HoneyPotScanPlus() {
	for ip, info := range sc.ScanResult {
		for _, service := range info.Services {
			if service.Protocol == "ssh" {
				address := fmt.Sprintf("%s:%d", ip, service.Port)
				fmt.Printf("开始检测%s是否存在Kippo蜜罐\r\n", address)
				if probe.IsKippoHoneyPot(address) {
					info.HoneyPot = append(info.HoneyPot.([]string), fmt.Sprintf("%d/%s", service.Port, "Kippo"))
					fmt.Println(address + " !! 发现Kippo蜜罐")
				}
				break
			}

			if service.Protocol == "http" {
				address := fmt.Sprintf("http://%s:%d", ip, service.Port)
				fmt.Printf("开始检测%s是否存在Glastopf 或者 HFish 蜜罐\n", address)
				if probe.IsGlastopfHoneyPot(address) {
					info.HoneyPot = append(info.HoneyPot.([]string), fmt.Sprintf("%d/%s", service.Port, "glastopf"))
					fmt.Println(address + " !! 发现Glastopf蜜罐")
				}

				if probe.IsHFishfHoneyPot(address) {
					info.HoneyPot = append(info.HoneyPot.([]string), fmt.Sprintf("%d/%s", service.Port, "HFish"))
					fmt.Println(address + " !! 发现HFish蜜罐")
				}
				break
			}

			if service.Protocol == "https" {
				address := fmt.Sprintf("https://%s:%d", ip, service.Port)
				fmt.Printf("开始检测%s是否存在Glastopf 或者 HFish 蜜罐\n", address)
				if probe.IsGlastopfHoneyPot(address) {
					info.HoneyPot = append(info.HoneyPot.([]string), fmt.Sprintf("%d/%s", service.Port, "glastopf"))
					fmt.Println(address + " !! glastopf")
				}

				if probe.IsHFishfHoneyPot(address) {
					info.HoneyPot = append(info.HoneyPot.([]string), fmt.Sprintf("%d/%s", service.Port, "HFish"))
					fmt.Println(address + " !! 发现HFish蜜罐")
				}
				break
			}
			//fmt.Printf("  Port: %d, Protocol: %v, ServiceApp: %v\n", service.Port, service.Protocol, service.ServiceApp)
		}
		//fmt.Println("DeviceInfo:", info.DeviceInfo)
		//fmt.Println("HoneyPot:", info.HoneyPot)
		//fmt.Println("TimeStamp:", info.TimeStamp)
		//fmt.Println()
	}
}
