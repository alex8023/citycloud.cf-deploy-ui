package entity

type CloudFoundry struct{
	Name string
}

func NewCloudFoundry(name string)(cloudfoundry CloudFoundry){
	cloudfoundry.Name = name
	return
}