package common

var Config *ConfigInfo

type ConfigInfo struct {
	IPFile      string
	Socks5Proxy string
}

func init() {
	Config = &ConfigInfo{
		IPFile:      "",
		Socks5Proxy: "",
	}
}
