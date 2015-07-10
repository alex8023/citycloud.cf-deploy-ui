package entity

type MicroBOSH struct {
	name            string
	network         NetWork
	resources       Resources
	cloudproperties CloudProperties
}

type NetWork struct {
	vip    string
	net_id string
}

type Resources struct {
	persistent_disk   string
	instance_type     string
	availability_zone string
}

type CloudProperties struct {
	auth_url          string
	username          string
	api_key           string
	tenant            string
	default_key_name  string
	private_key       string
	cci_ebs_url       string
	cci_ebs_accesskey string
	cci_ebs_secretkey string
}

func NewMicroBOSH(
	name string,
	network NetWork,
	resource Resources,
	cloudproperties CloudProperties,
) (microbosh MicroBOSH) {
	microbosh.name = name
	microbosh.network = network
	microbosh.resources = resource
	microbosh.cloudproperties = cloudproperties
	return
}

func NewNetWork(
	vip string,
	net_id string,
) (network NetWork) {
	network.vip = vip
	network.net_id = net_id
	return
}

func NewResources(
	persistent_disk string,
	instance_type string,
	availability_zone string,
) (resource Resources) {
	resource.persistent_disk = persistent_disk
	resource.instance_type = instance_type
	resource.availability_zone = availability_zone
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
	cloudproperties.auth_url = auth_url
	cloudproperties.username = username
	cloudproperties.api_key = api_key
	cloudproperties.tenant = tenant
	cloudproperties.default_key_name = default_key_name
	cloudproperties.private_key = private_key
	cloudproperties.cci_ebs_url = cci_ebs_url
	cloudproperties.cci_ebs_accesskey = cci_ebs_accesskey
	cloudproperties.cci_ebs_secretkey = cci_ebs_secretkey
	return
}
