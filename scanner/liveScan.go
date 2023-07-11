package scanner

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

// 此函数根据Scanner得IpFileName获取ip段，然后探测存活ip保存到LiveIP中
func (sc *Scanner) scanLiveIP() {
	ipListAll := parseIpFromFile(sc.IpFileName)

	workNumbers := len(ipListAll)
	jobs := make(chan string, workNumbers)
	results := make(chan []string, workNumbers)

	for i := 0; i < 8; i++ {
		go work(i, jobs, results)
	}

	for _, ip := range ipListAll {
		jobs <- ip
	}
	close(jobs)

	var liveIPs []string
	for i := 0; i < workNumbers; i++ {
		liveIP := <-results
		liveIPs = append(liveIPs, liveIP...)
	}

	sc.LiveIP = liveIPs
}

func work(id int, jobs <-chan string, result chan<- []string) {
	for ipSubnet := range jobs {
		log.Printf("scanner/work: worker %d is working\n", id)
		args := []string{"-sn", ipSubnet}
		output := Run(args)
		LiveIp := findIPv4Addresses(string(output))
		result <- LiveIp
	}
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
