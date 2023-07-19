package promotol

import (
	"ciscn/common"
	"ciscn/scan/protocol/get"
	"ciscn/scan/protocol/parse"
	"time"
)

func setResult(host string, port int, Args map[string]interface{}) map[string]interface{} {
	var banner []byte
	result := map[string]interface{}{
		"date":            time.Now().Unix(),
		"status":          "None",
		"banner.byte":     banner,
		"banner.string":   "None",
		"protocol":        Args["FlagMode"].(string),
		"type":            Args["FlagType"].(string),
		"host":            host,
		"port":            port,
		"uri":             "None",
		"note":            "None",
		"path":            "",
		"identify.bool":   false,
		"identify.string": "None",
	}
	return result
}

func DiscoverTls(host string, port int, Args map[string]interface{}) common.PortInfo {

	resp := &common.PortInfo{
		Port:       port,
		Protocol:   nil,
		ServiceApp: nil,
	}

	result := setResult(host, port, Args)
	b, err := get.TlsProtocol(host, port, Args["FlagTimeout"].(int))
	if err != nil {
		return *resp
	}
	result["type"] = "tls"
	result["status"] = "open"
	result["banner.byte"] = b
	result["banner.string"] = parse.ByteToStringParse1(b)
	if JudgeTls(result, Args, resp) {
		return *resp
	}
	return *resp
}

func DiscoverTcp(host string, port int, Args map[string]interface{}) common.PortInfo {
	resp := &common.PortInfo{
		Host:       host,
		Port:       port,
		Protocol:   nil,
		ServiceApp: nil,
	}
	result := setResult(host, port, Args)
	b, err := get.TcpProtocol(host, port, Args["FlagTimeout"].(int))
	if err != nil {
		return *resp
	}
	result["type"] = "tcp"
	result["status"] = "open"
	result["banner.byte"] = b
	result["banner.string"] = parse.ByteToStringParse1(b)
	if JudgeTcp(result, Args, resp) {
		return *resp
	}
	return *resp
}

func DiscoverUdp(host string, port int, Args map[string]interface{}) common.PortInfo {
	resp := &common.PortInfo{
		Port:       port,
		Protocol:   nil,
		ServiceApp: nil,
	}
	result := setResult(host, port, Args)
	var udpPort = []int{53, 111, 123, 137, 138, 139, 12345}
	if isContainInt(udpPort, port) {
		return *resp
	}
	b, err := get.UdpProtocol(host, port, Args["FlagTimeout"].(int))
	if err != nil {
		return *resp
	}
	result["type"] = "tcp"
	result["status"] = "open"
	result["banner.byte"] = b
	result["banner.string"] = parse.ByteToStringParse1(b)
	if JudgeUdp(result, Args, resp) {
		return *resp
	}
	return *resp
}

func isContainInt(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
