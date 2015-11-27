package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
)

type ServiceComponentController struct {
	beego.Controller
}

func (this *ServiceComponentController) Get() {
	id, err := this.GetInt64("id", 0)
	component := entity.Component{}
	if err == nil {
		component.Id = id
		err = component.Load()
		if err == nil {
			this.Data["json"] = &component
			this.ServeJson()
		}
	}
}

func (this *ServiceComponentController) Post() {
	action := this.GetString("action")
	sid, _ := this.GetInt64("sid", 0)
	switch action {
	case "":
		component := entity.Component{}
		component.Sid = sid
		component.Name = this.GetString("name")
		component.Value = this.GetString("value")
		id, err := this.GetInt64("id", 0)
		if err == nil {
			if id == 0 {
				component.Save()
			} else {
				component.Id = id
				component.Update()
			}
		}

	case "delete":
		this.DeleteComponent()
	}
	redirectUrl := fmt.Sprintf("templatesdetail?serviceId=%d", sid)
	this.Redirect(redirectUrl, 301) //this.Get()
}

func (this *ServiceComponentController) DeleteComponent() {
	id, err := this.GetInt64("id", 0)
	component := entity.Component{}
	if err == nil && id != 0 {
		component.Id = id
		err = component.Delete()
		if err != nil {
			this.Data["MessageErr"] = fmt.Sprintf("Errors: %s", err)
		}
	}
	if err != nil {
		this.Data["MessageErr"] = fmt.Sprintf("Errors: %s", err)
	}
}
