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
	} else if action == "paas" {
		initPaaS(this)
	} else {
		unRelease(this)
	}
	this.Layout = "index.tpl"
}

func initPaaS(this *IndexController) {
	this.Data["NavBarFocus"] = "paas"
	this.TplNames = "paas/config.tpl"
}

func unRelease(this *IndexController) {
	this.Data["NavBarFocus"] = "home"
	this.TplNames = "undo/undo.tpl"
}
