package entity

import (
	"testing"
)

func TestCloudFoundryTemplate(t *testing.T) {
	var cloudFoundryProperties = NewCloudFoundryProperties("cf-release",
		"57cfc863-786d-4495-bb97-86d2f650a038", "192.168.133.102", "ccipaas.net", "cci")

	var compilation = NewCompilation("flavor_91", "zone2", 6, "cf1")

	var networks = NewNetWorks("cf1",
		"manual",
		"8bb21e6e-dc6a-409c-82d0-a110fb3c9fe1",
		"192.168.129.0/24",
		"10.10.170.2",
		"192.168.133.108",
		"192.168.129.1 - 192.168.129.99",
		"192.168.129.100 - 192.168.129.126")

	var floatingNetWork = NewFloatingNetWork("192.168.133.102")
	var mapNetWorks = make(map[string]NetWorks)
	mapNetWorks[networks.Name] = networks
	mapNetWorks[floatingNetWork.Name] = floatingNetWork

	var resourcesPool = NewResourcesPools("normal",
		"flavor_91",
		"zone2",
		"cf1",
		4)
	var mapResourcesPools = make(map[string]ResourcesPools)
	mapResourcesPools[resourcesPool.Name] = resourcesPool
	var cft = NewCloudFoundryTemplate(CloudFoundry{
		CloudFoundryProperties: cloudFoundryProperties,
		Compilation:            compilation,
		NetWorks:               mapNetWorks,
		ResourcesPools:         mapResourcesPools,
	})
	_, err := cft.CreateCloudFoundryV2Yaml("/home/ubuntu/temp/cloudfoundryv2.yml")

	if err != nil {
		t.Error(err)
	}
}

func TestCloudFoundryV3Template(t *testing.T) {
	var cloudFoundryProperties = NewCloudFoundryProperties("cf-release",
		"57cfc863-786d-4495-bb97-86d2f650a038", "192.168.133.102", "ccipaas.net", "cci")
	var cft = NewCloudFoundryTemplate(CloudFoundry{CloudFoundryProperties: cloudFoundryProperties})
	_, err := cft.CreateCloudFoundryV3Yaml("/home/ubuntu/temp/cloudfoundryv3.yml")

	if err != nil {
		t.Error(err)
	}
}
