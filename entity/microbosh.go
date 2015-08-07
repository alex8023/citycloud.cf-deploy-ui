package entity

import (
	"github.com/astaxie/beego/orm"
)

type MicroBOSH struct {
	Id              int64
	Name            string
	NetWork         NetWork         `orm:"-"`
	Resources       Resources       `orm:"-"`
	CloudProperties CloudProperties `orm:"-"`
}

type NetWork struct {
	Id    int64
	Vip   string
	NetId string
}

type Resources struct {
	Id               int64
	PersistentDisk   string
	InstanceType     string
	AvailabilityZone string
}

type CloudProperties struct {
	Id              int64
	AuthUrl         string
	UserName        string
	ApiKey          string
	Tenant          string
	DefaultKeyName  string
	PrivateKey      string
	CciEbsUrl       string `orm:"null"`
	CciEbsAccesskey string `orm:"null"`
	CciEbsSecretkey string `orm:"null"`
}

func NewMicroBOSH(
	name string,
	network NetWork,
	resource Resources,
	cloudproperties CloudProperties,
) (microbosh MicroBOSH) {
	microbosh.Name = name
	microbosh.NetWork = network
	microbosh.Resources = resource
	microbosh.CloudProperties = cloudproperties
	return
}

func NewNetWork(
	vip string,
	net_id string,
) (network NetWork) {
	network.Vip = vip
	network.NetId = net_id
	return
}

func NewResources(
	persistent_disk string,
	instance_type string,
	availability_zone string,
) (resource Resources) {
	resource.PersistentDisk = persistent_disk
	resource.InstanceType = instance_type
	resource.AvailabilityZone = availability_zone
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
	cloudproperties.AuthUrl = auth_url
	cloudproperties.UserName = username
	cloudproperties.ApiKey = api_key
	cloudproperties.Tenant = tenant
	cloudproperties.DefaultKeyName = default_key_name
	cloudproperties.PrivateKey = private_key
	cloudproperties.CciEbsUrl = cci_ebs_url
	cloudproperties.CciEbsAccesskey = cci_ebs_accesskey
	cloudproperties.CciEbsSecretkey = cci_ebs_secretkey
	return
}

func (microbosh *MicroBOSH) TableName() string {
	return "micro_bosh"
}

func init() {
	orm.RegisterModel(new(NetWork), new(CloudProperties), new(MicroBOSH))
}
