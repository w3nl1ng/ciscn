package scanner

import (
	"bufio"
	"github.com/panjf2000/ants/v2"
	"log"
	"os"
	"regexp"
	"sync"
)

var TempLiveIP []string
var Mu sync.Mutex

func insertToLiveIP(liveIps []string) {
	Mu.Lock()
	TempLiveIP = append(TempLiveIP, liveIps...)
	Mu.Unlock()
}

func workFunc(i interface{}) {
	ipSubnet, ok := i.(string)
	if !ok {
		log.Printf("scanner/workFunc: can not convert type(%T) to type(string)\n", i)
		return
	}
	log.Printf("scanner/workFunc: begin scanning %s\n", ipSubnet)
	args := []string{"-sn", ipSubnet}
	output := Run(args)
	liveIps := findIPv4Addresses(string(output))
	insertToLiveIP(liveIps)
	log.Printf("scanner/workFunc: finish scanning %s\n", ipSubnet)
}

// 此函数根据Scanner得IpFileName获取ip段，然后探测存活ip保存到LiveIP中
func (sc *Scanner) scanLiveIP() {
	ipListAll := parseIpFromFile(sc.IpFileName)

	var wg sync.WaitGroup
	p, err := ants.NewPoolWithFunc(10, func(i interface{}) {
		workFunc(i)
		wg.Done()
	})
	if err != nil {
		log.Printf("scanner/scanLiveIP: %v\n", err)
	}
	defer p.Release()

	for _, ipSubnet := range ipListAll {
		wg.Add(1)
		_ = p.Invoke(ipSubnet)
	}
	wg.Wait()

	sc.LiveIP = TempLiveIP
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
