package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	if this.GetSession("username") != defaultUserName {
		this.TplNames = "login.tpl"
		this.Data["Message"] = ""
		this.Data["Website"] = "www.citycloud.com.cn"
		this.Data["Email"] = "cci-paas@citycloud.com.cn"
		this.Data["LoginAction"] = loginAction
	} else {
		this.Data["Message"] = ""
		this.Data["UserName"] = this.GetSession("UserName")
		this.TplNames = "index.tpl"
	}
}
