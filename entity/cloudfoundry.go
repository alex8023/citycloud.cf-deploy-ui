package entity

import (
	"github.com/astaxie/beego/orm"
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
	Id              int
	Name            string
	Uuid            string
	FloatingIp      string
	SystemDomain    string
	SystemDomainOrg string
	CloudProperties CloudProperties `orm:"-"`
}

// 编译机器配置
type Compilation struct {
	Id               int
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
	Id         int
	Name       string
	NetType    string
	NetId      string
	Cidr       string
	Dns        string
	PowerDns   string
	ReservedIp string
	StaticIp   string
}

//资源池配置
//此对象仅配置一个资源池对象
//在CloudFoundry中使用map组装资源池
type ResourcesPools struct {
	Id               int
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
	netType,
	netId,
	cidr,
	dns,
	powerDns,
	reservedIp,
	staticIp string) (netWorks NetWorks) {
	netWorks.Name = name
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

func init() {
	orm.RegisterModel(new(CloudFoundryProperties), new(Compilation), new(NetWorks), new(CloudFoundryJobs), new(ResourcesPools))
}
