package main

import (
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//	_ "github.com/citycloud/citycloud.cf-deploy-ui/routers"
	_ "github.com/citycloud/citycloud.cf-deploy-ui/entity"
	_ "github.com/citycloud/citycloud.cf-deploy-ui/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//	beego.Run()
	initDB()

}

func initDB() {
	orm.RunCommand()
}
