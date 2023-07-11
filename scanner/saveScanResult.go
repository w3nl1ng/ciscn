package scanner

import (
	"encoding/json"
	"log"
	"os"
)

// SaveScanResult 函数将扫描的最终结果保存到文件
func (sc *Scanner) SaveScanResult(fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Printf("scanner/SaveScanResult: %v\n", err)
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("scanner/SaveScanResult: %v\n", err)
		}
	}()

	content := resultToString(sc.ScanResult)
	_, err = file.WriteString(content)
	if err != nil {
		log.Printf("scanner/SavScanResult: %v\n", err)
		return
	}
}

func resultToString(result map[string]LiveIPInfo) string {
	jByte, err := json.Marshal(result)
	if err != nil {
		log.Printf("scanner/resultToString: %v\n", err)
		return ""
	}
	return string(jByte)
}
