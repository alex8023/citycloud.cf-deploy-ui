package controllers

import (
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
	}else {
		this.IndexCloudFoundry()
	}
}

func (this *CloudFoundryController) Post() {
	cloudfoundry := entity.NewCloudFoundry(this.GetString("name"))
	this.CommitData(cloudfoundry)
	this.Get()
}

func (this *CloudFoundryController) DeployCloudFoundry(){
	logger.Debug("%s","Deploy CloudFoundry")
	this.LoadData()
	this.TplNames = "cloudfoundry/config.tpl"
}

func (this *CloudFoundryController) IndexCloudFoundry(){
	this.LoadData()
	this.TplNames = "cloudfoundry/index.tpl"
}

func (this *CloudFoundryController) ConfigCloudFoundry(){
	logger.Debug("%s","Config CloudFoundry")
	this.LoadData()
	this.TplNames = "cloudfoundry/config.tpl"
}

//read data from const or database
func (this *CloudFoundryController) LoadData(){
	this.Data["CloudFoundry"] = cf
}

//write data to const or database
func (this *CloudFoundryController) CommitData(cloudfoundry entity.CloudFoundry){
	cf = cloudfoundry
}

var cf entity.CloudFoundry = entity.NewCloudFoundry("deployment-cloudfoundry")




