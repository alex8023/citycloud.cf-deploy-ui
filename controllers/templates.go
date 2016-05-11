package controllers

import (
	"fmt"
	_ "github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
)

type TemplatesController struct {
	BaseController
}

func (this *TemplatesController) Get() {
	action := this.GetString("action")
	this.Layout = "index.tpl"
	this.Data["NavBarFocus"] = "templates"
	this.Data["IaaSVersion"] = iaasVersion
	this.Data["DefaultVersion"] = defaultVersion

	switch action {
	case "", "list":
		this.List()
	case "detail":
		this.LoadCustomServiceDetail()
	default:
		this.List()
	}

}

func (this *TemplatesController) Post() {
	action := this.GetString("action")
	switch action {
	case "":
		service := entity.Service{}
		service.Name = this.GetString("name")
		service.Description = this.GetString("description")
		service.Where = this.GetString("where")

		id, err := this.GetInt64("id", 0)
		if err == nil {
			if id == 0 {
				service.Save()
				if service.Where == utils.Deploy_On_PaaS {
					deployOnPaaS := entity.OnPaaS{}
					deployOnPaaS.Api = this.GetString("Api")
					deployOnPaaS.Sid = service.Id
					deployOnPaaS.User = this.GetString("UserName")
					deployOnPaaS.Passwd = this.GetString("Passwd")
					deployOnPaaS.Org = this.GetString("Org")
					deployOnPaaS.Space = this.GetString("Space")
					deployOnPaaS.Save()
				}
				if service.Where == utils.Deploy_On_Vms {
					deployOnCustom := entity.OnCustom{}
					deployOnCustom.Sid = service.Id
					deployOnCustom.Ip = this.GetString("ip")
					deployOnCustom.User = this.GetString("VmUserName")
					deployOnCustom.Passwd = this.GetString("VmPasswd")
					deployOnCustom.Save()
				}
			} else {
				service.Id = id
				service.Update()
				if service.Where == utils.Deploy_On_PaaS {
					paasid, _ := this.GetInt64("paasid", 0)
					deployOnPaaS := entity.OnPaaS{}
					deployOnPaaS.Id = paasid
					deployOnPaaS.Api = this.GetString("Api")
					deployOnPaaS.Sid = service.Id
					deployOnPaaS.User = this.GetString("UserName")
					deployOnPaaS.Passwd = this.GetString("Passwd")
					deployOnPaaS.Org = this.GetString("Org")
					deployOnPaaS.Space = this.GetString("Space")
					deployOnPaaS.Update()
				}
				if service.Where == utils.Deploy_On_Vms {
					vmid, _ := this.GetInt64("vmid", 0)
					deployOnCustom := entity.OnCustom{}
					deployOnCustom.Id = vmid
					deployOnCustom.Sid = service.Id
					deployOnCustom.Ip = this.GetString("ip")
					deployOnCustom.User = this.GetString("VmUserName")
					deployOnCustom.Passwd = this.GetString("VmPasswd")
					deployOnCustom.Update()
				}
			}
		}

	case "delete":
		this.DeleteCustomService()
	}
	this.Redirect("templates", 301) //this.Get()
}

func (this *TemplatesController) List() {
	this.TplName = "templates/index.tpl"
	service, err := entity.LoadServiceList()
	if err != nil {
		this.Data["MessageErr"] = fmt.Sprintf("Errors: %s", err)
	}
	this.Data["Service"] = service
}

func (this *TemplatesController) LoadCustomServiceDetail() {
	id, err := this.GetInt64("id", 0)
	service := entity.Service{}
	if err == nil {
		service.Id = id
		err = service.Load()
		serviceDto := entity.ServiceDto{}
		serviceDto.Service = service
		err = serviceDto.Load()
		if err == nil {
			this.Data["json"] = &serviceDto
			this.ServeJSON()
		}
	}
}

func (this *TemplatesController) DeleteCustomService() {
	id, err := this.GetInt64("id", 0)
	service := entity.Service{}
	if err == nil && id != 0 {
		service.Id = id
		err = service.Load()
		if err == nil {
			serviceDto := entity.ServiceDto{}
			serviceDto.Service = service
			serviceDto.Load()
			if serviceDto.Service.Where == utils.Deploy_On_PaaS {
				serviceDto.OnPaaS.Delete()
			}
			if serviceDto.Service.Where == utils.Deploy_On_Vms {
				serviceDto.OnCustom.Delete()
			}
		}

		err = service.Delete()
		if err != nil {
			this.Data["MessageErr"] = fmt.Sprintf("Errors: %s", err)
		}
	}
	if err != nil {
		this.Data["MessageErr"] = fmt.Sprintf("Errors: %s", err)
	}
}
