// 此package负责解析命令行flag
package cmd

import "flag"

var IpFileName string

func Flag() {
	flag.StringVar(&IpFileName, "f", "", "specify the iplist filename")
	flag.Parse()
}
