package scanner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

//此函数扫描存活ip的设备信息

func TestScanner_DeviceScanPlus(t *testing.T) {
	sc := Scanner{}

	jsonData, err := ioutil.ReadFile("../out.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Define a map to unmarshal the JSON data
	var ipInfoMap map[string]LiveIPInfo

	// Unmarshal the JSON data into the map
	err = json.Unmarshal(jsonData, &ipInfoMap)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	sc.ScanResult = ipInfoMap
	// Print the results
	for ip, _ := range ipInfoMap {
		sc.LiveIP = append(sc.LiveIP, ip)
	}
	sc.HoneyPotScanPlus()
}
