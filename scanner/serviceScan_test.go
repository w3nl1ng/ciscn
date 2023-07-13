package scanner

import (
	"fmt"
	"testing"
)

func Test_parseServiceFromOutput(t *testing.T) {
	output := `Starting Nmap 7.80 ( https://nmap.org ) at 2023-07-12 17:44 CST
	Nmap scan report for 43.135.46.47
	Host is up (0.029s latency).
	
	PORT     STATE SERVICE  VERSION
	21/tcp   open  ftp      Pure-FTPd
	22/tcp   open  ssh      OpenSSH 7.4 (protocol 2.0)
	80/tcp   open  http     nginx
	443/tcp  open  ssl/http nginx
	888/tcp  open  http     nginx
	3306/tcp open  mysql    MySQL 5.6.50-log
	8888/tcp open  http     nginx
	
	Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
	Nmap done: 1 IP address (1 host up) scanned in 23.40 seconds`

	service := parseServiceFromOutput(output)
	fmt.Println(service)
}

func Test_ServiceScan(t *testing.T) {
	sc := NewScanner("iplist-test.txt")
	sc.ScanLiveIP()
	sc.PortScan()
	sc.ServiceScan()
	fmt.Println(sc.ScanResult)
}
