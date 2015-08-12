package entity

import (
	"github.com/astaxie/beego/orm"
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
	"strings"
)

type CloudFoundry struct {
	CloudFoundryProperties CloudFoundryProperties
	Compilation            Compilation
	NetWorks               map[string]NetWorks
	ResourcesPools         []ResourcesPools
	CloudFoundryJobs       map[string]CloudFoundryJobs
	Properties             Properties
	Release                Release
	Stemcells              Stemcells
}

// CloudFoundry 基本属性配置
type CloudFoundryProperties struct {
	Id              int64
	Name            string
	Uuid            string
	FloatingIp      string
	SystemDomain    string
	SystemDomainOrg string
	CloudProperties CloudProperties `orm:"-"`
}

// 编译机器配置
type Compilation struct {
	Id               int64
	InstanceType     string
	AvailabilityZone string
	Workers          int
	DefaultNetWork   string
}

//网络配置
//BOSHDns Readonly
//此对象仅配置一个网络，IaaS2.0使用默认网络和floating网络，IaaS3.0使用多个网络
//在CloudFoundry中使用map组装网络资源
//默认情况下，允许定义一个私有网络和一个共有网络，IaaS2.0使用的共有网络为floating网络
type NetWorks struct {
	NetWorkName string `orm:"pk"`
	Name        string
	NetType     string
	NetId       string
	Cidr        string
	Dns         string
	PowerDns    string
	ReservedIp  string
	StaticIp    string
}

//资源池配置
//此对象仅配置一个资源池对象
//在CloudFoundry中使用map组装资源池
type ResourcesPools struct {
	Id               int64
	Name             string
	InstanceType     string
	AvailabilityZone string
	Size             int
	DefaultNetWork   string
}

type CloudFoundryJobs struct {
	Name          string
	JobName       string `orm:"pk"`
	ResourcesPool string
	Instances     int      `orm:"default(1)"`
	StaticIp      []string `orm:"-"`
	StaticIps     string   `orm:"size(2048)"`
}

//Job属性配置
type Properties struct {
}

type Release struct {
	Name    string
	Version string
}

type Stemcells struct {
	Name    string
	Version string
}

func NewCloudFoundry(
	cloudFoundryProperties CloudFoundryProperties,
	compilation Compilation,
	netWorks map[string]NetWorks,
	resourcesPools []ResourcesPools,
	cloudFoundryJobs map[string]CloudFoundryJobs,
	properties Properties) (cloudfoundry CloudFoundry) {
	cloudfoundry.CloudFoundryProperties = cloudFoundryProperties
	cloudfoundry.Compilation = compilation
	cloudfoundry.NetWorks = netWorks
	cloudfoundry.ResourcesPools = resourcesPools
	cloudfoundry.CloudFoundryJobs = cloudFoundryJobs
	cloudfoundry.Properties = properties
	return
}

func NewCloudFoundrySimple() CloudFoundry {
	return CloudFoundry{}
}

func NewCloudFoundryProperties(
	name string,
	uuid string,
	floatingIp string,
	systemDomain string,
	systemDomainOrg string) (cloudFoundryProperties CloudFoundryProperties) {
	cloudFoundryProperties.Name = name
	cloudFoundryProperties.Uuid = uuid
	cloudFoundryProperties.FloatingIp = floatingIp
	cloudFoundryProperties.SystemDomain = systemDomain
	cloudFoundryProperties.SystemDomainOrg = systemDomainOrg
	return
}

func NewCompilation(
	instanceType,
	availabilityZone string,
	workers int,
	defaultNetWork string) (compilation Compilation) {
	compilation.AvailabilityZone = availabilityZone
	compilation.InstanceType = instanceType
	compilation.Workers = workers
	compilation.DefaultNetWork = defaultNetWork
	return
}

func NewNetWorks(
	name,
	netWorkName,
	netType,
	netId,
	cidr,
	dns,
	powerDns,
	reservedIp,
	staticIp string) (netWorks NetWorks) {
	netWorks.Name = name
	netWorks.NetWorkName = netWorkName
	netWorks.NetType = netType
	netWorks.NetId = netId
	netWorks.Cidr = cidr
	netWorks.Dns = dns
	netWorks.PowerDns = powerDns
	netWorks.ReservedIp = reservedIp
	netWorks.StaticIp = staticIp
	return
}

func NewResourcesPools(
	name,
	instanceType,
	availabilityZone,
	defaultNetWork string,
	size int) (resourcesPools ResourcesPools) {
	resourcesPools.Name = name
	resourcesPools.InstanceType = instanceType
	resourcesPools.AvailabilityZone = availabilityZone
	resourcesPools.DefaultNetWork = defaultNetWork
	resourcesPools.Size = size
	return
}

func NewCloudFoundryJobs(name, jobName, resourcePool string, instacnes int, staticIp []string) (cloudFoundryJobs CloudFoundryJobs) {
	cloudFoundryJobs.Name = name
	cloudFoundryJobs.JobName = jobName
	cloudFoundryJobs.ResourcesPool = resourcePool
	cloudFoundryJobs.Instances = instacnes
	cloudFoundryJobs.StaticIp = staticIp
	if staticIp != nil {
		cloudFoundryJobs.StaticIps = strings.Join(staticIp, ",")
	}
	return
}

func NewFloatingNetWork(floatingIp string) (netWorks NetWorks) {
	netWorks.StaticIp = floatingIp
	netWorks.Name = "floating"
	netWorks.NetType = "vip"
	return
}

func NewRelease(name, version string) (release Release) {
	if name == "" {
		name = "cf"
	}
	if version == "" {
		version = "170"
	}
	release.Name = name
	release.Version = version
	return
}

func NewStemcell(name, version string) (stemcell Stemcells) {
	if name == "" {
		name = "bosh-openstack-kvm-ubuntu-lucid-go_agent"
	}
	if version == "" {
		version = "2719"
	}
	stemcell.Name = name
	stemcell.Version = version
	return
}

func NewSimpleCloudFoundryProperties(name string) (cloudFoundryProperties CloudFoundryProperties) {
	cloudFoundryProperties.Name = name
	return
}

func (cloudFoundryProperties *CloudFoundryProperties) Load() error {

	cloud := cloudFoundryProperties.CloudProperties
	cloud.Load()

	queryOneErr := orm.NewOrm().QueryTable(cloudFoundryProperties).One(cloudFoundryProperties, "Id")
	if queryOneErr != nil {
		logger.Error("Query One failed %s", queryOneErr)
	}
	errors := orm.NewOrm().Read(cloudFoundryProperties, "Id")
	if errors != nil {
		logger.Error("Read CloudFoundryProperties error : %s", errors)
		_, err := orm.NewOrm().Insert(cloudFoundryProperties)
		if err != nil {
			logger.Error("Inert CloudFoundryProperties error %s ", err)
		}
	}

	cloudFoundryProperties.CloudProperties = cloud

	return errors
}

func (cloudFoundryProperties *CloudFoundryProperties) Update() error {
	_, err := orm.NewOrm().Update(cloudFoundryProperties)
	if err != nil {
		logger.Error("Update CloudFoundryProperties error %s ", err)
	}
	return err
}

func (compilation *Compilation) Load() error {

	queryOneErr := orm.NewOrm().QueryTable(compilation).One(compilation, "Id")
	if queryOneErr != nil {
		logger.Error("Query One failed %s", queryOneErr)
	}
	errors := orm.NewOrm().Read(compilation, "Id")
	if errors != nil {
		logger.Error("Read Compilation error : %s", errors)
		_, err := orm.NewOrm().Insert(compilation)
		if err != nil {
			logger.Error("Inert Compilation error %s ", err)
		}
	}
	return errors
}

func (compilation *Compilation) Update() error {
	_, err := orm.NewOrm().Update(compilation)
	if err != nil {
		logger.Error("Update Compilation error %s ", err)
	}
	return err
}

func (netWorks *NetWorks) Load() error {

	errors := orm.NewOrm().Read(netWorks, "NetWorkName")
	if errors != nil {
		logger.Error("Read NetWorks error : %s", errors)
		_, err := orm.NewOrm().Insert(netWorks)
		if err != nil {
			logger.Error("Inert NetWorks error %s ", err)
		}
	}
	return errors
}

func (netWorks *NetWorks) Update() error {
	_, err := orm.NewOrm().Update(netWorks)
	if err != nil {
		logger.Error("Update NetWorks error %s ", err)
	}
	return err
}

func (resourcesPools *ResourcesPools) Load() error {

	errors := orm.NewOrm().Read(resourcesPools, "Id")
	if errors != nil {
		logger.Error("Read ResourcesPools error : %s", errors)
		_, err := orm.NewOrm().Insert(resourcesPools)
		if err != nil {
			logger.Error("Inert ResourcesPools error %s ", err)
		}
	}
	return errors
}

func (resourcesPools *ResourcesPools) Update() error {
	_, err := orm.NewOrm().Update(resourcesPools)
	if err != nil {
		logger.Error("Update ResourcesPools error %s ", err)
	}
	return err
}

func LoadResourcePools(resourcesPools []ResourcesPools) ([]ResourcesPools, error) {
	qs := orm.NewOrm().QueryTable(new(ResourcesPools))

	cnt, _ := qs.Count()
	if cnt == 0 {
		for index, value := range resourcesPools {
			value.Load()
			resourcesPools[index] = value
		}
		return resourcesPools, nil
	}

	var res []ResourcesPools
	_, err := qs.All(&res)
	return res, err
}

func UpdateResourcePools(resourcesPools []ResourcesPools) ([]ResourcesPools, error) {
	qs := orm.NewOrm().QueryTable(new(ResourcesPools))
	cond := orm.NewCondition()

	cond1 := cond.And("id__isnull", false)
	qs.SetCond(cond1).Delete()

	insert, _ := qs.PrepareInsert()

	for index, values := range resourcesPools {
		id, err := insert.Insert(&values)

		if err != nil {
			logger.Error("Insert ResourcePools error :%s", err)
		} else {
			values.Id = id
			resourcesPools[index] = values
		}

	}

	insert.Close()

	return resourcesPools, nil
}

func (cloudFoundryJobs *CloudFoundryJobs) Load() error {
	logger.Debug(" before load jobs %s", cloudFoundryJobs)
	errors := orm.NewOrm().Read(cloudFoundryJobs, "JobName")
	if errors == nil {
		cloudFoundryJobs.StaticIp = strings.Split(cloudFoundryJobs.StaticIps, ",")
	}
	if errors != nil {
		logger.Error("Read CloudFoundryJobs error : %s", errors)
		cloudFoundryJobs.StaticIps = strings.Join(cloudFoundryJobs.StaticIp, ",")
		_, err := orm.NewOrm().Insert(cloudFoundryJobs)
		if err != nil {
			logger.Error("Inert CloudFoundryJobs error %s ", err)
		}
	}
	return errors
}

func (cloudFoundryJobs *CloudFoundryJobs) Update() error {
	logger.Debug(" before update jobs %s", cloudFoundryJobs)
	cloudFoundryJobs.StaticIps = strings.Join(cloudFoundryJobs.StaticIp, ",")
	_, err := orm.NewOrm().Update(cloudFoundryJobs)
	if err != nil {
		logger.Error("Update CloudFoundryJobs error %s ", err)
	}
	return err
}

func init() {
	orm.RegisterModel(new(CloudFoundryProperties), new(Compilation), new(NetWorks), new(CloudFoundryJobs), new(ResourcesPools))
}
