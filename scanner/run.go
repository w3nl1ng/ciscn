package scanner

import (
	"log"
	"os/exec"
)

// 此函数根据cmd参数调用os exec执行nmap，并返回执行的输出结果
func (sc *Scanner) Run(args []string) []byte {
	cmd := exec.Command("nmap", args...)
	output, err := cmd.Output()
	if err != nil {
		log.Println("scanner/Run: failed to run cmd")
		return nil
	}
	return output
}
