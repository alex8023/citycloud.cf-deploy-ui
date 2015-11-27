package utils_test

import (
	"bytes"
	"fmt"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("Test SCP", func() {
		var out bytes.Buffer
		scpRunner := utils.NewDeploySSHRunner("10.162.2.146", "root", "4rfv%TGB")
		fmt.Println()
		fmt.Println("=============")
		scpRunner.Push("/home/ubuntu/deploy/jdk1.8.0_40", "~/", &out)
		line, _ := out.ReadString('\n')
		fmt.Println(line)
		//		fmt.Println("=============")
		//		scpRunner.Push("/home/ubuntu/deploy/abc/abcd", "~/", &out)
		//		line, _ = out.ReadString('\n')
		//		fmt.Println(line)
		//		fmt.Println("=============")
		//		scpRunner.Push("/home/ubuntu/deploy/abc/cdef", "~/", &out)
		//		line, _ = out.ReadString('\n')
		//		fmt.Println(line)

	})
})
