package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
)

type MicroBoshController struct {
	beego.Controller
}

func (this *MicroBoshController) Get() {
	action := this.GetString("action")
	this.Layout = "index.tpl"
	this.Data["NavBarFocus"] = "microbosh"
	if action == "config" {
		this.ConfigMicroBOSH()
	} else if action == "deploy" {
		this.DeployMicroBOSH()
	} else {
		this.IndexMicroBOSH()
	}
}

func (this *MicroBoshController) Post() {

	network := entity.NewNetWork(this.GetString("vip"),
		this.GetString("net_id"))
	resources := entity.NewResources(this.GetString("persistent_disk"),
		this.GetString("instance_type"),
		this.GetString("availability_zone"))
	cloudproperties := entity.NewCloudProperties(this.GetString("auth_url"),
		this.GetString("username"),
		this.GetString("api_key"),
		this.GetString("tenant"),
		this.GetString("default_key_name"),
		"/home/ubuntu/bosh-workspace/key/test_microbosh.private",
		this.GetString("cci_ebs_url"),
		this.GetString("cci_ebs_accesskey"),
		this.GetString("cci_ebs_secretkey"))
	microbosh := entity.NewMicroBOSH(this.GetString("name"),
		network,
		resources,
		cloudproperties)

	this.CommitData(microbosh)
	this.Get()
}

func (this *MicroBoshController) IndexMicroBOSH() {
	this.LoadData()
	this.TplNames = "microbosh/index.tpl"
}

func (this *MicroBoshController) ConfigMicroBOSH() {
	logger.Debug("%s", "Config MicroBOSH")
	this.LoadData()
	this.TplNames = "microbosh/config.tpl"
}

func (this *MicroBoshController) DeployMicroBOSH() {
	logger.Debug("%s", "Deploy MicroBOSH")
	this.LoadData()
	this.Deploy()
	this.Data["HOST"] = this.Ctx.Request.Host
	this.Data["AppName"] = globaleAppName
	this.Data["WebSocket"] = microWebsocket
	this.TplNames = "microbosh/deploy.tpl"
}

//read data from const or database
func (this *MicroBoshController) LoadData() {
	this.Data["MicroBOSH"] = mi
}

//deploy
func (this *MicroBoshController) Deploy() {
	microBOSHTemplate := entity.NewMicroBOSHTemplate(mi)
	ok, err := microBOSHTemplate.CreateMicroBOSHYaml(microPath)
	if !ok {
		this.Data["Message"] = fmt.Sprintf("Error: %s", err)
	} else {
		this.Data["Message"] = fmt.Sprintf("Successful make deployment file: %s", microPath)
	}
}

//write data to const or database
func (this *MicroBoshController) CommitData(microbosh entity.MicroBOSH) {
	mi = microbosh
}

var mi entity.MicroBOSH = entity.NewMicroBOSH("microbosh-openstack",
	entity.NewNetWork("192.168.133.108", "8bb21e6e-dc6a-409c-82d0-a110fb3c9fe1"),
	entity.NewResources("16384", "flavor_71", "zone2"),
	entity.NewCloudProperties("http://192.168.128.2:5000/v2.0", "paas", "paas123", "Project_paas", "test_microbosh", "/home/ubuntu/bosh-workspace/key/test_microbosh.private", "http://192.168.128.5:8080/EBS", "7fd292943f3a46c491fbe4152cf3c8f0", "8e100b9430ab429397352359b16f01b0"))
