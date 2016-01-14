package controllers

import (
	"fmt"
	_ "github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
)

type TemplatesDetailController struct {
	BaseController
}

func (this *TemplatesDetailController) Get() {
	this.Layout = "index.tpl"
	this.Data["NavBarFocus"] = "templates"
	this.Data["IaaSVersion"] = iaasVersion
	this.Data["DefaultVersion"] = vsphereVersion

	action := this.GetString("action")
	switch action {
	case "", "list":
		serviceId, err := this.GetInt64("serviceId", 0)
		if err != nil {
			this.TplNames = "templates/404.tpl"
			return
		}
		if this.CheckService(serviceId) {
			this.List(serviceId)
		}
	case "detail":
		this.LoadCustomTemplateDetail()
	default:
		this.TplNames = "templates/404.tpl"
	}
}

func (this *TemplatesDetailController) CheckService(serviceId int64) bool {

	this.Data["ServiceId"] = serviceId
	if serviceId == 0 {
		this.TplNames = "templates/404.tpl"
		return false
	}
	service := entity.Service{}
	service.Id = serviceId
	loaderr := service.Load()
	if loaderr != nil {
		this.TplNames = "templates/404.tpl"
		return false
	}
	this.Data["Service"] = service
	return true
}

func (this *TemplatesDetailController) List(serviceId int64) {
	service := entity.Service{}
	service.Id = serviceId
	service.Load()

	template, errs := entity.LoadTemplateList(serviceId)
	this.Data["Template"] = template

	if errs != nil {
		this.Data["MessageErr"] = fmt.Sprintf("Errors: %s", errs)
	}

	component, componenterrors := entity.LoadComponentList(serviceId)

	this.Data["Component"] = component
	if componenterrors != nil {
		this.Data["MessageErr"] = fmt.Sprintf("Errors: %s", componenterrors)
	}

	if service.Where == utils.Deploy_On_Vms {
		operation := entity.Operation{}
		operation.Sid = serviceId
		operationerrors := operation.LoadBySid()

		this.Data["Operation"] = operation
		if operationerrors != nil {
			this.Data["MessageErr"] = fmt.Sprintf("Errors: %s", operationerrors)
		}
	}
	this.TplNames = "templates/index_detail.tpl"
}

func (this *TemplatesDetailController) LoadCustomTemplateDetail() {
	id, err := this.GetInt64("id", 0)
	template := entity.Template{}
	if err == nil {
		template.Id = id
		err = template.Load()
		if err == nil {
			this.Data["json"] = &template
			this.ServeJson()
		}
	}
}

func (this *TemplatesDetailController) Post() {
	action := this.GetString("action")
	sid, _ := this.GetInt64("sid", 0)
	switch action {
	case "":
		template := entity.Template{}
		template.Sid = sid
		template.Name = this.GetString("name")
		template.Description = this.GetString("description")
		template.FileType = this.GetString("fileType")
		template.TemplateFile = this.GetString("templatefile")
		template.TargetFile = this.GetString("targetfile")
		id, err := this.GetInt64("id", 0)
		if err == nil {
			if id == 0 {
				template.Save()
			} else {
				template.Id = id
				template.Update()
			}
		}

	case "delete":
		this.DeleteCustomTemplate()
	}
	redirectUrl := fmt.Sprintf("templatesdetail?serviceId=%d", sid)
	this.Redirect(redirectUrl, 301) //this.Get()
}

func (this *TemplatesDetailController) DeleteCustomTemplate() {
	id, err := this.GetInt64("id", 0)
	template := entity.Template{}
	if err == nil && id != 0 {
		template.Id = id
		err = template.Delete()
		if err != nil {
			this.Data["MessageErr"] = fmt.Sprintf("Errors: %s", err)
		}
	}
	if err != nil {
		this.Data["MessageErr"] = fmt.Sprintf("Errors: %s", err)
	}
}
