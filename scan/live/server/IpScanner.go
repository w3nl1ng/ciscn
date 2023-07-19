package server

import (
	"ciscn/scan/live/listener"
	"ciscn/scan/live/probe"
	"ciscn/utils"
	"github.com/panjf2000/ants"
	"log"
)

var LiveIP []string

func FetchIPWorker(targetsInterface interface{}) {
	defer WG.Done()
	targets, ok := targetsInterface.([]string)
	if !ok {
		return
	}

	for _, target := range targets {
		if probe.CheckIfIPLive(target) {
			LiveIP = append(LiveIP, target)
		}
	}
}

func StartLivedIPScan(hosts []string, thread int) []string {

	icmpListener := listener.NewIcmpListener()
	defer icmpListener.Cancel()

	splitHosts := utils.SplitTargetList(hosts, thread)

	Fetch, err := ants.NewPoolWithFunc(thread, FetchIPWorker)
	defer ants.Release()

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, hosts := range splitHosts {
		WG.Add(1)
		if err := Fetch.Invoke(hosts); err != nil {
			log.Fatal(err.Error())
		}
	}
	WG.Wait()

	return LiveIP
}
