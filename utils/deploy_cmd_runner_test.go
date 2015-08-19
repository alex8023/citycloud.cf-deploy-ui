package utils

import (
	"fmt"

	"bytes"
	. "github.com/onsi/ginkgo"
	"time"
)

var _ = Describe("Testing with Ginkgo", func() {
})

func ReadAndWriteBytes(out *bytes.Buffer, cmdRunner *DeployCmdRunner) {
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
