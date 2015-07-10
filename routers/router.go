package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/citycloud/citycloud.cf-deploy-ui/controllers"
	"log"
)

var globaleAppName = beego.AppConfig.String("appname")
var defaultUserName = beego.AppConfig.String("username")

func init() {
	appName := "/" + globaleAppName
	login :=
		beego.NewNamespace(appName,
			beego.NSBefore(func(ctx *context.Context) {
				username := ctx.Input.Session("UserName")
				loginAction := appName + "/login"
				if username != defaultUserName && ctx.Request.RequestURI != loginAction {
					log.Println("logout successful")
					ctx.Redirect(301, loginAction)
				}

			}),
			beego.NSRouter("/login", &controllers.LoginController{}),
			beego.NSRouter("/logout", &controllers.LogoutController{}, "*:Logout"),
			beego.NSRouter("/index", &controllers.IndexController{}),
			beego.NSRouter("/microbosh", &controllers.MicroBoshController{}, "post:Post"),
		)
	beego.AddNamespace(login)
}
