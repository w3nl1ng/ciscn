package judge

import (
	"regexp"
)

func TcpSSH(result map[string]interface{}) (string, bool) {
	var buff []byte
	buff, _ = result["banner.byte"].([]byte)
	re := regexp.MustCompile(`^SSH-([\d.]+)-OpenSSH_([\w._-]+)`)

	match := re.FindStringSubmatch(string(buff))

	if len(match) > 1 {
		version := match[2]
		return version, true
	}

	return "", false
}
