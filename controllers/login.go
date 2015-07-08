package controllers

import (
	"log"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

var globaleAppName = beego.AppConfig.String("appname")
var defaultUserName = beego.AppConfig.String("username")
var defaultPassword = beego.AppConfig.String("password")
var loginAction = "/" + globaleAppName + "/login"

var globaleSession *session.Manager

type LoginController struct {
	beego.Controller
}

type LogoutController struct {
	beego.Controller
}


func (this *LoginController) Get() {
	if this.GetSession("UserName") != defaultUserName {
		this.TplNames = "login.tpl"
		this.Data["Message"] = ""
		this.Data["Website"] = "www.citycloud.com.cn"
		this.Data["Email"] = "cci-paas@citycloud.com.cn"
		this.Data["LoginAction"] = loginAction
	} else {
		this.Redirect("/citycloud.cf-deploy-ui/index",302)
	}
}

func (this *LoginController) Post() {
	if this.GetString("username") != defaultUserName || this.GetString("password") != defaultPassword {
		this.TplNames = "login.tpl"
		this.Data["Message"] = "用户名密码错误"
		this.Data["Website"] = "www.citycloud.com.cn"
		this.Data["Email"] = "cci-paas@citycloud.com.cn"
		this.Data["LoginAction"] = loginAction
	} else {
		this.SetSession("UserName", this.GetString("username"))
		this.Redirect("/citycloud.cf-deploy-ui/index",302)
	}

}

func (this *LogoutController) Logout() {
	log.Println("logout successful")
	this.DestroySession()
	this.Redirect(loginAction, 302)
}