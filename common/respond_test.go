package common

import (
	"fmt"
	"testing"
)

func TestSaveRespondToFile(t *testing.T) {
	test := make(map[string]LiveIPInfo)
	test["123.456.789"] = LiveIPInfo{Services: []PortInfo{{Port: 123, Protocol: "http", ServiceApp: []string{"telnet/N", "xxx/xxx"}}}, DeviceInfo: "webcam", HoneyPot: []string{"xxx/xxx", "xx/xxx"}, TimeStamp: "xxxxxx"}
	test["123.567.543"] = LiveIPInfo{Services: []PortInfo{{Port: 123, Protocol: nil, ServiceApp: nil}}, DeviceInfo: "webcam", HoneyPot: []string{"xxx/xxx", "xx/xxx"}, TimeStamp: "xxxxxx"}
	test["123.567.553"] = LiveIPInfo{Services: []PortInfo{{Port: 123, Protocol: nil, ServiceApp: nil}}, DeviceInfo: "webcam", TimeStamp: "xxxxxx"}
	fmt.Println(ResultToString(test))

}
