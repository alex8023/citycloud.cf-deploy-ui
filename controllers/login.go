package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

var globaleAppName = beego.AppConfig.String("appname")
var defaultUserName = beego.AppConfig.String("username")
var defaultPassword = beego.AppConfig.String("password")
var loginAction = "/" + globaleAppName + "/login"
var indexAction = "/" + globaleAppName + "index"

var globaleSession *session.Manager

type LoginController struct {
	beego.Controller
}

type LogoutController struct {
	beego.Controller
}

func (this *LoginController) Get() {
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

func (this *LoginController) Post() {
	if this.GetString("username") != defaultUserName || this.GetString("password") != defaultPassword {
		this.TplNames = "login.tpl"
		this.Data["Message"] = "用户名密码错误"
		this.Data["Website"] = "www.citycloud.com.cn"
		this.Data["Email"] = "cci-paas@citycloud.com.cn"
		this.Data["LoginAction"] = loginAction
	} else {
		this.SetSession("UserName", this.GetString("username"))
		this.Data["UserName"] = this.GetSession("UserName")
		this.TplNames = "index.tpl"
		this.Redirect(indexAction, 302)
	}

}

func (this *LogoutController) Get() {
	this.Ctx.SetCookie("auth", "")
	this.Redirect(loginAction, 301)
}
