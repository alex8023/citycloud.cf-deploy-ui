package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
	"reflect"
	"strconv"
)

type VsphereCloudFoundryController struct {
	beego.Controller
}

func (this *VsphereCloudFoundryController) Get() {
	action := this.GetString("action")
	this.Layout = "index.tpl"
	this.Data["NavBarFocus"] = "cloudfoundry"
	this.Data["IaaSVersion"] = iaasVersion
	this.Data["DefaultVersion"] = defaultVersion

	switch action {
	case "config":
		this.ConfigCloudFoundry()
	case "deploy":
		this.DeployCloudFoundry()
	case "moreconf":
		this.ConfigProperties()
	default:
		this.IndexCloudFoundry()
	}
}

func (this *VsphereCloudFoundryController) Post() {
	model := this.GetString("model")
	erros := make([]error, 0)
	switch model {
	case utils.CloudFoundryProperties:
		cloudFoundryProperties = entity.NewCloudFoundryProperties(
			this.GetString("name"),
			this.GetString("uuid"),
			this.GetString("floatingIp"),
			this.GetString("systemDomain"),
			this.GetString("systemDomainOrg"))

		id, err := this.GetInt64("id", 0)
		if err == nil {
			cloudFoundryProperties.Id = id
		}
		this.CommitData(cloudFoundryProperties)
	case utils.NetWorks:
		// use default MicroBOSH vip
		privateNetWorks = entity.NewNetWorks(this.GetString("private-name"),
			this.GetString("private-netWorkName"),
			this.GetString("private-netType"),
			this.GetString("private-netId"),
			this.GetString("private-cidr"),
			this.GetString("private-dns"),
			vsphereMicro.VsphereNetWork.Ip,
			this.GetString("private-reservedIp"),
			this.GetString("private-staticIp"))
		//rebuild
		netWorksMap = make(map[string]entity.NetWorks)
		netWorksMap[utils.Net_Private] = privateNetWorks

		vsphereCompilation = vcf.VsphereCompilation
		vsphereCompilation.DefaultNetWork = privateNetWorks.Name
		this.CommitData(vsphereCompilation)

		for key, _ := range vsphereResourcePoolMap {
			vsphereResourcePoolMap[key].DefaultNetWork = privateNetWorks.Name
		}

		this.CommitData(vsphereResourcePoolMap)
		this.CommitData(netWorksMap)
	case utils.Compilation:
		workers, _ := this.GetInt("workers", 6)
		if workers <= 0 {
			workers = 6
		}
		//i am here
		vid, _ := this.GetInt64("vid", 0)

		vsphereCompilation = entity.NewVsphereCompilation(workers, this.GetString("defaultNetWork"), vid)

		id, err := this.GetInt64("id", 0)
		if err == nil {
			vsphereCompilation.Id = id
		}
		this.CommitData(vsphereCompilation)
	case utils.ResourcesPools:
		arrName := this.GetStrings("name")
		arrDefaultNetWork := this.GetStrings("defaultNetWork")
		arrSize := this.GetStrings("size")
		arrId := this.GetStrings("id")
		arrVid := this.GetStrings("vid")

		vsphereResourcePoolMap = make([]entity.VsphereResourcesPool, 0)

		for index, value := range arrName {
			size, _ := strconv.Atoi(arrSize[index])

			if size <= 0 {
				size = 1
			}

			vid, _ := strconv.ParseInt(arrVid[index], 10, 64)

			ids := arrId[index]

			if ids == "0" || ids == "" {
				ids = "0"
			}

			id, _ := strconv.ParseInt(ids, 10, 64)

			addResourcesPool := entity.NewVsphereResourcesPool(value,
				size,
				arrDefaultNetWork[index],
				vid,
			)
			if id != 0 {
				addResourcesPool.Id = id
			}

			if index == 0 {
				vsphereResourcePool = addResourcesPool
			}
			vsphereResourcePoolMap = append(vsphereResourcePoolMap, addResourcesPool)
		}
		fmt.Printf("============================%s", vsphereResourcePoolMap)
		this.CommitData(vsphereResourcePoolMap)
	case utils.CloudFoundryJobs:
		iPFactory := utils.NewIPFactory()
		ipArr := utils.SpliteIPs(privateNetWorks.StaticIp)

		success, iperr := iPFactory.InitFactory(ipArr[0], ipArr[1])

		if iperr != nil {
			erros = append(erros, iperr)
		}
		for key, value := range cloudFoundryJobsMap {
			value.Name = this.GetString(key + "_name")
			instance, _ := this.GetInt(key+"_instances", 1)
			if instance <= 0 {
				instance = 1
			}
			value.ResourcesPool = this.GetString(key + "_resourcesPool_select")
			if success {
				assignaIP, assignaerr := iPFactory.AssignaIP2Job(key, instance)
				if assignaerr == nil {
					value.StaticIp = assignaIP
					value.Instances = instance
				} else {
					erros = append(erros, fmt.Errorf("Assign IP for Job : %s , Errors: %s", key, assignaerr))
				}
			}
			cloudFoundryJobsMap[key] = value
		}
		this.CommitData(cloudFoundryJobsMap)

		for index, value := range vsphereResourcePoolMap {
			poolsize := 0
			for _, v := range cloudFoundryJobsMap {
				if value.Name == v.ResourcesPool {
					poolsize = poolsize + 1*v.Instances //fixed instance sum
				}
			}
			value.Size = poolsize
			vsphereResourcePoolMap[index] = value
		}
		this.CommitData(vsphereResourcePoolMap)

	case utils.Properties:
		property := properties.JobProperties
		for key, val := range property {
			val.Name = key
			val.Value = this.GetString(key)
			property[key] = val
		}
		properties.JobProperties = property
		this.CommitData(properties)

	default:
		erros = append(erros, fmt.Errorf("Unknow model %s", model))
	}

	if len(erros) != 0 {
		this.Data["MessageErr"] = fmt.Sprintf("Errors: %s", erros)
	}
	this.Get()
}

func (this *VsphereCloudFoundryController) DeployCloudFoundry() {
	logger.Debug("%s", "Deploy CloudFoundry")
	this.LoadData()
	this.Deploy()
	this.Data["HOST"] = this.Ctx.Request.Host
	this.Data["AppName"] = globaleAppName
	this.Data["WebSocket"] = cloudfoundryWebsocket
	this.TplNames = "cloudfoundry/deploy.tpl"
}

func (this *VsphereCloudFoundryController) IndexCloudFoundry() {
	this.LoadData()
	this.TplNames = "cloudfoundry/vsphere/index.tpl"
}

func (this *VsphereCloudFoundryController) ConfigCloudFoundry() {
	logger.Debug("%s", "Config CloudFoundry")
	this.LoadData()
	model := this.GetString("model")
	if model == "" {
		model = utils.CloudFoundryProperties
	}
	this.Data["Model"] = model
	switch model {
	case utils.CloudFoundryProperties:
		this.TplNames = "cloudfoundry/vsphere/config_properties.tpl"
	case utils.NetWorks:
		this.TplNames = "cloudfoundry/vsphere/config_networks.tpl"
	case utils.Compilation:
		this.TplNames = "cloudfoundry/vsphere/config_compilation.tpl"
	case utils.ResourcesPools:
		this.Data["Pools"] = len(cf.ResourcesPools)
		this.TplNames = "cloudfoundry/vsphere/config_resourcespools.tpl"
	case utils.CloudFoundryJobs:
		this.TplNames = "cloudfoundry/vsphere/config_jobs.tpl"
	default:
		this.IndexCloudFoundry()
	}
}

func (this *VsphereCloudFoundryController) ConfigProperties() {
	logger.Debug("%s", "Config CloudFoundry Properties")
	model := this.GetString("model")
	if model == "" {
		model = utils.Properties
	}
	this.Data["Model"] = model
	properties.Load()
	cf.Properties = properties
	this.Data["Properties"] = properties
	this.TplNames = "cloudfoundry/vsphere/config_more.tpl"
}

//read data from const or database
func (this *VsphereCloudFoundryController) LoadData() {
	vsphereCompilation.Load()

	vsphereCompilationVid := vsphereCompilation.Vid

	vsphereResource.Id = vsphereCompilationVid

	err := vsphereResource.Load()

	if err != nil {
		this.Data["MessageErr"] = "UnKnow instance flavor,please set instance flavor for comiplation componement"
	}

	vsphereCompilation.VsphereResource = vsphereResource

	vcf.VsphereCompilation = vsphereCompilation

	vsphereResourcePoolMap, _ = entity.LoadVsphereResourcePool(vsphereResourcePoolMap)

	vcf.VsphereResourcesPools = vsphereResourcePoolMap

	for key, jobs := range cloudFoundryJobsMap {
		jobs.Load()
		cloudFoundryJobsMap[key] = jobs
	}

	vcf.CloudFoundryJobs = cloudFoundryJobsMap

	for key, network := range netWorksMap {
		network.Load()
		netWorksMap[key] = network
	}

	vcf.NetWorks = netWorksMap

	vcf.Stemcells = stemcells

	vcf.Release = releases

	properties.Load()

	vcf.Properties = properties

	cloudFoundryProperties.Load()

	vcf.CloudFoundryProperties = cloudFoundryProperties

	this.Data["CloudFoundry"] = vcf

	vsphereResources, _ = entity.LoadVsphereResource(vsphereResources)

	this.Data["VsphereResource"] = vsphereResources
}

//deploy,only set deployment
func (this *VsphereCloudFoundryController) Deploy() {
	// reload data
	this.LoadData()

	vcf.Release = entity.NewRelease(cloudFoundryName, cloudFoundryVersion)
	vcf.Stemcells = entity.NewStemcell(stemcellName, stemcellVersion)

	cloudFoundryTemplate := entity.NewVsphereCloudFoundryTemplate(vcf)

	ok, err := cloudFoundryTemplate.CreateDefaultCloudFoundryYamlFileWithAppPath(iaasVersion, cloudFoundryPath, beego.AppPath)
	if !ok {
		this.Data["Message"] = fmt.Sprintf("Error: %s", err)
	} else {
		this.Data["Message"] = fmt.Sprintf("Successful make deployment file: %s", cloudFoundryPath)
	}
}

//write data to const or database
func (this *VsphereCloudFoundryController) CommitData(data interface{}) {
	switch data.(type) {
	case map[string]entity.CloudFoundryJobs:
		cloudFoundryJobsMap = data.(map[string]entity.CloudFoundryJobs)
		for key, value := range cloudFoundryJobsMap {
			value.Update()
			cloudFoundryJobsMap[key] = value
		}
		vcf.CloudFoundryJobs = cloudFoundryJobsMap

	case entity.CloudFoundryProperties:
		cloudFoundryProperties = data.(entity.CloudFoundryProperties)
		cloudFoundryProperties.Update()
		vcf.CloudFoundryProperties = cloudFoundryProperties

	case entity.VsphereCompilation:
		vsphereCompilation = data.(entity.VsphereCompilation)
		vsphereCompilation.Update()
		vcf.VsphereCompilation = vsphereCompilation
	case entity.VsphereCloudFoundry:
		vcf = data.(entity.VsphereCloudFoundry)

	case map[string]entity.NetWorks:
		netWorksMap = data.(map[string]entity.NetWorks)
		for key, value := range netWorksMap {
			value.Update()
			netWorksMap[key] = value
		}
		vcf.NetWorks = netWorksMap

	case []entity.VsphereResourcesPool:
		vsphereResourcePoolMap = data.([]entity.VsphereResourcesPool)
		vsphereResourcePoolMap, _ = entity.UpdateVsphereResourcePool(vsphereResourcePoolMap)
		vcf.VsphereResourcesPools = vsphereResourcePoolMap
	case entity.Properties:
		properties = data.(entity.Properties)
		properties.Update()
		vcf.Properties = properties

	default:
		this.Data["MessageErr"] = fmt.Sprintf("Unknow type %s", reflect.TypeOf(data))
	}
}

var (
	vsphereCompilation = entity.NewVsphereCompilation(6, privateNetWorks.Name, 0)

	vsphereResource = entity.NewVsphereResource("0", "4096", "4096", "2")

	vsphereResourcePool = entity.NewVsphereResourcesPool("normal", 14, "VM Network", 0)

	vsphereResourcePoolMap []entity.VsphereResourcesPool = []entity.VsphereResourcesPool{
		vsphereResourcePool,
	}

	vsphereResources []entity.VsphereResource = []entity.VsphereResource{
		vsphereResource,
	}

	stemcells = entity.NewStemcell(stemcellName, stemcellVersion)

	releases = entity.NewRelease(cloudFoundryName, cloudFoundryVersion)

	vcf = entity.VsphereCloudFoundry{}
)
