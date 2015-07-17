package utils

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

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

func TestSetDeployMents(test *testing.T) {
	//	cmdRunner := NewDeployCmdRunner()
	//	deployment := "microbosh/micro_bosh.yml"
	//	dir := "/home/ubuntu/bosh-workspace/deploy"
	//	var out bytes.Buffer
	//	command := Command{Name: "bosh", Args: []string{"micro", "deployment", deployment}, Dir: dir, Stdin: ""}
	//	cmdRunner.RunCommandAsyncCmd(command, &out)
	//	ReadAndWriteBytes(&out, &cmdRunner)

	//	fmt.Println("Over")
}

func TestDeployMicrBOSH(test *testing.T) {

	//	dir := "/home/ubuntu/bosh-workspace/deploy"
	//	var out bytes.Buffer
	//	stemcells := "/home/ubuntu/bosh-workspace/stemcells/bosh-stemcell-2719-openstack-kvm-ubuntu-lucid-go_agent.tgz"
	//	microCommand := Command{Name: "bosh", Args: []string{"micro", "deploy", stemcells}, Dir: dir, Stdin: "yes"}
	//	//cmdRunner需要新建一个，否则后续的命令不会执行
	//	cmdRunner1 := NewDeployCmdRunner()
	//	cmdRunner1.RunCommandAsyncCmd(microCommand, &out)
	//	ReadAndWriteBytes(&out, &cmdRunner1)

	//	fmt.Println("Over")
}

func TestTargetMicrBOSH(test *testing.T) {
	dir := "/home/ubuntu/bosh-workspace/deploy"
	var out bytes.Buffer
	loginCommand := Command{Name: "bosh", Args: []string{"target", "https://192.168.133.108:25555"}, Dir: dir, Stdin: "admin\nadmin\n"}
	cmdRunner1 := NewDeployCmdRunner()
	cmdRunner1.RunCommandAsyncCmd(loginCommand, &out)
	ReadAndWriteBytes(&out, &cmdRunner1)
}

func TestLoginMicrBOSH(test *testing.T) {
	dir := "/home/ubuntu/bosh-workspace/deploy"
	var out bytes.Buffer
	loginCommand := Command{Name: "bosh", Args: []string{"login"}, Dir: dir, Stdin: "admin\nadmin\n"}
	cmdRunner1 := NewDeployCmdRunner()
	cmdRunner1.RunCommandAsyncCmd(loginCommand, &out)
	ReadAndWriteBytes(&out, &cmdRunner1)
}
