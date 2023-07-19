package probe

import (
	"ciscn/common"
	"ciscn/embed"
	"strconv"
	"strings"
)

//
//func PortConnect(addr server.Addr, respondingHosts chan<- string, wg *sync.WaitGroup) {
//	host, port := addr.Ip, addr.Port
//	conn, err := WrapperTcpWithTimeout("tcp4", fmt.Sprintf("%s:%v", host, port), time.Duration(1)*time.Second)
//	defer func() {
//		if conn != nil {
//			conn.Close()
//		}
//	}()
//	if err == nil {
//		address := host + ":" + strconv.Itoa(port)
//		wg.Add(1)
//		respondingHosts <- address
//	}
//}

//func CheckIfPortLiveAndJudgeServer(target string, port int) bool {
//	var checkServers []embed.TooLTT
//
//	for i := 0; i < len(embed.NmapServers); i++ {
//		checkServer, flag := ParsePortToProbeFromServers(embed.NmapServers[i], port)
//		if flag != false {
//			checkServers = append(checkServers, checkServer)
//		}
//	}
//
//	for j := 0; j < len(checkServers); j++ {
//		switch checkServers[j].Protocol {
//		case "TCP":
//			TcpProbe(checkServers[j], fmt.Sprintf("%s:%d", target, port))
//		}
//	}
//
//	return false
//}

func ParsePortToProbeFromServers(nmapServer embed.TooLTT, port int) (embed.TooLTT, bool) {
	var serverPorts []string
	var checkServer embed.TooLTT
	var flag = false

	for j := 0; j < len(nmapServer.Ports); j++ {
		if strings.Contains(nmapServer.Ports[j].(string), "-") {
			rangePort := common.ParsePortRangeToSlice(nmapServer.Ports[j].(string))
			serverPorts = append(serverPorts, rangePort...)
		}
	}

	if common.StringInSlice(strconv.Itoa(port), serverPorts) {
		flag = true
		checkServer = nmapServer
	}

	return checkServer, flag
}
