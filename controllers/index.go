package controllers

import (
	"log"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
		this.Data["UserName"] = this.GetSession("UserName")
		Get(this.GetSession("UserName"))
		this.TplNames = "index.tpl"
}

func Get(username interface{} ) {
	log.Println("hello")
	log.Println(username)
}
