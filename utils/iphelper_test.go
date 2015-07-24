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
