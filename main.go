package main

import (
	"github.com/astaxie/beego"
	_ "github.com/citycloud/citycloud.cf-deploy-ui/routers"
)

func main() {
	beego.Run()
}
