package controllers

import (
	"github.com/astaxie/beego"
)

var (
	globaleAppName string = "/" + beego.AppConfig.String("appname")
	defaultUserName string = beego.AppConfig.String("username")
 	defaultPassword string = beego.AppConfig.String("password")
	loginAction string = globaleAppName + "/login"
 	indexAction string = globaleAppName + "/index"
)

