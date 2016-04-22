package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"strings"
)

type BaseController struct {
	beego.Controller
	i18n.Locale
}

func (this *BaseController) Prepare() {
	lang := this.Input().Get("lang")
	if len(lang) == 0 {
		lang = this.Ctx.GetCookie("lang")
	}
	if len(lang) == 0 {
		al := this.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5]
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}
	if len(lang) == 0 {
		lang = "zh-CN"
		fmt.Println("zh-CN")
	}
	this.Data["Lang"] = lang
}

type langType struct {
	Lang string
	Name string
}

func init() {
	langs := strings.Split(beego.AppConfig.String("lang::types"), "|")
	names := strings.Split(beego.AppConfig.String("lang::names"), "|")
	langTypes := make([]*langType, 0, len(langs))
	for i, v := range langs {
		langTypes = append(langTypes, &langType{
			Lang: v,
			Name: names[i],
		})
	}

	for _, lang := range langs {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			return
		}
	}
	beego.AddFuncMap("i18n", i18n.Tr)
	
}
