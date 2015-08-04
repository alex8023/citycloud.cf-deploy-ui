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

	workSpace    string = beego.AppConfig.String("workSpace")
	workDir      string = workSpace + "/" + beego.AppConfig.String("workDir")
	microManiest string = "microbosh/micro_bosh.yml"
	microPath    string = workDir + "/" + microManiest

	stemcells string = workSpace + "/" + beego.AppConfig.String("stemcells")

	cloudFoundryManiest string = "cloudfoundry/cloudfoundry.yml"
	cloudFoundryPath    string = workDir + "/" + cloudFoundryManiest
	cloudFoundryRelease string = workSpace + "/" + beego.AppConfig.String("paas-release")

	iaasVersion    = beego.AppConfig.String("iaas-version")
	defaultVersion = "CCI-IaaS3.0"

	paasName    = beego.AppConfig.String("paasName")
	paasVersion = beego.AppConfig.String("paasVersion")

	stemcellName    = beego.AppConfig.String("stemcellName")
	stemcellVersion = beego.AppConfig.String("stemcellVersion")
)
