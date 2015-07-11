package entity

import (
	"testing"
)

func TestCloudFoundryTemplate(t *testing.T){
	var test = NewCloudFoundry("deployment-cloudfoundry")
	var cft = NewCloudFoundryTemplate(test)
	_,err := cft.CreateCloudFoundryYaml("/home/ubuntu/cloudfoundry.yml")
	
	if err != nil{
		t.Error(err)
	}
}