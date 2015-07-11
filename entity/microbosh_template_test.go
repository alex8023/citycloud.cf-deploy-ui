package entity

import (
	"testing"
)



func TestMicroboshTemplate(t *testing.T){
	var test = NewMicroBOSH("deployment-microbosh",
	NewNetWork("vip","netid"),
	NewResources("16384","flavor_100","zone2"),
	NewCloudProperties("auth_url","username","apikey","tenant","defaultkeyname","privatekey","ebsurl","ebskey","ebssercetkey"))
	
	var mt = NewMicroBOSHTemplate(test)
	_,err := mt.CreateMicroBOSHYaml("/home/ubuntu/microbosh.yml")
	
	if err != nil{
		t.Error(err)
	}
}