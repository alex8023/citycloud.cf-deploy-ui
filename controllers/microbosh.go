package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
)

type MicroBoshController struct {
	beego.Controller
}

func (this *MicroBoshController) Get() {
	action := this.GetString("action")
	this.Layout = "index.tpl"
	this.Data["NavBarFocus"] = "microbosh"
	this.Data["MicroBOSH"] = mi
	if action == "config" {
		this.TplNames = "microbosh/config.tpl"
	} else {
		this.TplNames = "microbosh/index.tpl"
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

	PrintMicroBOSH(microbosh)
	this.Data["MicroBOSH"] = microbosh
	this.Get()
}
var mi entity.MicroBOSH = entity.NewMicroBOSH("deployment-name",
	entity.NewNetWork("vip","netid"),
	entity.NewResources("16384","flavor_100","zone2"),
	entity.NewCloudProperties("auth_url","username","apikey","tenant","defaultkeyname","privatekey","ebsurl","ebskey","ebssercetkey"))
func PrintMicroBOSH(microbosh entity.MicroBOSH) {
	fmt.Println(microbosh)
	mi = microbosh
}
