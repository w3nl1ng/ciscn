package scanner

import (
	"fmt"
	"testing"
)

func Test_parseIpFromFile(t *testing.T) {
	ipList := parseIpFromFile("..\\iplist.txt")
	fmt.Println(ipList)
}
