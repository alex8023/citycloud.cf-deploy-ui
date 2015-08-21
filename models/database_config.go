package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	mysqlHost := beego.AppConfig.String("mysql_host")
	mysqlPort := beego.AppConfig.String("mysql_port")
	mysqlUserName := beego.AppConfig.String("mysql_username")
	mysqlPasswd := beego.AppConfig.String("mysql_password")
	mysqlDatabase := beego.AppConfig.String("mysql_database")
	if mysqlPort == "" {
		mysqlPort = "3306"
	}
	mysqlUrl := mysqlUserName + ":" + mysqlPasswd + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDatabase + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", mysqlUrl)
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}
