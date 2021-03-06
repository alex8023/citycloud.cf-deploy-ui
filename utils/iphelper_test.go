package utils_test

import (
	"fmt"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
	. "github.com/onsi/ginkgo"
	"net"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("parse ip", func() {

		ip := net.ParseIP("192.168.133.108").To4()
		fmt.Println(len(ip))
		fmt.Println(ip.To4())

		for i := 0; i < len(ip); i++ {
			fmt.Println(ip[i])
		}
	})
	It("check c i d r contains i p", func() {

		cidr := "10.10.10.0/24"
		ip := "10.10.10.24"
		fmt.Println(utils.CheckCIDRContainsIP(cidr, ip))
	})
	It("parse static i p", func() {

		start := "192.168.133.108"
		end := "192.168.133.120"
		ips := utils.ParseStaticIP(start, end)

		fmt.Println(ips)
	})
	It("ass ip", func() {

		iPFactory := utils.NewIPFactory()
		iPFactory.InitFactory("192.168.1.1", "192.168.1.20")

		fmt.Println("iiiiiiiiiiiiiiiiii")
		fmt.Printf("%s", iPFactory.GetAssignaIP4Job(utils.Job_Cloud_Controller_Clock))
		fmt.Println()
		abs, e := iPFactory.AssignaIP2Job(utils.Job_Cloud_Controller_Clock, 2)
		fmt.Printf("AssignaIP %s , %s", abs, e)
		fmt.Println()
		fmt.Printf("%s", iPFactory.GetAssignaIP4Job(utils.Job_Cloud_Controller_Clock))
		fmt.Println()

		iPFactory.AssignaIP2Job(utils.Job_Cloud_Controller_Clock, 3)

		fmt.Printf("%s", iPFactory.GetAssignaIP4Job(utils.Job_Cloud_Controller_Clock))
		fmt.Println()

		iPFactory.AssignaIP2Job(utils.Job_Cloud_Controller_Clock, 1)

		fmt.Printf("%s", iPFactory.GetAssignaIP4Job(utils.Job_Cloud_Controller_Clock))
		fmt.Println()
	})
	It("splite", func() {

		str := utils.SpliteIPs("192.168.129.100 - 192.168.129.126")
		fmt.Println(str)
		for _, i := range str {
			fmt.Println(i)
		}
	})
})
