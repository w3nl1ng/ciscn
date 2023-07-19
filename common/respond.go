package common

import (
	"encoding/json"
	"log"
)

type HttpServerDetect struct {
	IsDetected bool     `json:"is_detected"`
	ServiceApp []string `json:"service_app"`
	Deviceinfo []string `json:"device"`
}

type LiveIPInfo struct {
	Services   []PortInfo `json:"services"`
	DeviceInfo string     `json:"device"`
	HoneyPot   []string   `json:"honeypot"`
	TimeStamp  string     `json:"timestamp"`
}

type PortInfo struct {
	Host       string      `json:"host"`
	Port       int         `json:"port"`
	Protocol   interface{} `json:"protocol"`
	ServiceApp interface{} `json:"service_app"`
	Deviceinfo interface{} `json:"device"`
}

var ScanResult map[string]*LiveIPInfo

func ResultToString(scanResult map[string]LiveIPInfo) string {
	jByte, err := json.Marshal(scanResult)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(jByte)
}
