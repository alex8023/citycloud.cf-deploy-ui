package controllers

import (
	"github.com/astaxie/beego"
)

//  变量声明
var globaleAppName string = "/" + beego.AppConfig.String("appname")
var defaultUserName string = beego.AppConfig.String("username")
var defaultPassword string = beego.AppConfig.String("password")
var loginAction string = globaleAppName + "/login"
var indexAction string = globaleAppName + "/index"
