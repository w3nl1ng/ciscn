package scanner

import (
	"fmt"
	"testing"
)

func Test_findDeviceInfo(t *testing.T) {
	out := `Starting Nmap 7.80 ( https://nmap.org ) at 2023-07-12 13:31 CST
Nmap scan report for ec2-16-163-13-0.ap-east-1.compute.amazonaws.com (16.163.13.0)
Host is up (0.081s latency).
Not shown: 997 filtered ports
PORT    STATE  SERVICE
22/tcp  open   ssh
80/tcp  open   http
443/tcp closed https
Device type: general purpose|storage-misc|firewall|webcam
Running (JUST GUESSING): Linux 4.X|3.X|2.6.X (93%), Synology DiskStation Manager 5.X (86%), WatchGuard Fireware 11.X (86%), Tandberg embedded (86%), FreeBSD 6.X (85%)
OS CPE: cpe:/o:linux:linux_kernel:4.2 cpe:/o:linux:linux_kernel:3.8 cpe:/o:linux:linux_kernel:2.6.32 cpe:/o:linux:linux_kernel cpe:/a:synology:diskstation_manager:5.1 cpe:/o:watchguard:fireware:11.8 cpe:/h:tandberg:vcs cpe:/o:freebsd:freebsd:6.2
Aggressive OS guesses: Linux 4.2 (93%), Linux 3.8 (88%), Linux 4.4 (88%), Linux 2.6.32 (87%), Linux 2.6.32 or 3.10 (87%), Linux 3.5 (87%), Synology DiskStation Manager 5.1 (86%), Linux 2.6.35 (86%), Linux 2.6.39 (86%), Linux 3.10 - 3.12 (86%)
No exact OS matches for host (test conditions non-ideal).
OS detection performed. Please report any incorrect results at https://nmap.org/submit/ .
Nmap done: 1 IP address (1 host up) scanned in 25.09 seconds`
	result := findDeviceInfo(out)
	fmt.Println(result)
}

func Test_deviceScan(t *testing.T) {
	sc := NewScanner("iplist-test.txt")
	// sc.ScanLiveIP()
	sc.LiveIP = []string{"16.163.13.255", "16.163.13.251", "16.163.13.250", "16.163.13.249"}
	sc.PortScan()
	sc.deviceScan()

	fmt.Println(sc.ScanResult)
}
