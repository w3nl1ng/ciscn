package promotol

import (
	"bytes"
	"ciscn/common"
	"ciscn/scan/protocol/judge"
)

func JudgeTls(result map[string]interface{}, Args map[string]interface{}, resp *common.PortInfo) bool {
	protocol := result["protocol"].(string)
	runAll := true
	//if protocol != "" {
	//	runAll = false
	//}

	//在此处填充结构体中的协议类型和
	if protocol == "http" || protocol == "https" || runAll {
		if judge.TlsHTTPS(result, Args) {
			resp.Protocol = "https"
			return true
		}
	}
	if protocol == "rdp" || runAll {
		if judge.TlsRDP(result, Args) {
			resp.Protocol = "rdp"
			return true
		}
	}
	if protocol == "redis-ssl" || runAll {
		if judge.TlsRedisSsl(result) {
			resp.Protocol = "redis-ssl"
			resp.ServiceApp = []string{"redis/N"}
			return true
		}
	}

	status := result["status"].(string)
	if status == "open" && runAll {
		resp.Protocol = "unknown"
	}
	return false
}

func JudgeTcp(result map[string]interface{}, Args map[string]interface{}, resp *common.PortInfo) bool {
	protocol := result["protocol"].(string)
	runAll := true
	//if protocol != "" {
	//	runAll = false
	//}
	if protocol == "http" || protocol == "https" || runAll {
		result, ok := judge.TcpHTTP(result, Args)
		if ok {
			resp.Protocol = "http"
			resp.ServiceApp = result["identify.string"]
			return true
		}
		//if result.IsDetected {
		//	resp.Protocol = "http"
		//	resp.ServiceApp = result.ServiceApp
		//	resp.Deviceinfo = result.Deviceinfo
		//	return true
		//}

	}
	if protocol == "mysql" || runAll {
		if judge.TcpMysql(result) {
			resp.Protocol = "mysql"
			resp.ServiceApp = []string{"mysql/N"}
			return true
		}
	}
	if protocol == "redis" || runAll {
		if judge.TcpRedis(result) {
			resp.Protocol = "redis"
			resp.ServiceApp = []string{"redis/N"}
			return true
		}
	}
	if protocol == "smtp" || runAll {
		if judge.TcpSMTP(result) {
			resp.Protocol = "smtp"
			return true
		}
	}
	if protocol == "imap" || runAll {
		if judge.TcpIMAP(result) {
			resp.Protocol = "imap"
			return true
		}
	}
	if protocol == "ssh" || runAll {
		serviceApp, ok := judge.TcpSSH(result)
		if ok {
			resp.Protocol = "ssh"
			resp.ServiceApp = []string{"SSH/" + serviceApp}
			return true
		}
	}
	if protocol == "pop3" || runAll {
		if judge.TcpPOP3(result) {
			resp.Protocol = "pop3"
			return true
		}
	}
	if protocol == "vnc" || runAll {
		if judge.TcpVNC(result) {
			resp.Protocol = "vnc"
			return true
		}
	}
	if protocol == "telnet" || runAll {
		if judge.TcpTelnet(result) {
			resp.Protocol = "telnet"
			return true
		}
	}
	if protocol == "ftp" || runAll {
		if judge.TcpFTP(result) {
			resp.Protocol = "ftp"
			return true
		}
	}
	if protocol == "snmp" || runAll {
		if judge.TcpSNMP(result) {
			resp.Protocol = "snmp"
			return true
		}
	}
	if protocol == "oracle" || runAll {
		if judge.TcpOracle(result, Args) {
			resp.Protocol = "oracle"
			resp.ServiceApp = []string{"oracle/N"}
			return true
		}
	}
	if protocol == "frp" || runAll {
		if judge.TcpFrp(result, Args) {
			resp.Protocol = "frp"
			return true
		}
	}
	if protocol == "socks" || runAll {
		if judge.TcpSocks(result, Args) {
			resp.Protocol = "socks"
			return true
		}
	}
	if protocol == "ldap" || runAll {
		if judge.TcpLDAP(result, Args) {
			resp.Protocol = "ldap"
			return true
		}
	}
	if protocol == "rmi" || runAll {
		if judge.TcpRMI(result, Args) {
			resp.Protocol = "rmi"
			return true
		}
	}
	if protocol == "activemq" || runAll {
		if judge.TcpActiveMQ(result) {
			resp.Protocol = "activemq"
			return true
		}
	}
	if protocol == "rtsp" || runAll {
		if judge.TcpRTSP(result, Args) {
			resp.Protocol = "rtsp"
			return true
		}
	}
	if protocol == "rdp" || runAll {
		if judge.TcpRDP(result, Args) {
			resp.Protocol = "rdp"
			return true
		}
	}

	if protocol == "dcerpc" || runAll {
		if judge.TcpDceRpc(result, Args) {
			resp.Protocol = "dcerpc"
			return true
		}
	}
	if protocol == "mssql" || runAll {
		if judge.TcpMssql(result, Args) {
			resp.Protocol = "mssql"
			resp.ServiceApp = []string{"mssql/N"}
			return true
		}
	}
	if protocol == "smb" || runAll {
		if judge.TcpSMB(result, Args) {
			resp.Protocol = "smb"
			return true
		}
	}
	if protocol == "giop" || runAll {
		if judge.TcpGIOP(result, Args) {
			resp.Protocol = "giop"
			return true
		}
	}

	status := result["status"].(string)
	if status == "open" && runAll {
		resp.Protocol = "unknown"
	}
	return false
}

func JudgeUdp(result map[string]interface{}, Args map[string]interface{}, resp *common.PortInfo) bool {
	protocol := result["protocol"].(string)
	runAll := true
	if protocol != "" {
		runAll = false
	}
	if protocol == "nbns" || runAll {
		if judge.UdpNbns(result, Args) {
			resp.Protocol = "nbns"
			return true
		}
	}

	var buffer [256]byte
	status := result["status"].(string)
	if bytes.Equal(result["banner.byte"].([]byte), buffer[:]) {
		result["status"] = "close"
		return false
	} else if status == "open" && runAll {
		resp.Protocol = "unknown"
		//logger.Failed(fmt.Sprintf("[%s] %s [%s]",logger.Cyan("UDP/unknown"), parse.SchemeParse(result), logger.Blue(banner)))
	}
	return false
}
