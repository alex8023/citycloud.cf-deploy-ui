package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	action := this.GetString("action")
	if action == "home" || action == "" {
		this.Data["NavBarFocus"] = "home"
		this.TplNames = "home/home.tpl"
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
	this.TplNames = "cloudfoundry/config.tpl"
}

func (this *IndexController) initBOSH() {
	this.Data["NavBarFocus"] = "bosh"
	this.TplNames = "bosh/config.tpl"
}


func (this *IndexController) unRelease() {
	this.Data["NavBarFocus"] = this.GetString("action")
	this.TplNames = "undo/undo.tpl"
}
