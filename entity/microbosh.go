package entity

type MicroBOSH struct {
	Name            string
	Network         NetWork
	Resources       Resources
	Cloudproperties CloudProperties
}

type NetWork struct {
	Vip    string
	Net_id string
}

type Resources struct {
	Persistent_disk   string
	Instance_type     string
	Availability_zone string
}

type CloudProperties struct {
	Auth_url          string
	Username          string
	Api_key           string
	Tenant            string
	Default_key_name  string
	Private_key       string
	Cci_ebs_url       string
	Cci_ebs_accesskey string
	Cci_ebs_secretkey string
}

func NewMicroBOSH(
	name string,
	network NetWork,
	resource Resources,
	cloudproperties CloudProperties,
) (microbosh MicroBOSH) {
	microbosh.Name = name
	microbosh.Network = network
	microbosh.Resources = resource
	microbosh.Cloudproperties = cloudproperties
	return
}

func NewNetWork(
	vip string,
	net_id string,
) (network NetWork) {
	network.Vip = vip
	network.Net_id = net_id
	return
}

func NewResources(
	persistent_disk string,
	instance_type string,
	availability_zone string,
) (resource Resources) {
	resource.Persistent_disk = persistent_disk
	resource.Instance_type = instance_type
	resource.Availability_zone = availability_zone
	return
}

func NewCloudProperties(
	auth_url string,
	username string,
	api_key string,
	tenant string,
	default_key_name string,
	private_key string,
	cci_ebs_url string,
	cci_ebs_accesskey string,
	cci_ebs_secretkey string,
) (cloudproperties CloudProperties) {
	cloudproperties.Auth_url = auth_url
	cloudproperties.Username = username
	cloudproperties.Api_key = api_key
	cloudproperties.Tenant = tenant
	cloudproperties.Default_key_name = default_key_name
	cloudproperties.Private_key = private_key
	cloudproperties.Cci_ebs_url = cci_ebs_url
	cloudproperties.Cci_ebs_accesskey = cci_ebs_accesskey
	cloudproperties.Cci_ebs_secretkey = cci_ebs_secretkey
	return
}
