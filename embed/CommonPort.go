package embed

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed TopThousandPort.txt
var portData string
var CommonPort []int

func init() {
	CommonPort = parsePorts(portData)
}

func parsePorts(data string) []int {
	strPorts := strings.Split(data, ",")

	ports := make([]int, 0, len(strPorts))

	for _, strPort := range strPorts {
		port, err := strconv.Atoi(strings.TrimSpace(strPort))
		if err != nil {
			log.Printf("无法解析端口号：%v\n", err)
			continue
		}

		ports = append(ports, port)
	}

	return ports
}
