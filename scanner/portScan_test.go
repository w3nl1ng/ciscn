package scanner

import (
	"fmt"
	"testing"
)

func Test_findOpenPort(t *testing.T) {
	output := `Starting Nmap 7.80 ( https://nmap.org ) at 2023-07-11 20:41 CST
Nmap scan report for baidu.com (39.156.66.10)
Host is up (0.029s latency).
Other addresses for baidu.com (not scanned): 110.242.68.66
	
PORT     STATE    SERVICE
22/tcp   filtered ssh
80/tcp   open     http
443/tcp  open     https
8080/tcp filtered http-proxy
	
Nmap done: 1 IP address (1 host up) scanned in 2.32 seconds`
	ports, porotcols := findOpenPort(output)
	fmt.Println(ports, porotcols)
}

func Test_portScan(t *testing.T) {
	sc := NewScanner("../iplist.txt")
	sc.ScanLiveIP()
	// sc.LiveIP = []string{"16.163.13.255", "16.163.13.251", "16.163.13.250", "16.163.13.249"}
	// sc.LiveIP = []string{"16.163.13.255"}
	sc.PortScan()
	fmt.Println(sc.ScanResult)
}

// func Test_scanPort(t *testing.T) {
// 	_ = p.Invoke(arg)
// 	// fmt.Println(sc.ScanResult)
// }
