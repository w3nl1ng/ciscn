package scanner

import (
	"ciscn/config"
	"fmt"
	"log"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/panjf2000/ants/v2"
)

type serviceInfo struct {
	Port     int
	Protocol string
	Service  string
	Version  string
}

type workArg struct {
	IP    string
	Ports string //类似 22.80.8888
}

var TempLiveIPInfo map[string]LiveIPInfo
var Mu_TempLiveIPInfo sync.Mutex

func init() {
	if TempLiveIPInfo == nil {
		TempLiveIPInfo = make(map[string]LiveIPInfo)
	}
}

func insetToTempLiveIPInfo(ip string, serInfo []serviceInfo) {
	Mu_TempLiveIPInfo.Lock()
	defer Mu_TempLiveIPInfo.Unlock()

	liveipinfo := LiveIPInfo{Services: make([]PortInfo, 0)}
	for _, si := range serInfo {
		temp := PortInfo{}
		temp.Port = si.Port
		temp.Protocol = si.Protocol
		if si.Service != "" {
			if si.Version != "" {
				temp.ServiceApp = []string{fmt.Sprintf("%s/%s", si.Service, si.Version)}
			} else { //版本没有探测出来
				temp.ServiceApp = []string{fmt.Sprintf("%s/N", si.Service)}
			}
		}
		liveipinfo.Services = append(liveipinfo.Services, temp)
	}
	TempLiveIPInfo[ip] = liveipinfo
}

func serviceScan(i interface{}) {
	args, ok := i.(workArg)
	if !ok {
		log.Printf("scanner/serviceScan: failed to convert type(%T) to type(workArg)\n", i)
		return
	}

	ip, ports := args.IP, args.Ports
	log.Printf("scanner/serviceScan: begin scanning %s\n", ip)
	runArgs := []string{"-sV", "-p", ports, ip}
	output := Run(runArgs)
	serInfo := parseServiceFromOutput(string(output))
	insetToTempLiveIPInfo(ip, serInfo)
	log.Printf("scanner/serviceScan: finish scanning %s\n", ip)
}

// 此函数扫描存活的端口，确定服务和版本
func (sc *Scanner) ServiceScan() {

	var wg sync.WaitGroup
	p, err := ants.NewPoolWithFunc(10, func(i interface{}) {
		serviceScan(i)
		wg.Done()
	})
	if err != nil {
		log.Printf("scanner/scanLiveIP: %v\n", err)
	}
	defer p.Release()

	for ip, liveIPInfo := range sc.ScanResult {
		wg.Add(1)
		var ports string
		//拼接端口
		for _, service := range liveIPInfo.Services {
			ports += fmt.Sprintf("%d,", service.Port)
		}
		args := workArg{IP: ip, Ports: ports[:len(ports)-1]} //除去ports的最后一个逗号
		err := p.Invoke(args)
		if err != nil {
			log.Printf("scanner/ServiceScan: %v\n", err)
		}
	}
	wg.Wait()

	for ip, liveipinfo := range TempLiveIPInfo {
		sc.ScanResult[ip] = liveipinfo
	}
}

func parseServiceFromOutput(output string) []serviceInfo {
	var sep string
	if runtime.GOOS == "windows" {
		sep = "\r\n"
	} else {
		sep = "\n"
	}

	lines := strings.Split(output, sep)

	var startIndex int
	for startIndex = 0; startIndex < len(lines); startIndex++ {
		if strings.Contains(lines[startIndex], "PORT") && strings.Contains(lines[startIndex], "VERSION") {
			startIndex += 1
			break
		}
	}

	var result []serviceInfo
	for ; startIndex < len(lines); startIndex++ {
		if !strings.Contains(lines[startIndex], "tcp") && !strings.Contains(lines[startIndex], "open") {
			break
		}
		lineArr := strings.Fields(lines[startIndex])
		if len(lineArr) < 3 {
			continue
		}
		port, err := strconv.Atoi(strings.Split(lineArr[0], "/")[0])
		if err != nil {
			log.Printf("scanner/parseServiceFromOutput: %v\n", err)
			continue
		}
		protocol := parseProtocol(lineArr[2])
		service, version := paserServiceAndVersion(strings.Join(lineArr[3:], " "))

		temp := serviceInfo{Port: port, Protocol: protocol, Service: service, Version: version}
		result = append(result, temp)
	}

	return result
}

func parseProtocol(protocol string) string {
	protocol = strings.ReplaceAll(protocol, "?", "")
	protocol = strings.ToLower(protocol)
	switch protocol {
	case "ssl/http":
		return "httpp"
	case "ssl/https":
		return "https"
	case "ssh":
		return "ssh"
	case "http":
		return "http"
	case "https":
		return "https"
	case "rtsp":
		return "rtsp"
	case "ftp":
		return "ftp"
	case "telnet":
		return "telnet"
	case "amqp":
		return "amqp"
	case "mongodb":
		return "mongodb"
	case "redis":
		return "redis"
	case "mysql":
		return "mysql"
	case "vnc-http":
		return "http"
	case "zeus-admin":
		return "http"
	case "opsmessaging":
		return "amqp"
	case "http-alt":
		return "http"
	case "ipp":
		return "http"

	default:
		return "unknown"
	}
}

func paserServiceAndVersion(s string) (string, string) {
	var service, version string
	//匹配服务
	for _, target := range config.ServiceList {
		if containWithoutCase(s, target) {
			service = target
		}
	}
	//匹配版本号
	versionlist := trimVersions(s)
	if versionlist != nil {
		version = versionlist[0] //默认取第一个类似版本号的字符串
	}

	return service, version
}

// containWithoutCase 函数判断str是否包含substr，忽略大小写
func containWithoutCase(str string, substr string) bool {
	strUp := strings.ToUpper(str)
	substrUp := strings.ToUpper(substr)
	return strings.Contains(strUp, substrUp)
}

// 匹配一个字符串中的所有版本号，如果未匹配成功则返回nil
func trimVersions(str string) []string {
	re := regexp.MustCompile(`\d+\.\d+(\.\d+)*`) // 匹配版本号的正则表达式
	return re.FindAllString(str, -1)
}
