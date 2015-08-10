package entity

import (
	"github.com/astaxie/beego/orm"
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
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

func (microBOSH *MicroBOSH) Load() error {
	queryOneErr := orm.NewOrm().QueryTable(microBOSH).One(microBOSH, "Id")
	if queryOneErr != nil {
		logger.Error("Query One failed %s", queryOneErr)
	}

	errors := orm.NewOrm().Read(microBOSH, "Id")
	if errors != nil {
		logger.Error("Read MicroBOSH error : %s", errors)
		_, err := orm.NewOrm().Insert(microBOSH)
		if err != nil {
			logger.Error("Inert MicroBOSH error %s ", err)
		}
	}

	return nil
}

func (microBOSH *MicroBOSH) Update() error {
	_, err := orm.NewOrm().Update(microBOSH)
	if err != nil {
		logger.Error("Update MicroBOSH error %s ", err)
	}
	return err
}

func (netWork *NetWork) Load() error {
	queryOneErr := orm.NewOrm().QueryTable(netWork).One(netWork, "Id")
	if queryOneErr != nil {
		logger.Error("Query One failed %s", queryOneErr)
	}
	errors := orm.NewOrm().Read(netWork, "Id")
	if errors != nil {
		logger.Error("Read NetWork error : %s", errors)
		_, err := orm.NewOrm().Insert(netWork)
		if err != nil {
			logger.Error("Inert NetWork error %s ", err)
		}
	}
	return errors
}

func (netWork *NetWork) Update() error {
	_, err := orm.NewOrm().Update(netWork)
	if err != nil {
		logger.Error("Update NetWork error %s ", err)
	}
	return err
}

func (cloudProperties *CloudProperties) Load() error {

	queryOneErr := orm.NewOrm().QueryTable(cloudProperties).One(cloudProperties, "Id")
	if queryOneErr != nil {
		logger.Error("Query One failed %s", queryOneErr)
	}

	errors := orm.NewOrm().Read(cloudProperties, "Id")
	if errors != nil {
		logger.Error("Read CloudProperties error : %s", errors)
		_, err := orm.NewOrm().Insert(cloudProperties)
		if err != nil {
			logger.Error("Inert CloudProperties error %s ", err)
		}
	}
	return errors
}

func (cloudProperties *CloudProperties) Update() error {
	_, err := orm.NewOrm().Update(cloudProperties)
	if err != nil {
		logger.Error("Update CloudProperties error %s ", err)
	}
	return err
}

func (resources *Resources) Load() error {
	queryOneErr := orm.NewOrm().QueryTable(resources).One(resources, "Id")
	if queryOneErr != nil {
		logger.Error("Query One failed %s", queryOneErr)
	}
	errors := orm.NewOrm().Read(resources, "Id")
	if errors != nil {
		logger.Error("Read Resources error : %s", errors)
		_, err := orm.NewOrm().Insert(resources)
		if err != nil {
			logger.Error("Inert Resources error %s ", err)
		}
	}
	return errors
}

func (resources *Resources) Update() error {
	_, err := orm.NewOrm().Update(resources)
	if err != nil {
		logger.Error("Update Resources error %s ", err)
	}
	return err
}

func init() {
	orm.RegisterModel(new(NetWork), new(CloudProperties), new(MicroBOSH), new(Resources))
}
