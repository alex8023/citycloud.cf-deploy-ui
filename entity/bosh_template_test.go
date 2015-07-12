package entity

import (
	"testing"
)

func TestBOSHTemplate(t *testing.T){
	var test = NewBOSH("deployment-bosh")
	var bt = NewBOSHTemplate(test)
	_,err := bt.CreateBOSHYaml("/home/ubuntu/bosh.yml")
	
	if err != nil{
		t.Error(err)
	}
}