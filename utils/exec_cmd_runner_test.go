package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
	"time"
)

func TestRunCommand(test *testing.T) {
	//	cmdRunner := NewExecCmdRunner()
	//	stdout, stderr, _, _ := cmdRunner.RunCommand("ping", "127.0.0.1", "-c 2")
	//	test.Log(stdout, stderr)
}

func TestRunCommand1(test *testing.T) {
	//	cmd := exec.Command("ping", "127.0.0.1", "-c 1")
	//	var out bytes.Buffer
	//	cmd.Stdout = &out
	//	cmd.Start()
	//	cd := make([]byte, 100)

	//	for {
	//		time.Sleep(2 * time.Second)
	//		abc, err := out.Read(cd)
	//		if err == nil {
	//			fmt.Println(string(cd)) //读取输出的日志
	//		}
	//		if abc == 0 {
	//			break
	//		}
	//	}
}

func TestRunAync(test *testing.T) {
	cmd := exec.Command("ping", "127.0.0.1", "-c 20")
	time.Sleep(time.Second)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	cmd.Start()
	fmt.Println(cmd.Process.Pid)

	for {
		time.Sleep(time.Second)
		abc, err := out.ReadBytes('\n')
		if err == nil {
			fmt.Print(string(abc)) //读取输出的日志
		}
		if err != nil {
			fmt.Println(err)
			//			break
		}
	}
	cmd.Wait()
}
