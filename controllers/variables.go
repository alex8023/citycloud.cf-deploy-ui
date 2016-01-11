package controllers

import (
	"github.com/astaxie/beego"
)

var (
	globaleAppName         string = "/" + beego.AppConfig.String("appname")
	defaultUserName        string = beego.AppConfig.String("username")
	defaultPassword        string = beego.AppConfig.String("password")
	loginAction            string = globaleAppName + "/login"
	indexAction            string = globaleAppName + "/index"
	microWebsocket         string = "/microboshwebsocket"
	cloudfoundryWebsocket  string = "/cloudfoundrywebsocket"
	boshWebsocket          string = "/boshwebsocket"
	serviceDeployWebScoket string = "/servicedeploywebsocket"
	workSpace              string = beego.AppConfig.String("workSpace")
	workDir                string = workSpace + "/" + beego.AppConfig.String("workDir")
	microManiest           string = "microbosh/micro_bosh.yml"
	microPath              string = workDir + "/" + microManiest
	cloudFoundryManiest    string = "cloudfoundry/cloudfoundry.yml"
	cloudFoundryPath       string = workDir + "/" + cloudFoundryManiest
	cloudFoundryName       string = beego.AppConfig.String("paasName")
	cloudFoundryVersion    string = beego.AppConfig.String("paasVersion")
	cloudFoundryRelease    string = workSpace + "/" + beego.AppConfig.String("paasRelease")
	iaasVersion            string = beego.AppConfig.String("iaas-version")
	defaultVersion         string = "CCI-IaaS3.0"
	stemcellName           string = beego.AppConfig.String("stemcellName")
	stemcellVersion        string = beego.AppConfig.String("stemcellVersion")
	stemcellRelease        string = workSpace + "/" + beego.AppConfig.String("stemcellRelease")
	customServiceDir       string = beego.AppConfig.String("customServiceDeployDir")
	vsphereVersion         string = "Vsphere"
)
