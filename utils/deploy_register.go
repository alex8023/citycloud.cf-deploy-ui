package utils

import ()

const (
	DeployMicroBOSH    string = "MicroBOSH"
	DeployBOSH         string = "BOSH"
	DeployCloudFoundry string = "CloudFoundry"
)

const (
	DeployInit = iota
	Deploying
	DeployDone
)

type DeployStat struct {
	Name   string
	Status int
}

var adapters = make(map[string]DeployStat)

func Register(name string, adapter DeployStat) (bool, error) {

	if _, ok := adapters[name]; ok {
	}
	adapters[name] = adapter
	return false, nil
}
