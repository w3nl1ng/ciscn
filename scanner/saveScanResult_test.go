package scanner

import "testing"

func TestScanner_SaveScanResult(t *testing.T) {
	sc := NewScanner("test.txt")
	sc.ScanResult = make(map[string]LiveIPInfo)
	sc.ScanResult["159.65.92.42"] = LiveIPInfo{
		Services: []PortInfo{
			{Port: 445, Protocol: "https", ServiceApp: []string{
				"apache/2.4.6", "centos/N", "openssl/1.0.2", "php/N", "wordpress/N",
			}},
			{Port: 80, Protocol: "http", ServiceApp: []string{
				"apache/2.4.6", "centos/N", "openssl/1.0.2", "php/N", "wordpress/N",
			}},
		},
	}
	sc.ScanResult["113.30.191.229"] = LiveIPInfo{
		Services: []PortInfo{
			{Port: 443, Protocol: "https", ServiceApp: []string{
				"nginx/N", "php/N",
			}},
			{Port: 2222, Protocol: "ssh", ServiceApp: []string{
				"openssh/7.9",
			}},
		},
		DeviceInfo: []string{"firewall/pfsense"},
	}
	sc.ScanResult["165.22.22.193"] = LiveIPInfo{
		Services: []PortInfo{
			{Port: 22, Protocol: "ssh", ServiceApp: []string{
				"openssh/8.7",
			}},
			{Port: 80, Protocol: "http", ServiceApp: []string{
				"nginx/1.18.0",
				"ubuntu/N",
			}},
			{Port: 554, Protocol: "rtsp"},
		},
		DeviceInfo: []string{"webcam/dahua"},
	}
	sc.ScanResult["103.252.119.251"] = LiveIPInfo{
		Services: []PortInfo{
			{Port: 6379, Protocol: "redis"},
			{Port: 3306, Protocol: "mysql"},
			{Port: 21, Protocol: "ftp"},
			{Port: 80, Protocol: "http", ServiceApp: []string{
				"nginx/1.18.0", "ubuntu/N",
			}},
			{Port: 1022, Protocol: "ssh", ServiceApp: []string{"openssh/7.4"}},
			{Port: 8083, Protocol: "http", ServiceApp: []string{
				"apache/2.4.56",
				"debian/N",
				"php/8.0.29",
			}},
			{Port: 7001, Protocol: "weblogic"},
			{Port: 9200, Protocol: "http", ServiceApp: []string{"elasticsearch/6.5.1"}},
			{Port: 2202, Protocol: "ssh", ServiceApp: []string{"openssh/6.6.1"}},
		},
	}
	sc.ScanResult["185.139.228.48"] = LiveIPInfo{
		Services: []PortInfo{
			{Port: 2222, Protocol: "ssh", ServiceApp: []string{"openssh/5.1"}},
			{Port: 3306, Protocol: "mysql"},
			{Port: 8080, Protocol: "http", ServiceApp: []string{"apache/1.1"}},
		},
		HoneyPot: []string{"2222/kippo"},
	}

	sc.SaveScanResult("result.json")
}
