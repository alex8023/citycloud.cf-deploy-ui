package entity

type BOSH struct{
	Name string
	Uuid string
	InstanceType string
	PowerDns string
	Director string
	MicroBOSHIP string
	CloudProperties CloudProperties
}

func NewBOSH(name,
uuid,
instanceType,
powerdns,
director ,
microboshIp string,
cloudProperties CloudProperties)(bosh BOSH){
	bosh.Name = name
	bosh.Uuid = uuid
	bosh.InstanceType = instanceType
	bosh.PowerDns = powerdns
	bosh.Director = director
	bosh.MicroBOSHIP =microboshIp
	bosh.CloudProperties = cloudProperties
	return
}