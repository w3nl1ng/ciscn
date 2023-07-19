package probe

import (
	"ciscn/common"
	"errors"
	"golang.org/x/net/proxy"
	"net"
	"net/url"
	"strings"
	"time"
)

//func SendMessageTCP(address string) bool {
//	timeout := 10 * time.Second
//	conn, err := net.DialTimeout("tcp", address, timeout)
//	if err != nil {
//		return false
//	}
//	defer conn.Close()
//
//	//_, err = conn.Write([]byte(strings.ReplaceAll(message, "\\r\\n", "\r\n")))
//	//if err != nil {
//	//	return "", err
//	//}
//	//
//	//response := make([]byte, 1024)
//	//n, err := conn.Read(response)
//	//
//	//if err != nil && err.Error() != "EOF" {
//	//	return "", err
//	//}
//	//return string(response[:n]), nil
//	return true
//}

func WrapperTcpWithTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	d := &net.Dialer{Timeout: timeout}
	return WrapperTCP(network, address, d)
}

func WrapperTCP(network, address string, forward *net.Dialer) (net.Conn, error) {
	//get conn
	var conn net.Conn
	if common.Config.Socks5Proxy == "" {
		var err error
		conn, err = forward.Dial(network, address)
		if err != nil {
			return nil, err
		}
	} else {
		dailer, err := Socks5Dailer(forward)
		if err != nil {
			return nil, err
		}
		conn, err = dailer.Dial(network, address)
		if err != nil {
			return nil, err
		}
	}
	return conn, nil

}

func Socks5Dailer(forward *net.Dialer) (proxy.Dialer, error) {
	u, err := url.Parse(common.Config.Socks5Proxy)
	if err != nil {
		return nil, err
	}
	if strings.ToLower(u.Scheme) != "socks5" {
		return nil, errors.New("Only support socks5")
	}
	address := u.Host
	var auth proxy.Auth
	var dailer proxy.Dialer
	if u.User.String() != "" {
		auth = proxy.Auth{}
		auth.User = u.User.Username()
		password, _ := u.User.Password()
		auth.Password = password
		dailer, err = proxy.SOCKS5("tcp", address, &auth, forward)
	} else {
		dailer, err = proxy.SOCKS5("tcp", address, nil, forward)
	}

	if err != nil {
		return nil, err
	}
	return dailer, nil
}

func TcpProbe(target string) bool {
	timeout := 2000 * time.Millisecond
	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

//func TcpProbe(probe embed.TooLTT, target string) bool {
//	probeString := common.FixRegexpPattenFormat(probe.Probestring)
//	resp, err := SendMessageTCP(target, probeString)
//	if err != nil {
//		log.Printf("在向目标 %s 发送探针 %s 失败 : %s", target, probeString, err)
//	}
//	for i := 0; i < len(probe.Matches); i++ {
//		re := regexp.MustCompile(common.FixRegexpPattenFormat(probe.Matches[i].Pattern))
//		match := re.FindStringSubmatch(resp)
//		if len(match) > 1 {
//			version := match[1]
//			fmt.Println("发现目标: ", target)
//			fmt.Println("服务名称:", probe.Matches[i].Versioninfo.Cpename)
//			fmt.Println("设备名称:", probe.Matches[i].Versioninfo.Devicetype)
//			fmt.Println("版本号:", version)
//			return true
//		}
//	}
//	//fmt.Println(resp)
//	return false
//}
