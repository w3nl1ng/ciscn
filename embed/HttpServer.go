package embed

import (
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed HttpServer.json
var httpJson []byte
var HttpServers []HttpServer

type HttpServer struct {
	Pattern     string      `json:"pattern"`
	Name        string      `json:"name"`
	PatternFlag string      `json:"pattern_flag"`
	VersionInfo Versioninfo `json:"versioninfo"`
}

func init() {
	err := json.Unmarshal(httpJson, &HttpServers)
	if err != nil {
		log.Fatal(err)
	}
}
