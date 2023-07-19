package embed

import "testing"

func TestCommonPort(t *testing.T) {
	for _, port := range CommonPort {
		println(port)
	}
}
