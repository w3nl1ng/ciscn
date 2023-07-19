package probe

import (
	"fmt"
	"testing"
)

func TestIsHFishfHoneyPot(t *testing.T) {
	target := "http://webapptestbi.f1071216.k8s.pupuvip.com/"
	//target = strings.TrimSuffix(target, "/")
	//fmt.Println(strings.Split(target, "//")[1])

	fmt.Println(IsHFishfHoneyPot(target))
}
