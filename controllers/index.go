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
		initHome(this)
	} else if action == "microbosh" {
		initDeployMicroBosh(this)
	} else if action == "paas" {
		initPaaS(this)
	} else {
		unRelease(this)
	}
	this.Layout = "index.tpl"
}

func initHome(this *IndexController) {
	this.Data["NavBarFocus"] = "home"
	this.TplNames = "home/home.tpl"
}

func initDeployMicroBosh(this *IndexController) {
	this.Data["NavBarFocus"] = "microbosh"
	this.TplNames = "microbosh/config.tpl"
}

func initPaaS(this *IndexController) {
	this.Data["NavBarFocus"] = "paas"
	this.TplNames = "paas/config.tpl"
}

func unRelease(this *IndexController) {
	this.Data["NavBarFocus"] = "home"
	this.TplNames = "undo/undo.tpl"
}
