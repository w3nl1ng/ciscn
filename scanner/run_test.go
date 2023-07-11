package scanner

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	args := []string{"-sL", "16.163.13.0/24"}
	output := Run(args)
	if output == nil {
		t.Fatal("scanner/TestRun: sc.Run return nil")
	}
	fmt.Println(string(output))
}
