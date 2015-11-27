package controllers

import (
	"github.com/astaxie/beego"
)

type ServiceDeployController struct {
	beego.Controller
}

func (this *ServiceDeployController) Get() {
	this.Layout = "index.tpl"
	this.Data["NavBarFocus"] = "templates"

	this.Data["HOST"] = this.Ctx.Request.Host
	this.Data["AppName"] = globaleAppName
	this.Data["WebSocket"] = serviceDeployWebScoket
	this.Data["ServiceId"] = this.GetString("serviceId")

	this.TplNames = "templates/deploy.tpl"
}
