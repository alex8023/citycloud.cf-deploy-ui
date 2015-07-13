package controllers

import (
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
		this.GetString("private_key"),
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

func (this *MicroBoshController) IndexMicroBOSH(){
	this.LoadData()
	this.TplNames = "microbosh/index.tpl"
}

func (this *MicroBoshController) ConfigMicroBOSH(){
	logger.Debug("%s","Config MicroBOSH")
	this.LoadData()
	this.TplNames = "microbosh/config.tpl"
}

func (this *MicroBoshController) DeployMicroBOSH(){
	logger.Debug("%s","Deploy MicroBOSH")
	this.LoadData()
	this.TplNames = "microbosh/config.tpl"
}

//read data from const or database
func (this *MicroBoshController) LoadData(){
	this.Data["MicroBOSH"] = mi
}

//write data to const or database
func (this *MicroBoshController) CommitData(microbosh entity.MicroBOSH){
	mi = microbosh
}
	
var mi entity.MicroBOSH = entity.NewMicroBOSH("deployment-microbosh",
	entity.NewNetWork("vip","netid"),
	entity.NewResources("16384","flavor_100","zone2"),
	entity.NewCloudProperties("auth_url","username","apikey","tenant","defaultkeyname","privatekey","ebsurl","ebskey","ebssercetkey"))
