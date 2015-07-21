package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
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
	cloudfoundry := entity.CloudFoundry{}
	this.CommitData(cloudfoundry)
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
	this.TplNames = "cloudfoundry/config.tpl"
}

//read data from const or database
func (this *CloudFoundryController) LoadData() {
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

var cf entity.CloudFoundry = entity.CloudFoundry{
	CloudFoundryProperties: cloudFoundryProperties,
	Compilation:            compilation,
	NetWorks:               networks,
	ResourcesPool:          resourcesPool,
}

var cloudFoundryProperties = entity.NewCloudFoundryProperties("cf-release",
	"57cfc863-786d-4495-bb97-86d2f650a038", "192.168.133.102", "ccipaas.net", "cci")
var compilation = entity.NewCompilation("flavor_91", "zone2", 6, "cf1")

var networks = entity.NewNetWorks("cf1",
	"manual",
	"8bb21e6e-dc6a-409c-82d0-a110fb3c9fe1",
	"192.168.129.0/24",
	"10.10.170.2",
	"192.168.129.1 - 192.168.129.99",
	"192.168.129.100 - 192.168.129.126")

var resourcesPool = entity.NewResourcesPool("normal",
	"flavor_91",
	"zone2",
	"cf1",
	4)
