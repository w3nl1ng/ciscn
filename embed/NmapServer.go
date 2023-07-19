package embed

import (
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed nmap.json
var nmapJson []byte
var NmapServers []TooLTT

type TooLTT struct {
	Protocol     string        `json:"protocol"`
	Probename    string        `json:"probename"`
	Probestring  string        `json:"probestring"`
	Ports        []interface{} `json:"ports"`
	Sslports     []interface{} `json:"sslports"`
	Totalwaitms  string        `json:"totalwaitms"`
	Tcpwrappedms string        `json:"tcpwrappedms"`
	Rarity       string        `json:"rarity"`
	Fallback     string        `json:"fallback"`
	Matches      []Matches     `json:"matches"`
	Softmatches  []Softmatches `json:"softmatches"`
}
type Versioninfo struct {
	Cpename           string `json:"cpename"`
	Devicetype        string `json:"devicetype"`
	Hostname          string `json:"hostname"`
	Info              string `json:"info"`
	Operatingsystem   string `json:"operatingsystem"`
	Vendorproductname string `json:"vendorproductname"`
	Version           string `json:"version"`
}
type Matches struct {
	Pattern     string      `json:"pattern"`
	Name        string      `json:"name"`
	PatternFlag string      `json:"pattern_flag"`
	Versioninfo Versioninfo `json:"versioninfo"`
}
type Softmatches struct {
	Pattern     string      `json:"pattern"`
	Name        string      `json:"name"`
	PatternFlag string      `json:"pattern_flag"`
	Versioninfo Versioninfo `json:"versioninfo"`
}

func init() {
	err := json.Unmarshal(nmapJson, &NmapServers)
	if err != nil {
		log.Fatal(err)
	}
}
