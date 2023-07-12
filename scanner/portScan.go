package scanner

import (
	"ciscn/config"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/panjf2000/ants/v2"
)

var Mu1 sync.Mutex
var Mu2 sync.Mutex
var Mu3 sync.Mutex
var Mu4 sync.Mutex

var TempOpenPorts []string
var TempProtocols []string
var TempResult map[string]LiveIPInfo

func insertToPort(ports []string) {
	Mu1.Lock()
	TempOpenPorts = append(TempOpenPorts, ports...)
	Mu1.Unlock()
}

func insertToProtocol(protocol []string) {
	Mu2.Lock()
	TempProtocols = append(TempProtocols, protocol...)
	Mu2.Unlock()
}

func InsertToTempResult(ip string, lp LiveIPInfo) {
	Mu3.Lock()
	TempResult[ip] = lp
	Mu3.Unlock()
}

func clearPortsAndProtocol() {
	Mu4.Lock()
	TempOpenPorts = []string{}
	TempProtocols = []string{}
	Mu4.Unlock()
}

func scanPort(i interface{}) {
	arg, ok := i.(struct {
		port string
		ip   string
	})
	if !ok {
		log.Printf("scanner/scanPort: can not convert type(%T) to type(string)\n", i)
		return
	}
	port := arg.port
	ip := arg.ip

	log.Printf("scanner/scanPort: begin scanning %s", ip)
	args := []string{"-p", port, ip}
	output := Run(args)
	// fmt.Println(string(output))
	openPorts, protocols := findOpenPort(string(output))
	if openPorts == nil || protocols == nil {
		log.Printf("scanner/scanPort: not found open in %s", ip)
	} else {
		insertToPort(openPorts)
		insertToProtocol(protocols)

		// fmt.Println(openPorts)
		// fmt.Println(protocols)
		// fmt.Println("here")
		var liveIpInfo LiveIPInfo

		for i := 0; i < len(TempOpenPorts); i++ {
			var portInfo PortInfo
			portInfo.Port, _ = strconv.Atoi(TempOpenPorts[i])
			portInfo.Protocol = TempProtocols[i]
			liveIpInfo.Services = append(liveIpInfo.Services, portInfo)
		}

		InsertToTempResult(ip, liveIpInfo)

		// fmt.Println(TempOpenPorts)
		// fmt.Println(TempProtocols)
		clearPortsAndProtocol()
	}
	log.Printf("scanner/scanPort: finish scanning %s", ip)
}

// 此函数扫描LiveIP中的IP，找出存活的端口和对应的协议
func (sc *Scanner) portScan() {
	var Ports string
	var wg sync.WaitGroup
	//var args []string

	TempResult = make(map[string]LiveIPInfo)

	for _, v := range config.TopPorts {
		Ports += strconv.Itoa(v)
		Ports += ","
	}
	Ports = strings.TrimSuffix(Ports, ",")

	p, err := ants.NewPoolWithFunc(10, func(i interface{}) {
		scanPort(i)
		wg.Done()
	})

	if err != nil {
		log.Printf("scanner/portScan: %v\n", err)
	}
	defer p.Release()

	for _, liveIp := range sc.LiveIP {
		wg.Add(1)
		var arg struct {
			port string
			ip   string
		}
		arg.port = Ports
		arg.ip = liveIp
		_ = p.Invoke(arg)

		// for i := 0; i < len(TempOpenPorts); i++{
		// 	var portInfo PortInfo
		// 	portInfo.Port, _ = strconv.Atoi(TempOpenPorts[i])
		// 	portInfo.Protocol = TempProtocols[i]
		// 	var liveIpInfo LiveIPInfo
		// 	liveIpInfo.Services = append(liveIpInfo.Services, portInfo)
		// 	sc.ScanResult = make(map[string]LiveIPInfo)
		// 	sc.ScanResult[liveIp] = liveIpInfo
		// }

		// // fmt.Println(TempOpenPorts)
		// // fmt.Println(TempProtocols)

		// TempOpenPorts = []string{}
		// TempProtocols = []string{}

	}
	wg.Wait()
	// fmt.Println(TempResult)

	sc.ScanResult = TempResult

	//ip = "baidu.com"
	//cmd := exec.Command("nmap", "-p", Ports, ip)
	//out, err := cmd.CombinedOutput()
	//if err != nil {
	//	log.Fatal("cmd.Run() failed with %s\n", err)
	//	// fmt.Println(string(out))
	//}
	//// fmt.Println(string(out))

	////确定存活列表标志
	//Flag := "PORT    STATE SERVICE"
	//Flag1 := "open"
	//Flag2 := "filtered"
	//
	//var lines []string
	//var portLive []string
	//var protocol []string
	//var lineNum int
	//
	//lines = strings.Split(string(out), "\n")
	//
	//for i, v := range lines {
	//	if strings.Contains(v, Flag) {
	//		lineNum = i
	//	}
	//}
	//
	//for i := lineNum + 1; i < len(lines); i++ {
	//	v := lines[i]
	//	if (strings.Contains(v, Flag1)) || (strings.Contains(v, Flag2)) {
	//		liveStr1 := strings.Split(v, " ")
	//		liveStr2 := strings.Split(liveStr1[0], "/")
	//		portLive = append(portLive, liveStr2[0])
	//		protocol = append(protocol, liveStr2[1])
	//	}
	//}
	//// fmt.Println(lines)
	//fmt.Println(portLive)
	//fmt.Println(protocol)

}

func findOpenPort(out string) ([]string, []string) {
	//确定存活列表标志
	Flag := "STATE"
	Flag1 := "open"
	Flag2 := "filtered"

	var lines []string
	var portLive []string
	var protocol []string
	var lineNum int

	lines = strings.Split(string(out), "\n")

	for i, v := range lines {
		if strings.Contains(v, Flag) {
			lineNum = i
			break
		}
		// print(strings.Contains(lines[4],Flag))
		// fmt.Print("i, v:")
		// fmt.Println(i)
		// fmt.Println(v)
	}

	// fmt.Print("linenum:")
	// fmt.Println(lineNum)
	for i := lineNum + 1; i < len(lines) && lineNum != 0; i++ {
		v := lines[i]
		// fmt.Println(v)
		if (strings.Contains(v, Flag1)) || (strings.Contains(v, Flag2)) {
			liveStr1 := strings.Split(v, " ")
			liveStr2 := strings.Split(liveStr1[0], "/")
			portLive = append(portLive, liveStr2[0])
			protocol = append(protocol, liveStr2[1])
		}
	}
	// fmt.Println(lines)
	// fmt.Println(portLive)
	// fmt.Println(protocol)

	return portLive, protocol
}
