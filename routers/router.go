package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/citycloud/citycloud.cf-deploy-ui/controllers"
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
					ctx.Redirect(301, loginAction)
				}

			}),
			beego.NSRouter("/login", &controllers.LoginController{}),
			beego.NSRouter("/logout", &controllers.LogoutController{}, "*:Logout"),
			beego.NSRouter("/index", &controllers.IndexController{}),
			beego.NSRouter("/microbosh", &controllers.MicroBoshController{}),
			beego.NSRouter("/bosh", &controllers.BOSHController{}),
			beego.NSRouter("/cloudfoundry", &controllers.CloudFoundryController{}),
			beego.NSRouter("/microboshwebsocket", &controllers.MicroBOSHWebSocketController{}),
			beego.NSRouter("/cloudfoundrywebsocket", &controllers.CloudFoundryWebSocketController{}),
			beego.NSRouter("/templates", &controllers.TemplatesController{}),
			beego.NSRouter("/templatesdetail", &controllers.TemplatesDetailController{}),
			beego.NSRouter("/servicedeploy", &controllers.ServiceDeployController{}),
			beego.NSRouter("/servicedeploywebsocket", &controllers.ServiceDeployWebScoketController{}),
		)
	beego.AddNamespace(login)
}
