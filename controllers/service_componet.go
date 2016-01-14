package controllers

import (
	"fmt"
	_ "github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
)

type ServiceComponentController struct {
	BaseController
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

type ServiceOperationController struct {
	BaseController
}

func (this *ServiceOperationController) Get() {
	id, err := this.GetInt64("sid", 0)
	if id == 0 {
		this.Data["json"] = ""
		this.ServeJson()
	}
	operation := entity.Operation{}
	if err == nil {
		operation.Id = id
		err = operation.Load()
		if err == nil {
			this.Data["json"] = &operation
			this.ServeJson()
		} else {
			this.Data["json"] = ""
			this.ServeJson()
		}
	}
}

func (this *ServiceOperationController) Post() {
	sid, _ := this.GetInt64("sid", 0)
	id, _ := this.GetInt64("id", 0)
	operation := entity.Operation{}
	operation.Id = id
	operation.Sid = sid

	operation.Start = this.GetString("start")
	operation.Restart = this.GetString("restart")
	operation.Stop = this.GetString("stop")

	if id == 0 {
		operation.Save()
	} else {
		operation.Update()
	}

	redirectUrl := fmt.Sprintf("templatesdetail?serviceId=%d", sid)
	this.Redirect(redirectUrl, 301) //this.Get()
}
