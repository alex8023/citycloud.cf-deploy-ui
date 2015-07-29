package utils

import (
	"fmt"
	"net"
	"testing"
)

func TestParseIp(testing *testing.T) {
	ip := net.ParseIP("192.168.133.108").To4()
	fmt.Println(len(ip))
	fmt.Println(ip.To4())

	for i := 0; i < len(ip); i++ {
		fmt.Println(ip[i])
	}

}

func TestCheckCIDRContainsIP(testing *testing.T) {
	cidr := "10.10.10.0/24"
	ip := "10.10.10.24"
	fmt.Println(CheckCIDRContainsIP(cidr, ip))
}

func TestParseStaticIP(testing *testing.T) {
	start := "192.168.133.108"
	end := "192.168.133.120"
	ips := parseStaticIP(start, end)

	fmt.Println(ips)
}

func TestAssIp(testing *testing.T) {
	iPFactory := NewIPFactory()
	iPFactory.InitFactory("192.168.1.1", "192.168.1.20")
	fmt.Printf(" ips \n = %s", iPFactory.ips)
	fmt.Printf(" \nippool \n = %s", iPFactory.ippool)
	fmt.Printf(" \njobMaps \n = %s", iPFactory.jobMaps)
	fmt.Println()
	fmt.Printf("%s", iPFactory.GetAssignaIP4Job(Job_Cloud_Controller_Clock))
	fmt.Println()
	abs, e := iPFactory.AssignaIP2Job(Job_Cloud_Controller_Clock, 2)
	fmt.Printf("AssignaIP %s , %s", abs, e)
	fmt.Println()
	fmt.Printf("%s", iPFactory.GetAssignaIP4Job(Job_Cloud_Controller_Clock))
	fmt.Println()

	iPFactory.AssignaIP2Job(Job_Cloud_Controller_Clock, 3)

	fmt.Printf("%s", iPFactory.GetAssignaIP4Job(Job_Cloud_Controller_Clock))
	fmt.Println()

	iPFactory.AssignaIP2Job(Job_Cloud_Controller_Clock, 1)

	fmt.Printf("%s", iPFactory.GetAssignaIP4Job(Job_Cloud_Controller_Clock))
	fmt.Println()
}

func TestSplite(testing *testing.T) {
	str := SpliteIPs("192.168.129.100 - 192.168.129.126")
	fmt.Println(str)
	for _, i := range str {
		fmt.Println(i)
	}
}
