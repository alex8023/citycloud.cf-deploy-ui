package entity

import (
	"testing"
)

func TestBOSHTemplate(t *testing.T) {
	var test = NewSimpleBOSH("deployment-bosh")
	var bt = NewBOSHTemplate(test)
	_, err := bt.CreateBOSHYaml("/home/ubuntu/temp/bosh.yml")

	if err != nil {
		t.Error(err)
	}
}
