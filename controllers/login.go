package controllers

import (
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
)

type LoginController struct {
	BaseController
}

type LogoutController struct {
	BaseController
}

func (this *LoginController) Get() {
	if this.GetSession("UserName") != defaultUserName {
		this.TplNames = "login.tpl"
		this.Data["Message"] = ""
		this.Data["Website"] = "www.citycloud.com.cn"
		this.Data["Email"] = "cci-paas@citycloud.com.cn"
		this.Data["LoginAction"] = loginAction
	} else {
		this.Redirect(indexAction, 302)
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
		logger.Debug("%s Login", this.GetString("username"))
		this.Redirect(indexAction, 302)
	}

}

func (this *LogoutController) Logout() {
	logger.Debug("%s", "Logout Successful")
	this.DestroySession()
	this.Redirect(loginAction, 302)
}
