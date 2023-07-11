package scanner

import (
	"bufio"
	"log"
	"os"
)

// 此函数根据Scanner得IpFileName获取ip段，然后探测存活ip保存到LiveIP中
func (sc *Scanner) scanLiveIP() {

}

func parseIpFromFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("scanner/parseIpFromFile: %v\n", err)
		return nil
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("scanner/parseIpFromFile: %v\n", err)
		}
	}()

	var result []string
	textScanner := bufio.NewScanner(file)
	for textScanner.Scan() {
		result = append(result, textScanner.Text())
	}

	return result
}
