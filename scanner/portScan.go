package scanner

import (
	"ciscn/config"
	"log"
	"runtime"
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
	args := []string{"-sS", "-p", port, ip}
	output := Run(args)
	// fmt.Println(string(output))
	openPorts, protocols := findOpenPort(string(output))
	if openPorts == nil || protocols == nil {
		log.Printf("scanner/scanPort: not found open in %s", ip)
	} else {
		insertToPort(openPorts)
		insertToProtocol(protocols)

		var liveIpInfo LiveIPInfo

		for i := 0; i < len(TempOpenPorts); i++ {
			var portInfo PortInfo
			portInfo.Port, _ = strconv.Atoi(TempOpenPorts[i])
			portInfo.Protocol = TempProtocols[i]
			liveIpInfo.Services = append(liveIpInfo.Services, portInfo)
		}

		InsertToTempResult(ip, liveIpInfo)

		clearPortsAndProtocol()
	}
	log.Printf("scanner/scanPort: finish scanning %s", ip)
}

// 此函数扫描LiveIP中的IP，找出存活的端口和对应的协议
func (sc *Scanner) PortScan() {
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
	}
	wg.Wait()

	sc.ScanResult = TempResult

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

	//windows与linux的换行符不一致，需要区别对待
	var sep string
	if runtime.GOOS == "windows" {
		sep = "\r\n"
	} else {
		sep = "\n"
	}

	lines = strings.Split(string(out), sep)

	for i, v := range lines {
		if strings.Contains(v, Flag) {
			lineNum = i
			break
		}
	}

	for i := lineNum + 1; i < len(lines) && lineNum != 0; i++ {
		v := lines[i]
		if (strings.Contains(v, Flag1)) || (strings.Contains(v, Flag2)) {
			liveStr1 := strings.Split(v, " ")
			liveStr2 := strings.Split(liveStr1[0], "/")
			portLive = append(portLive, liveStr2[0])
			if strings.Compare(liveStr1[len(liveStr1)-1], "http-proxy") == 0 {
				protocol = append(protocol, "http")
			} else {
				protocol = append(protocol, liveStr1[len(liveStr1)-1])
			}
		}
	}

	return portLive, protocol
}
