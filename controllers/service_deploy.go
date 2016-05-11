package controllers

import (
	_ "github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
)

type ServiceDeployController struct {
	BaseController
}

func (this *ServiceDeployController) Get() {
	this.Layout = "index.tpl"
	this.Data["NavBarFocus"] = "templates"
	this.Data["IaaSVersion"] = iaasVersion
	this.Data["DefaultVersion"] = defaultVersion

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

	this.TplName = "templates/deploy.tpl"
}
