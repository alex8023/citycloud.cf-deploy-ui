package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
	"strconv"
)

type CloudFoundryController struct {
	beego.Controller
}

func (this *CloudFoundryController) Get() {
	action := this.GetString("action")
	this.Layout = "index.tpl"
	this.Data["NavBarFocus"] = "cloudfoundry"
	if action == "config" {
		this.ConfigCloudFoundry()
	} else if action == "deploy" {
		this.DeployCloudFoundry()
	} else {
		this.IndexCloudFoundry()
	}
}

func (this *CloudFoundryController) Post() {
	model := this.GetString("model")
	erros := make([]error, 0)
	if model == "CloudFoundryProperties" {
		cloudFoundryProperties = entity.NewCloudFoundryProperties(
			this.GetString("name"),
			this.GetString("uuid"),
			this.GetString("floatingIp"),
			this.GetString("systemDomain"),
			this.GetString("systemDomainOrg"))
		cf.CloudFoundryProperties = cloudFoundryProperties

	} else if model == "NetWorks" {
		// use default MicroBOSH vip
		networks = entity.NewNetWorks(this.GetString("name"),
			this.GetString("netType"),
			this.GetString("netId"),
			this.GetString("cidr"),
			this.GetString("dns"),
			mi.Network.Vip,
			this.GetString("reservedIp"),
			this.GetString("staticIp"))
		//rebuild
		netWorksMap = make([]entity.NetWorks, 0)
		netWorksMap = append(netWorksMap, networks)
		cf.Compilation.DefaultNetWork = networks.Name
		for key, _ := range cf.ResourcesPools {
			cf.ResourcesPools[key].DefaultNetWork = networks.Name
		}
		cf.NetWorks = netWorksMap
	} else if model == "Compilation" {
		workers, _ := this.GetInt("workers", 6)
		if workers <= 0 {
			workers = 6
		}

		compilation = entity.NewCompilation(this.GetString("instanceType"),
			this.GetString("availabilityZone"),
			workers,
			this.GetString("defaultNetWork"))
		cf.Compilation = compilation
	} else if model == "ResourcesPools" {
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
		cf.ResourcesPools = resourcesPoolsMap
	} else if model == "CloudFoundryJobs" {
		iPFactory := utils.NewIPFactory()
		ipArr := utils.SpliteIPs(networks.StaticIp)

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
			value.Instances = instance
			value.ResourcesPool = this.GetString(key + "_resourcesPool_select")
			value.StaticIp = []string{""}
			if success {
				assignaIP, assignaerr := iPFactory.AssignaIP2Job(key, instance)
				if assignaerr == nil {
					value.StaticIp = assignaIP
				} else {
					erros = append(erros, assignaerr)
				}
			}
			cloudFoundryJobsMap[key] = value
		}
		cf.CloudFoundryJobs = cloudFoundryJobsMap
	}

	//	cloudfoundry := entity.CloudFoundry{}
	//	this.CommitData(cloudfoundry)
	fmt.Println(erros)
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
		model = "CloudFoundryProperties"
	}
	this.Data["Model"] = model
	if model == "CloudFoundryProperties" {
		this.TplNames = "cloudfoundry/config_properties.tpl"
	} else if model == "NetWorks" {
		this.TplNames = "cloudfoundry/config_networks.tpl"
	} else if model == "Compilation" {
		this.TplNames = "cloudfoundry/config_compilation.tpl"
	} else if model == "ResourcesPools" {
		this.Data["Pools"] = len(cf.ResourcesPools)
		this.TplNames = "cloudfoundry/config_resourcespools.tpl"
	} else if model == "CloudFoundryJobs" {
		this.TplNames = "cloudfoundry/config_jobs.tpl"
	}

}

//read data from const or database
func (this *CloudFoundryController) LoadData() {
	//logger.Debug("cloudfoundry properties: %s", cf)
	this.Data["CloudFoundry"] = cf
}

//deploy,only set deployment
func (this *CloudFoundryController) Deploy() {
	cloudFoundryTemplate := entity.NewCloudFoundryTemplate(cf)
	ok, err := cloudFoundryTemplate.CreateCloudFoundryV2Yaml(cloudFoundryPath)
	if !ok {
		this.Data["Message"] = fmt.Sprintf("Error: %s", err)
	} else {
		this.Data["Message"] = fmt.Sprintf("Successful make deployment file: %s", cloudFoundryPath)
	}
}

//write data to const or database
func (this *CloudFoundryController) CommitData(cloudfoundry entity.CloudFoundry) {
	cf = cloudfoundry
}

var (
	cloudFoundryProperties = entity.NewCloudFoundryProperties("cf-release",
		"57cfc863-786d-4495-bb97-86d2f650a038", "192.168.133.102", "ccipaas.net", "cci")

	compilation = entity.NewCompilation("flavor_91", "zone2", 6, networks.Name)

	networks = entity.NewNetWorks("cf1",
		"manual",
		"8bb21e6e-dc6a-409c-82d0-a110fb3c9fe1",
		"192.168.129.0/24",
		"10.10.170.2",
		"192.168.133.108",
		"192.168.129.1 - 192.168.129.99",
		"192.168.129.100 - 192.168.129.126")

	netWorksMap []entity.NetWorks = []entity.NetWorks{
		networks,
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
		utils.Job_Stats: entity.NewCloudFoundryJobs("stats",
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
