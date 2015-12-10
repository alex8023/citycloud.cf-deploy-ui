package controllers

import (
	"github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
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

	serviceId, _ := this.GetInt64("serviceId", 0)
	service := entity.Service{}
	service.Id = serviceId
	loaderr := service.Load()

	if loaderr == nil {
		this.Data["Service"] = service
	}

	this.TplNames = "templates/deploy.tpl"
}
