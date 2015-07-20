package entity

import (
	"testing"
)

func TestCloudFoundryTemplate(t *testing.T) {
	var test = NewSimpleCloudFoundry("deployment-cloudfoundry")
	var cft = NewCloudFoundryTemplate(test)
	_, err := cft.CreateCloudFoundryV2Yaml("/home/ubuntu/temp/cloudfoundryv2.yml")

	if err != nil {
		t.Error(err)
	}
}

func TestCloudFoundryV3Template(t *testing.T) {
	var test = NewSimpleCloudFoundry("deployment-cloudfoundry")
	var cft = NewCloudFoundryTemplate(test)
	_, err := cft.CreateCloudFoundryV3Yaml("/home/ubuntu/temp/cloudfoundryv3.yml")

	if err != nil {
		t.Error(err)
	}
}
