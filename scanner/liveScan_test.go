package scanner

import (
	"fmt"
	"testing"
)

func Test_parseIpFromFile(t *testing.T) {
	ipList := parseIpFromFile("..\\iplist.txt")
	fmt.Println(ipList)
}

func Test_findIPv4Addresses(t *testing.T) {
	output := `Starting Nmap 7.80 ( https://nmap.org ) at 2023-07-11 16:27 CST
Nmap scan report for ec2-16-163-13-0.ap-east-1.compute.amazonaws.com (16.163.13.0)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-7.ap-east-1.compute.amazonaws.com (16.163.13.7)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-8.ap-east-1.compute.amazonaws.com (16.163.13.8)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-11.ap-east-1.compute.amazonaws.com (16.163.13.11)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-12.ap-east-1.compute.amazonaws.com (16.163.13.12)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-13.ap-east-1.compute.amazonaws.com (16.163.13.13)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-16.ap-east-1.compute.amazonaws.com (16.163.13.16)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-17.ap-east-1.compute.amazonaws.com (16.163.13.17)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-23.ap-east-1.compute.amazonaws.com (16.163.13.23)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-24.ap-east-1.compute.amazonaws.com (16.163.13.24)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-25.ap-east-1.compute.amazonaws.com (16.163.13.25)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-26.ap-east-1.compute.amazonaws.com (16.163.13.26)
Host is up (0.25s latency).
Nmap scan report for ec2-16-163-13-29.ap-east-1.compute.amazonaws.com (16.163.13.29)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-31.ap-east-1.compute.amazonaws.com (16.163.13.31)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-42.ap-east-1.compute.amazonaws.com (16.163.13.42)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-44.ap-east-1.compute.amazonaws.com (16.163.13.44)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-46.ap-east-1.compute.amazonaws.com (16.163.13.46)
Host is up (0.25s latency).
Nmap scan report for ec2-16-163-13-48.ap-east-1.compute.amazonaws.com (16.163.13.48)
Host is up (0.25s latency).
Nmap scan report for ec2-16-163-13-49.ap-east-1.compute.amazonaws.com (16.163.13.49)
Host is up (0.25s latency).
Nmap scan report for ec2-16-163-13-54.ap-east-1.compute.amazonaws.com (16.163.13.54)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-57.ap-east-1.compute.amazonaws.com (16.163.13.57)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-60.ap-east-1.compute.amazonaws.com (16.163.13.60)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-63.ap-east-1.compute.amazonaws.com (16.163.13.63)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-64.ap-east-1.compute.amazonaws.com (16.163.13.64)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-68.ap-east-1.compute.amazonaws.com (16.163.13.68)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-75.ap-east-1.compute.amazonaws.com (16.163.13.75)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-76.ap-east-1.compute.amazonaws.com (16.163.13.76)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-79.ap-east-1.compute.amazonaws.com (16.163.13.79)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-81.ap-east-1.compute.amazonaws.com (16.163.13.81)
Host is up (0.25s latency).
Nmap scan report for ec2-16-163-13-84.ap-east-1.compute.amazonaws.com (16.163.13.84)
Host is up (0.25s latency).
Nmap scan report for ec2-16-163-13-91.ap-east-1.compute.amazonaws.com (16.163.13.91)
Host is up (0.25s latency).
Nmap scan report for ec2-16-163-13-93.ap-east-1.compute.amazonaws.com (16.163.13.93)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-106.ap-east-1.compute.amazonaws.com (16.163.13.106)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-109.ap-east-1.compute.amazonaws.com (16.163.13.109)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-112.ap-east-1.compute.amazonaws.com (16.163.13.112)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-113.ap-east-1.compute.amazonaws.com (16.163.13.113)
Host is up (0.25s latency).
Nmap scan report for ec2-16-163-13-114.ap-east-1.compute.amazonaws.com (16.163.13.114)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-115.ap-east-1.compute.amazonaws.com (16.163.13.115)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-122.ap-east-1.compute.amazonaws.com (16.163.13.122)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-126.ap-east-1.compute.amazonaws.com (16.163.13.126)
Host is up (0.25s latency).
Nmap scan report for ec2-16-163-13-127.ap-east-1.compute.amazonaws.com (16.163.13.127)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-131.ap-east-1.compute.amazonaws.com (16.163.13.131)
Host is up (0.25s latency).
Nmap scan report for ec2-16-163-13-138.ap-east-1.compute.amazonaws.com (16.163.13.138)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-139.ap-east-1.compute.amazonaws.com (16.163.13.139)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-141.ap-east-1.compute.amazonaws.com (16.163.13.141)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-142.ap-east-1.compute.amazonaws.com (16.163.13.142)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-145.ap-east-1.compute.amazonaws.com (16.163.13.145)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-150.ap-east-1.compute.amazonaws.com (16.163.13.150)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-152.ap-east-1.compute.amazonaws.com (16.163.13.152)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-154.ap-east-1.compute.amazonaws.com (16.163.13.154)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-155.ap-east-1.compute.amazonaws.com (16.163.13.155)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-156.ap-east-1.compute.amazonaws.com (16.163.13.156)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-159.ap-east-1.compute.amazonaws.com (16.163.13.159)
Host is up (0.25s latency).
Nmap scan report for ec2-16-163-13-163.ap-east-1.compute.amazonaws.com (16.163.13.163)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-164.ap-east-1.compute.amazonaws.com (16.163.13.164)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-165.ap-east-1.compute.amazonaws.com (16.163.13.165)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-169.ap-east-1.compute.amazonaws.com (16.163.13.169)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-170.ap-east-1.compute.amazonaws.com (16.163.13.170)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-171.ap-east-1.compute.amazonaws.com (16.163.13.171)
Host is up (0.25s latency).
Nmap scan report for ec2-16-163-13-172.ap-east-1.compute.amazonaws.com (16.163.13.172)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-173.ap-east-1.compute.amazonaws.com (16.163.13.173)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-181.ap-east-1.compute.amazonaws.com (16.163.13.181)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-182.ap-east-1.compute.amazonaws.com (16.163.13.182)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-183.ap-east-1.compute.amazonaws.com (16.163.13.183)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-187.ap-east-1.compute.amazonaws.com (16.163.13.187)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-191.ap-east-1.compute.amazonaws.com (16.163.13.191)
Host is up (0.25s latency).
Nmap scan report for ec2-16-163-13-194.ap-east-1.compute.amazonaws.com (16.163.13.194)
Host is up (0.25s latency).
Nmap scan report for ec2-16-163-13-195.ap-east-1.compute.amazonaws.com (16.163.13.195)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-196.ap-east-1.compute.amazonaws.com (16.163.13.196)
Host is up (0.25s latency).
Nmap scan report for ec2-16-163-13-197.ap-east-1.compute.amazonaws.com (16.163.13.197)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-199.ap-east-1.compute.amazonaws.com (16.163.13.199)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-202.ap-east-1.compute.amazonaws.com (16.163.13.202)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-204.ap-east-1.compute.amazonaws.com (16.163.13.204)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-205.ap-east-1.compute.amazonaws.com (16.163.13.205)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-207.ap-east-1.compute.amazonaws.com (16.163.13.207)
Host is up (0.25s latency).
Nmap scan report for ec2-16-163-13-211.ap-east-1.compute.amazonaws.com (16.163.13.211)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-212.ap-east-1.compute.amazonaws.com (16.163.13.212)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-216.ap-east-1.compute.amazonaws.com (16.163.13.216)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-220.ap-east-1.compute.amazonaws.com (16.163.13.220)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-223.ap-east-1.compute.amazonaws.com (16.163.13.223)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-226.ap-east-1.compute.amazonaws.com (16.163.13.226)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-227.ap-east-1.compute.amazonaws.com (16.163.13.227)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-230.ap-east-1.compute.amazonaws.com (16.163.13.230)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-233.ap-east-1.compute.amazonaws.com (16.163.13.233)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-234.ap-east-1.compute.amazonaws.com (16.163.13.234)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-237.ap-east-1.compute.amazonaws.com (16.163.13.237)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-239.ap-east-1.compute.amazonaws.com (16.163.13.239)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-242.ap-east-1.compute.amazonaws.com (16.163.13.242)
Host is up (0.25s latency).
Nmap scan report for ec2-16-163-13-244.ap-east-1.compute.amazonaws.com (16.163.13.244)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-245.ap-east-1.compute.amazonaws.com (16.163.13.245)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-246.ap-east-1.compute.amazonaws.com (16.163.13.246)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-247.ap-east-1.compute.amazonaws.com (16.163.13.247)
Host is up (0.23s latency).
Nmap scan report for ec2-16-163-13-248.ap-east-1.compute.amazonaws.com (16.163.13.248)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-249.ap-east-1.compute.amazonaws.com (16.163.13.249)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-250.ap-east-1.compute.amazonaws.com (16.163.13.250)
Host is up (0.24s latency).
Nmap scan report for ec2-16-163-13-251.ap-east-1.compute.amazonaws.com (16.163.13.251)
Host is up (0.25s latency).
Nmap scan report for ec2-16-163-13-255.ap-east-1.compute.amazonaws.com (16.163.13.255)
Host is up (0.24s latency).
Nmap done: 256 IP addresses (97 hosts up) scanned in 11.96 seconds`
	LiveIp := findIPv4Addresses(output)
	fmt.Println(LiveIp)
}

func Test_scanLiveIP(t *testing.T) {
	sc := NewScanner("iplist-test.txt")
	sc.ScanLiveIP()
	fmt.Println(sc.LiveIP)
}
