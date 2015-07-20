package entity

type CloudFoundry struct {
	Name            string
	Uuid            string
	FloatingIp      string
	SystemDomain    string
	SystemDomainOrg string
	Compilation     Compilation
	NetWorks        NetWorks
	ResourcesPool   ResourcesPool
	Properties      Properties
}

// 编译机器配置
type Compilation struct {
	InstanceType     string
	AvailabilityZone string
	Workers          int
	DefaultNetWork   string
}

//网络配置
type NetWorks struct {
	Name       string
	NetType    string
	NetId      string
	Cidr       string
	Dns        string
	ReservedIp string
	StaticIp   string
}

type ResourcesPool struct {
	Name             string
	InstanceType     string
	AvailabilityZone string
	Size             int
	DefaultNetWork   string
}

type Properties struct {
}

func NewCloudFoundry(name string,
	uuid string,
	floatingIp string,
	systemDomain string,
	systemDomainOrg string,
	compilation Compilation,
	netWorks NetWorks,
	resourcesPool ResourcesPool,
	properties Properties) (cloudfoundry CloudFoundry) {
	cloudfoundry.Name = name
	cloudfoundry.FloatingIp = floatingIp
	cloudfoundry.SystemDomain = systemDomain
	cloudfoundry.SystemDomainOrg = systemDomainOrg
	cloudfoundry.Compilation = compilation
	cloudfoundry.NetWorks = netWorks
	cloudfoundry.ResourcesPool = resourcesPool
	cloudfoundry.Properties = properties
	return
}

func NewCompilation(instanceType,
	availabilityZone string,
	workers int,
	defaultNetWork string) (compilation Compilation) {
	compilation.AvailabilityZone = availabilityZone
	compilation.InstanceType = instanceType
	compilation.Workers = workers
	compilation.DefaultNetWork = defaultNetWork
	return
}

func NewNetWorks(name,
	netType,
	netId,
	cidr,
	dns,
	reservedIp,
	staticIp string) (netWorks NetWorks) {
	netWorks.Name = name
	netWorks.NetType = netType
	netWorks.NetId = netId
	netWorks.Cidr = cidr
	netWorks.Dns = dns
	netWorks.ReservedIp = reservedIp
	netWorks.StaticIp = staticIp
	return
}

func NewResourcesPool(name,
	instanceType,
	availabilityZone,
	defaultNetWork string,
	size int) (resourcesPool ResourcesPool) {
	resourcesPool.Name = name
	resourcesPool.InstanceType = instanceType
	resourcesPool.AvailabilityZone = availabilityZone
	resourcesPool.DefaultNetWork = defaultNetWork
	resourcesPool.Size = size
	return
}

func NewSimpleCloudFoundry(name string) (cloudfoundry CloudFoundry) {
	cloudfoundry.Name = name
	return
}
