package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
)

type VsphereMicroBoshController struct {
	beego.Controller
}

func (this *VsphereMicroBoshController) Get() {
	action := this.GetString("action")
	this.Layout = "index.tpl"
	this.Data["NavBarFocus"] = "microbosh"
	this.Data["IaaSVersion"] = iaasVersion
	this.Data["DefaultVersion"] = vsphereVersion
	if action == "config" {
		this.ConfigMicroBOSH()
	} else if action == "deploy" {
		this.DeployMicroBOSH()
	} else {
		this.IndexMicroBOSH()
	}
}

func (this *VsphereMicroBoshController) Post() {

	vsphereNetWork := entity.NewVsphereNetWork(this.GetString("ip"),
		this.GetString("netMask"),
		this.GetString("gateWay"),
		this.GetString("dns"),
		this.GetString("vlanName"),
	)
	vsphereNetworkId, _ := this.GetInt64("vsphereNetworkId", 0)
	vsphereNetWork.Id = vsphereNetworkId

	vsphereResource := entity.NewVsphereResource(this.GetString("persistentDisk"),
		this.GetString("ram"),
		this.GetString("disk"),
		this.GetString("cpu"),
	)
	vsphereResourceId, _ := this.GetInt64("vsphereResourceId", 0)
	vsphereResource.Id = vsphereResourceId

	vsphereVcenter := entity.NewVsphereVcenter(this.GetString("userName"),
		this.GetString("passwd"),
		this.GetString("host"),
		this.GetString("dataCenterName"),
		this.GetString("vmFolder"),
		this.GetString("templateFolder"),
		this.GetString("diskPath"),
		this.GetString("dataStore"),
		this.GetString("clustersName"),
	)
	vsphereVcenterId, _ := this.GetInt64("vsphereVcenterId", 0)
	vsphereVcenter.Id = vsphereVcenterId

	vsphere := entity.NewVsphereMicro(this.GetString("name"), vsphereNetWork, vsphereResource, vsphereVcenter)
	vsphereMicroId, _ := this.GetInt64("vsphereMicroId", 0)
	vsphere.Id = vsphereMicroId
	vsphere.VsphereResourceId = vsphereResourceId
	this.CommitData(vsphere)
	this.Get()
}

func (this *VsphereMicroBoshController) IndexMicroBOSH() {
	this.LoadData()
	this.TplNames = "microbosh/index_vsphere.tpl"
}

func (this *VsphereMicroBoshController) ConfigMicroBOSH() {
	logger.Debug("%s", "Config VsphereMicroBOSH")
	this.LoadData()
	this.TplNames = "microbosh/config_vsphere.tpl"
}

func (this *VsphereMicroBoshController) DeployMicroBOSH() {
	logger.Debug("%s", "Deploy VsphereMicroBOSH")
	this.LoadData()
	this.Deploy()
	this.Data["HOST"] = this.Ctx.Request.Host
	this.Data["AppName"] = globaleAppName
	this.Data["WebSocket"] = microWebsocket
	this.TplNames = "microbosh/deploy.tpl"
}

//read data from const or database
func (this *VsphereMicroBoshController) LoadData() {

	vsphereMicro.Load()

	vsphereNetWork := vsphereMicro.VsphereNetWork
	vsphereNetWork.Load()

	vsphereResource := vsphereMicro.VsphereResource

	if vsphereMicro.VsphereResourceId == 0 {
		vsphereResource.Load()
		vsphereMicro.VsphereResourceId = vsphereResource.Id
		vsphereMicro.Update()
	} else {
		vsphereResource.Id = vsphereMicro.VsphereResourceId
		vsphereResource.Load()
	}

	vsphereVcenter := vsphereMicro.VsphereVcenter
	vsphereVcenter.Load()

	vsphereMicro.VsphereNetWork = vsphereNetWork
	vsphereMicro.VsphereResource = vsphereResource
	vsphereMicro.VsphereVcenter = vsphereVcenter

	this.Data["VsphereMicroBOSH"] = vsphereMicro
}

//deploy
func (this *VsphereMicroBoshController) Deploy() {

	// reload data
	this.LoadData()

	microBOSHTemplate := entity.NewVspherewMicroTemplate(vsphereMicro)

	ok, err := microBOSHTemplate.CreateMicroBOSHYamlFile(iaasVersion, microPath)
	if !ok {
		this.Data["Message"] = fmt.Sprintf("Error: %s", err)
	} else {
		this.Data["Message"] = fmt.Sprintf("Successful make deployment file: %s", microPath)
	}
}

//write data to const or database
func (this *VsphereMicroBoshController) CommitData(vsphere entity.VsphereMicro) {

	vsphere.Update()

	vsphere.VsphereNetWork.Update()

	vsphere.VsphereResource.Update()

	vsphere.VsphereVcenter.Update()

	vsphereMicro = vsphere

}

var vsphereMicro entity.VsphereMicro = entity.NewVsphereMicro("micro01",
	entity.NewVsphereNetWork("10.10.170.12",
		"255.255.255.0",
		"10.10.170.254",
		"10.10.170.2",
		"VM Network"),
	entity.NewVsphereResource("16384",
		"4096",
		"16384",
		"2"),
	entity.NewVsphereVcenter("administrator",
		"P@ssw0rd",
		"10.10.90.233",
		"vCenter",
		"vm_folder",
		"template_folder",
		"boshdeployer",
		"datastore1302",
		"vCluster"),
)
