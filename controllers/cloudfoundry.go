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

type CloudFoundryController struct {
	beego.Controller
}

func (this *CloudFoundryController) Get() {
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

func (this *CloudFoundryController) Post() {
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
		iaas := this.GetString("iaasVsersion")
		// use default MicroBOSH vip
		privateNetWorks = entity.NewNetWorks(this.GetString("private-name"),
			this.GetString("private-netWorkName"),
			this.GetString("private-netType"),
			this.GetString("private-netId"),
			this.GetString("private-cidr"),
			this.GetString("private-dns"),
			mi.NetWork.Vip,
			this.GetString("private-reservedIp"),
			this.GetString("private-staticIp"))
		//rebuild
		netWorksMap = make(map[string]entity.NetWorks)
		netWorksMap[utils.Net_Private] = privateNetWorks

		if iaas == defaultVersion {
			publicNetWorks = entity.NewNetWorks(this.GetString("public-name"),
				this.GetString("public-netWorkName"),
				this.GetString("public-netType"),
				this.GetString("public-netId"),
				this.GetString("public-cidr"),
				"",
				"",
				"",
				this.GetString("public-staticIp"))
		} else {
			publicNetWorks = entity.NewFloatingNetWork(cf.CloudFoundryProperties.FloatingIp)
		}
		netWorksMap[utils.Net_Public] = publicNetWorks

		compilation = cf.Compilation
		compilation.DefaultNetWork = privateNetWorks.Name
		this.CommitData(compilation)

		for key, _ := range resourcesPoolsMap {
			resourcesPoolsMap[key].DefaultNetWork = privateNetWorks.Name
		}

		this.CommitData(resourcesPoolsMap)
		this.CommitData(netWorksMap)
	case utils.Compilation:
		workers, _ := this.GetInt("workers", 6)
		if workers <= 0 {
			workers = 6
		}
		compilation = entity.NewCompilation(this.GetString("instanceType"),
			this.GetString("availabilityZone"),
			workers,
			this.GetString("defaultNetWork"))

		id, err := this.GetInt64("id", 0)
		if err == nil {
			compilation.Id = id
		}
		this.CommitData(compilation)
	case utils.ResourcesPools:
		arrName := this.GetStrings("name")
		arrInstanceType := this.GetStrings("instanceType")
		arrAvailabilityZone := this.GetStrings("availabilityZone")
		arrDefaultNetWork := this.GetStrings("defaultNetWork")
		arrSize := this.GetStrings("size")

		resourcesPoolsMap = make([]entity.ResourcesPools, 0)

		for index, value := range arrName {
			size, _ := strconv.Atoi(arrSize[index])

			if size <= 0 {
				size = 1
			}

			addResourcesPool := entity.NewResourcesPools(value,
				arrInstanceType[index],
				arrAvailabilityZone[index],
				arrDefaultNetWork[index],
				size,
			)

			if index == 0 {
				resourcesPool = addResourcesPool
			}
			resourcesPoolsMap = append(resourcesPoolsMap, addResourcesPool)
		}
		this.CommitData(resourcesPoolsMap)
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

		for index, value := range resourcesPoolsMap {
			poolsize := 0
			for _, v := range cloudFoundryJobsMap {
				if value.Name == v.ResourcesPool {
					poolsize = poolsize + 1
				}
			}
			value.Size = poolsize
			resourcesPoolsMap[index] = value
		}
		this.CommitData(resourcesPoolsMap)

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

func (this *CloudFoundryController) DeployCloudFoundry() {
	logger.Debug("%s", "Deploy CloudFoundry")
	this.LoadData()
	this.Deploy()
	this.Data["HOST"] = this.Ctx.Request.Host
	this.Data["AppName"] = globaleAppName
	this.Data["WebSocket"] = cloudfoundryWebsocket
	this.TplNames = "cloudfoundry/deploy.tpl"
}

func (this *CloudFoundryController) IndexCloudFoundry() {
	this.LoadData()
	this.TplNames = "cloudfoundry/index.tpl"
}

func (this *CloudFoundryController) ConfigCloudFoundry() {
	logger.Debug("%s", "Config CloudFoundry")
	this.LoadData()
	model := this.GetString("model")
	if model == "" {
		model = utils.CloudFoundryProperties
	}
	this.Data["Model"] = model
	switch model {
	case utils.CloudFoundryProperties:
		this.TplNames = "cloudfoundry/config_properties.tpl"
	case utils.NetWorks:
		this.TplNames = "cloudfoundry/config_networks.tpl"
	case utils.Compilation:
		this.TplNames = "cloudfoundry/config_compilation.tpl"
	case utils.ResourcesPools:
		this.Data["Pools"] = len(cf.ResourcesPools)
		this.TplNames = "cloudfoundry/config_resourcespools.tpl"
	case utils.CloudFoundryJobs:
		this.TplNames = "cloudfoundry/config_jobs.tpl"
	default:
		this.IndexCloudFoundry()
	}
}

func (this *CloudFoundryController) ConfigProperties() {
	logger.Debug("%s", "Config CloudFoundry Properties")
	model := this.GetString("model")
	if model == "" {
		model = utils.Properties
	}
	this.Data["Model"] = model
	properties.Load()
	cf.Properties = properties
	this.Data["Properties"] = properties
	this.TplNames = "cloudfoundry/config_more.tpl"
}

//read data from const or database
func (this *CloudFoundryController) LoadData() {
	compilation.Load()

	cloudFoundryProperties.Load()

	for key, network := range netWorksMap {
		network.Load()
		netWorksMap[key] = network
	}

	for key, jobs := range cloudFoundryJobsMap {
		jobs.Load()
		cloudFoundryJobsMap[key] = jobs
	}

	resourcesPoolsMap, _ = entity.LoadResourcePools(resourcesPoolsMap)

	cf.Compilation = compilation

	cf.CloudFoundryProperties = cloudFoundryProperties

	cf.NetWorks = netWorksMap

	cf.CloudFoundryJobs = cloudFoundryJobsMap

	cf.ResourcesPools = resourcesPoolsMap

	properties.Load()

	cf.Properties = properties

	this.Data["CloudFoundry"] = cf
}

//deploy,only set deployment
func (this *CloudFoundryController) Deploy() {
	// reload data
	this.LoadData()

	cf.Release = entity.NewRelease(cloudFoundryName, cloudFoundryVersion)
	cf.Stemcells = entity.NewStemcell(stemcellName, stemcellVersion)

	templateText := ""
	if iaasVersion == defaultVersion {
		templateText = entity.CloudFoundryTemplateV3
		ipArr := utils.SpliteIPs(publicNetWorks.StaticIp)
		cf.CloudFoundryProperties.FloatingIp = ipArr[0]
		//设置ip段的第一个ip为haproxy的浮动ip
	} else {
		templateText = entity.CloudFoundryTemplateV2
	}

	cloudFoundryTemplate := entity.NewCloudFoundryTemplate(cf)

	ok, err := cloudFoundryTemplate.CreateCloudFoundryYaml(templateText, cloudFoundryPath)
	if !ok {
		this.Data["Message"] = fmt.Sprintf("Error: %s", err)
	} else {
		this.Data["Message"] = fmt.Sprintf("Successful make deployment file: %s", cloudFoundryPath)
	}
}

//write data to const or database
func (this *CloudFoundryController) CommitData(data interface{}) {
	switch data.(type) {
	case map[string]entity.CloudFoundryJobs:
		cloudFoundryJobsMap = data.(map[string]entity.CloudFoundryJobs)
		for key, value := range cloudFoundryJobsMap {
			value.Update()
			cloudFoundryJobsMap[key] = value
		}
		cf.CloudFoundryJobs = cloudFoundryJobsMap

	case entity.CloudFoundryProperties:
		cloudFoundryProperties = data.(entity.CloudFoundryProperties)
		cloudFoundryProperties.Update()
		cf.CloudFoundryProperties = cloudFoundryProperties

	case entity.Compilation:
		compilation = data.(entity.Compilation)
		compilation.Update()
		cf.Compilation = compilation

	case entity.CloudFoundry:
		cf = data.(entity.CloudFoundry)

	case map[string]entity.NetWorks:
		netWorksMap = data.(map[string]entity.NetWorks)
		for key, value := range netWorksMap {
			value.Update()
			netWorksMap[key] = value
		}
		cf.NetWorks = netWorksMap

	case []entity.ResourcesPools:
		resourcesPoolsMap = data.([]entity.ResourcesPools)
		resourcesPoolsMap, _ = entity.UpdateResourcePools(resourcesPoolsMap)
		cf.ResourcesPools = resourcesPoolsMap

	case entity.Properties:
		properties = data.(entity.Properties)
		properties.Update()
		cf.Properties = properties

	default:
		this.Data["MessageErr"] = fmt.Sprintf("Unknow type %s", reflect.TypeOf(data))
	}

	// //or
	//	switch result := data.(type) {
	//	case map[string]entity.CloudFoundryJobs:

	//		for key, value := range result {
	//			value.Update()
	//			result[key] = value
	//		}
	//		cf.CloudFoundryJobs = result
	//	case entity.CloudFoundryProperties:
	//		result.Update()
	//		cf.CloudFoundryProperties = result
	//	case entity.Compilation:
	//		result.Update()
	//		cf.Compilation = result
	//	case entity.CloudFoundry:
	//		cf = result

	//	case map[string]entity.NetWorks:
	//		for key, value := range result {
	//			value.Update()
	//			result[key] = value
	//		}
	//		cf.NetWorks = result

	//	case []entity.ResourcesPools:
	//		result, _ = entity.UpdateResourcePools(result)
	//		cf.ResourcesPools = result
	//	default:
	//		this.Data["MessageErr"] = fmt.Sprintf("Unknow type %s", reflect.TypeOf(data))
	//	}

}

var (
	cloudFoundryProperties = entity.NewCloudFoundryProperties("cf-release",
		"57cfc863-786d-4495-bb97-86d2f650a038", "192.168.133.102", "ccipaas.net", "cci")

	compilation = entity.NewCompilation("flavor_91", "zone2", 6, privateNetWorks.Name)

	privateNetWorks = entity.NewNetWorks("cf1",
		utils.Net_Private,
		"manual",
		"8bb21e6e-dc6a-409c-82d0-a110fb3c9fe1",
		"192.168.129.0/24",
		"10.10.170.2",
		"192.168.133.108",
		"192.168.129.1 - 192.168.129.99",
		"192.168.129.100 - 192.168.129.126")

	publicNetWorks = entity.NewNetWorks("publc",
		utils.Net_Public,
		"manual",
		"8bb21e6e-dc6a-409c-82d0-a110fb3c9fe1",
		"10.162.2.0/24",
		"",
		"",
		"",
		"10.162.2.28 - 10.162.2.29")

	netWorksMap map[string]entity.NetWorks = map[string]entity.NetWorks{
		utils.Net_Private: privateNetWorks, utils.Net_Public: publicNetWorks,
	}

	resourcesPool = entity.NewResourcesPools(
		"normal",
		"flavor_91",
		"zone2",
		"cf1",
		4)
	resourcesPool2 = entity.NewResourcesPools(
		"normal2",
		"flavor_91",
		"zone2",
		"cf1",
		4)

	resourcesPoolsMap []entity.ResourcesPools = []entity.ResourcesPools{
		resourcesPool, resourcesPool2,
	}

	properties = entity.Properties{}

	cf entity.CloudFoundry = entity.NewCloudFoundry(
		cloudFoundryProperties,
		compilation,
		netWorksMap,
		resourcesPoolsMap,
		cloudFoundryJobsMap,
		properties)

	cloudFoundryJobsMap = map[string]entity.CloudFoundryJobs{
		utils.Job_Haproxy: entity.NewCloudFoundryJobs(
			"haproxy",
			utils.Job_Haproxy,
			"normal2",
			1,
			[]string{""}),
		utils.Job_Gorouter: entity.NewCloudFoundryJobs(
			"gorouter",
			utils.Job_Gorouter,
			"normal",
			1,
			[]string{""}),
		utils.Job_Postgres: entity.NewCloudFoundryJobs(
			"postgres",
			utils.Job_Postgres,
			"normal",
			1,
			[]string{""}),
		utils.Job_NFS: entity.NewCloudFoundryJobs(
			"nfs",
			utils.Job_NFS,
			"normal",
			1,
			[]string{""}),
		utils.Job_NATS: entity.NewCloudFoundryJobs(
			"nats",
			utils.Job_NATS,
			"normal",
			1,
			[]string{""}),
		utils.Job_Syslog: entity.NewCloudFoundryJobs(
			"syslog_aggregator",
			utils.Job_Syslog,
			"normal",
			1,
			[]string{""}),
		utils.Job_Etcd: entity.NewCloudFoundryJobs(
			"etcd",
			utils.Job_Etcd,
			"normal",
			1,
			[]string{""}),
		utils.Job_Loggregator: entity.NewCloudFoundryJobs(
			"loggregator",
			utils.Job_Loggregator,
			"normal",
			1,
			[]string{""}),
		utils.Job_Loggregator_Traffic: entity.NewCloudFoundryJobs(
			"loggregator_traffic",
			utils.Job_Loggregator_Traffic,
			"normal",
			1,
			[]string{""}),
		utils.Job_Uaa: entity.NewCloudFoundryJobs(
			"uaa", utils.Job_Uaa,
			"normal",
			1,
			[]string{""}),
		utils.Job_Cloud_Controller_NG: entity.NewCloudFoundryJobs(
			"cloud_controller_ng",
			utils.Job_Cloud_Controller_NG,
			"normal",
			1,
			[]string{""}),
		utils.Job_Cloud_Controller_Worker: entity.NewCloudFoundryJobs(
			"cloud_controller_worker",
			utils.Job_Cloud_Controller_Worker,
			"normal",
			1,
			[]string{""}),
		utils.Job_Cloud_Controller_Clock: entity.NewCloudFoundryJobs(
			"cloud_controller_clock",
			utils.Job_Cloud_Controller_Clock,
			"normal",
			1,
			[]string{""}),
		utils.Job_Hm9000: entity.NewCloudFoundryJobs(
			"hm9000",
			utils.Job_Hm9000,
			"normal",
			1,
			[]string{""}),
		utils.Job_Stats: entity.NewCloudFoundryJobs(
			"stats",
			utils.Job_Stats,
			"normal",
			1,
			[]string{""}),
		utils.Job_Dea_Next: entity.NewCloudFoundryJobs(
			"dea_next",
			utils.Job_Dea_Next,
			"normal",
			1,
			[]string{""}),
	}
)
