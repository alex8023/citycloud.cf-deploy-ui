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
	//deployment := "/home/ubuntu/bosh-workspace/deploy/microbosh/micro_bosh.yml"
	//	go cmd.RunCommand("bosh", "", &out, "micro", "deployment", deployment)
	//stemcell := "/home/ubuntu/bosh-workspace/stemcells/bosh-stemcell-2719-openstack-kvm-ubuntu-lucid-go_agent.tgz"
	//	go cmd.RunCommand("bosh", "", &out, "micro", "deploy", stemcell)
	go cmd.RunCommand("sh", "yes", &out, "deploy.sh")
	fmt.Println("hello")
	for {
		time.Sleep(time.Second)
	again:
		abc, err := out.ReadBytes('\n')
		if err != nil && cmd.Success() {

			break
		} else {
			if err != nil {
				if abc != nil {
					fmt.Print(string(abc))
				}

				goto again
			}
			fmt.Print(string(abc))
		}
	}

}
