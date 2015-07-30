package controllers

import (
	"github.com/astaxie/beego"
)

var (
	globaleAppName        string = "/" + beego.AppConfig.String("appname")
	defaultUserName       string = beego.AppConfig.String("username")
	defaultPassword       string = beego.AppConfig.String("password")
	loginAction           string = globaleAppName + "/login"
	indexAction           string = globaleAppName + "/index"
	microWebsocket        string = "/microboshwebsocket"
	cloudfoundryWebsocket string = "/cloudfoundrywebsocket"
	boshWebsocket         string = "/boshwebsocket"

	workDir      string = "/home/ubuntu/bosh-workspace/dephloy"
	microManiest string = "microbosh/micro_bosh.yml"
	microPath    string = workDir + "/" + microManiest

	stemcells string = beego.AppConfig.String("stemcells")

	cloudFoundryManiest string = "cloudfoundry/cloudfoundry.yml"
	cloudFoundryPath    string = workDir + "/" + cloudFoundryManiest
	cloudFoundryRelease string = beego.AppConfig.String("paas-release")
)
