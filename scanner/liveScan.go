package scanner

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

// 此函数根据Scanner得IpFileName获取ip段，然后探测存活ip保存到LiveIP中
func (sc *Scanner) scanLiveIP() {
	//ipListAll := parseIpFromFile(sc.IpFileName)

}

func work(id int, jobs <-chan string, result chan<- []string) {
	//for ipSubnet := range jobs {
	//	args := []string{"-sn", ipSubnet}
	//	output := Run(args)
	//}
}

// 处理nmap -sn ip的输出，找到存活的ip
func findIPv4Addresses(s string) []string {
	re := regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)
	return re.FindAllString(s, -1)
}

// parseIpFromFile 函数将ip段从文件中读取出来
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
