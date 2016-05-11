package controllers

import (
	_ "github.com/astaxie/beego"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	action := this.GetString("action")
	this.Data["IaaSVersion"] = iaasVersion
	this.Data["DefaultVersion"] = defaultVersion
	if action == "home" || action == "" {
		this.Data["NavBarFocus"] = "home"
		this.TplName = "home/home.tpl"
	} else if action == "cloudfoundry" {
		this.initPaaS()
	} else if action == "bosh" {
		this.initBOSH()
	} else {
		this.unRelease()
	}
	this.Layout = "index.tpl"
}

func (this *IndexController) initPaaS() {
	this.Data["NavBarFocus"] = "cloudfoundry"
	this.TplName = "cloudfoundry/config.tpl"
}

func (this *IndexController) initBOSH() {
	this.Data["NavBarFocus"] = "bosh"
	this.TplName = "bosh/config.tpl"
}

func (this *IndexController) unRelease() {
	this.Data["NavBarFocus"] = this.GetString("action")
	this.TplName = "undo/undo.tpl"
}
