package controllers

import (
	"github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
)

type BOSHController struct {
	beego.Controller
}

func (this *BOSHController) Get() {
	action := this.GetString("action")
	this.Layout = "index.tpl"
	this.Data["NavBarFocus"] = "bosh"
	this.Data["BOSH"] = bo
	if action == "config" {
		this.ConfigBOSH()
	} else if action == "deploy" {
		this.DeployBOSH()
	} else {
		this.IndexBOSH()
	}
}

func (this *BOSHController) Post() {
	bosh := entity.NewSimpleBOSH(this.GetString("name"))
	this.CommitData(bosh)
	this.Get()
}

func (this *BOSHController) DeployBOSH() {
	this.LoadData()
	this.TplNames = "bosh/config.tpl"
}

func (this *BOSHController) IndexBOSH() {
	this.LoadData()
	this.TplNames = "bosh/index.tpl"
}

func (this *BOSHController) ConfigBOSH() {
	this.LoadData()
	this.TplNames = "bosh/config.tpl"
}

func (this *BOSHController) LoadData() {
	this.Data["BOSH"] = bo
}

func (this *BOSHController) CommitData(bosh entity.BOSH) {
	bo = bosh
}

var bo entity.BOSH = entity.NewSimpleBOSH("deployment-bosh")
