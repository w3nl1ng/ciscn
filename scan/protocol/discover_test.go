package promotol

import (
	"fmt"
	"testing"
)

func TestDiscover(t *testing.T) {
	tests := []struct {
		host string
		port int
	}{
		//{"211.22.90.156", 22},
		//{"211.22.90.175:22", 22},
		//{"211.22.90.156", 80},
		//{"120.53.91.189", 80},
		//{"154.23.86.62", 443},
		//{"18.166.56.248", 80},
		{"103.252.118.139", 80},
		//{"16.163.13.0", 9999},
		//{"16.163.13.0", 8888},
		//{"16.163.13.0", 23},
		//{"16.163.13.0", 1234},
	}

	args := map[string]interface{}{
		"FlagTimeout": 2,
		"FlagType":    "tcp",
		"FlagMode":    "",
		"FlagUrl":     "",
	}

	for _, target := range tests {
		resp := DiscoverTcp(target.host, target.port, args)

		//resp.ServiceApp = common.ExtractStrings(resp.ServiceApp.(string))
		//resp.ServiceApp = common.MergeDuplicates(resp.ServiceApp.([]string))
		//for _, s := range resp.ServiceApp.([]string) {
		//	if s == "Nginx" {
		//
		//	}
		//}
		fmt.Println(resp)
	}
}
