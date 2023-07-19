package embed

import (
	"fmt"
	"testing"
)

func TestHttpServer(t *testing.T) {
	for i, httpServer := range HttpServers {
		fmt.Println(i)
		fmt.Println(httpServer.Name)
		fmt.Println(httpServer.Pattern)
		fmt.Println(httpServer.VersionInfo.Vendorproductname)
		fmt.Println(httpServer.PatternFlag)
	}
}
