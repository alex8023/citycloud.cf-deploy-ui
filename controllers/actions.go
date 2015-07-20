package controllers

import (
	"github.com/astaxie/beego"
)

var (
	globaleAppName  string = "/" + beego.AppConfig.String("appname")
	defaultUserName string = beego.AppConfig.String("username")
	defaultPassword string = beego.AppConfig.String("password")
	loginAction     string = globaleAppName + "/login"
	indexAction     string = globaleAppName + "/index"

	workDir      string = "/home/ubuntu/bosh-workspace/deploy"
	microManiest string = "microbosh/micro_bosh.yml"
	microPath    string = workDir + "/" + microManiest
	stemcells    string = "/home/ubuntu/bosh-workspace/stemcells/bosh-stemcell-2719-openstack-kvm-ubuntu-lucid-go_agent.tgz"

	cloudFoundryManiest string = "cloudfoundry/cloudfoundry.yml"
	cloudFoundryPath    string = workDir + "/" + cloudFoundryManiest
	cloudFoundryRelease string = "/home/ubuntu/bosh-workspace/stemcells/boshrelease-cf-170.tgz"
)
