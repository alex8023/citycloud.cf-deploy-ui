package utils

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestRunCommand(test *testing.T) {
	cmd := NewDeployCmdRunner()
	var out bytes.Buffer
	go cmd.RunCommand("ping", "", &out, "127.0.0.1", "-c 10")
	fmt.Println("hello")
	for {
		time.Sleep(time.Second)
	again:
		abc, err := out.ReadBytes('\n')
		if err != nil && cmd.Success() {
			break
		} else {
			if err != nil {
				goto again
			}
			fmt.Print(string(abc))
		}
	}

}
