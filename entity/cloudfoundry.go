package entity

type CloudFoundry struct{
	Name string
}

type Compilation struct{
	
}

func NewCloudFoundry(name string)(cloudfoundry CloudFoundry){
	cloudfoundry.Name = name
	return
}