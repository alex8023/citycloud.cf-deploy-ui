package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
)

type TemplatesControllers struct {
	beego.Controller
}

func (this *TemplatesControllers) Get() {
	action := this.GetString("action")
	this.Layout = "index.tpl"
	this.Data["NavBarFocus"] = "templates"

	switch action {
	case "", "list":
		this.List()
	case "config":
	case "deploy":
	default:
	}

}

func (this *TemplatesControllers) Post() {

}

func (this *TemplatesControllers) List() {
	this.TplNames = "templates/index.tpl"
	service, err := entity.LoadServiceList()
	if err != nil {
		this.Data["MessageErr"] = fmt.Sprintf("Errors: %s", err)
	}
	this.Data["Service"] = service
}
