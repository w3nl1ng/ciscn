package scanner

import "log"

type Scanner struct {
	IpFileName string
	LiveIP     []string
	ScanResult map[string]LiveIPInfo //键为存活的ip
}

type LiveIPInfo struct {
	Services   []PortInfo  `json:"services"`
	DeviceInfo interface{} `json:"deviceinfo"`
	HoneyPot   []string    `json:"honeypot"`
	TimeStamp  string      `json:"timestamp"`
}

type PortInfo struct {
	Port       int         `json:"port"`
	Protocol   interface{} `json:"protocol"`
	ServiceApp interface{} `json:"service_app"`
}

func NewScanner(fileName string) *Scanner {

	if fileName == "" {
		log.Println("scanner/NewScanner: fileName is not specify")
		return nil
	}

	return &Scanner{IpFileName: fileName}
}
