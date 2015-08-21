package utils_test

import (
	"fmt"

	"bytes"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
	. "github.com/onsi/ginkgo"
	"time"
)

var _ = Describe("Testing with Ginkgo", func() {
})

func ReadAndWriteBytes(out *bytes.Buffer, cmdRunner *utils.DeployCmdRunner) {
	for {
	loop:
		if out.Len() == 0 {
			if cmdRunner.Exited() {
				break
			} else {
				time.Sleep(time.Second)
				goto loop
			}
		}

	again:
		line, _ := out.ReadBytes('\n')
		if line != nil {
			fmt.Println(string(line))
			goto again
		} else {
			goto loop
		}
	}
}
