package server

import (
	"ciscn/embed"
	"ciscn/scan/live/probe"
	"ciscn/utils"
	"fmt"
	"github.com/panjf2000/ants"
	"log"
	"strconv"
	"sync"
	"time"
)

var (
	mu       sync.Mutex
	LiveAddr []string
)

//	func FetchPortWorker2(targetsInterface interface{}) {
//		defer WG2.Done()
//		address := targetsInterface.(string)
//		if probe.CheckIfPortLive(address) {
//			log.Println(address)
//			mu.Lock()
//			LiveAddr = append(LiveAddr, address)
//			mu.Unlock()
//		}
//	}
//
// func FetchPortWorker(targetsInterface interface{}) {
//
//	defer WG.Done()
//	targets, ok := targetsInterface.([]string)
//	if !ok {
//		return
//	}
//	Fetch2, err := ants.NewPoolWithFunc(20, FetchPortWorker2)
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//
//	for _, target := range targets {
//		log.Printf("开始扫描 %s 的常用端口，数量为: %d", target, len(embed.CommonPort))
//		for _, port := range embed.CommonPort {
//			address := fmt.Sprintf("%s:%d", target, port)
//			WG2.Add(1)
//			if err := Fetch2.Invoke(address); err != nil {
//				log.Fatal(err.Error())
//			}
//		}
//	}
//	WG2.Wait()
//
// }

type Addr struct {
	Ip   string
	Port int
}

func StartLivedPortScan(address []string, thread int) []string {

	log.Println(len(address))
	splitAddress := utils.SplitTargetList(address, thread)

	Fetch, err := ants.NewPoolWithFunc(thread, FetchPortWorker)
	defer ants.Release()

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, hosts := range splitAddress {
		WG.Add(1)
		if err := Fetch.Invoke(hosts); err != nil {
			log.Fatal(err.Error())
		}
	}
	WG.Wait()

	return LiveAddr
}

func PortConnect(addr Addr, respondingHosts chan<- string, wg *sync.WaitGroup) {
	host, port := addr.Ip, addr.Port
	conn, err := probe.WrapperTcpWithTimeout("tcp4", fmt.Sprintf("%s:%v", host, port), time.Duration(1)*time.Second)
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	if err == nil {
		address := host + ":" + strconv.Itoa(port)
		result := fmt.Sprintf("%s open", address)
		fmt.Println(result)
		wg.Add(1)
		respondingHosts <- address
	}
}

func FetchPortWorker(i interface{}) {
	defer WG.Done()
	var AliveAddress []string
	var wg sync.WaitGroup
	hostslist := i.([]string)

	probePorts := embed.CommonPort
	workers := 16
	Addrs := make(chan Addr, len(hostslist)*len(probePorts))
	results := make(chan string, len(hostslist)*len(probePorts))

	go func() {
		for found := range results {
			AliveAddress = append(AliveAddress, found)
			wg.Done()
		}
	}()

	for i := 0; i < workers; i++ {
		go func() {
			for addr := range Addrs {
				PortConnect(addr, results, &wg)
				wg.Done()
			}
		}()
	}
	for _, port := range probePorts {
		//fmt.Printf("开始扫描%d端口 ,共有%d host\r\n", port, len(hostslist))
		for _, host := range hostslist {
			wg.Add(1)
			Addrs <- Addr{host, port}
		}
	}

	wg.Wait()
	close(Addrs)
	close(results)
	mu.Lock()
	LiveAddr = append(LiveAddr, AliveAddress...)
	mu.Unlock()
}
