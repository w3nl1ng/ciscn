package main

import (
	"ciscn/cmd"
	"ciscn/scanner"
	"log"
	"os"
)

func main() {
	cmd.Flag()
	sc := scanner.NewScanner(cmd.IpFileName)
	if sc == nil {
		log.Println("main/main: failed to new a scanner")
		os.Exit(0)
	}
}
