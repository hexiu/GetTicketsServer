// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"net/http"
	"strings"

	"12306_server/controllers"
	"12306_server/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	ns := beego.NewNamespace("/v1",
		//登录

		beego.NSRouter("/auth/login", &controllers.UserController{}, "Post:Login"),
		beego.NSRouter("/auth/verifyCode", &controllers.UserController{}, "Get:VerifyCode"),
		beego.NSRouter("/auth/init", &controllers.UserController{}, "Get:InitLogin"),
		//车次处理
		beego.NSNamespace("/schedule",
			beego.NSBefore(Auth),
			beego.NSInclude(
				&controllers.ScheduleController{},
			),
		),
		//站台处理
		beego.NSNamespace("/station",
			beego.NSBefore(Auth),
			beego.NSInclude(
				&controllers.StationController{},
			),
		),
		//乘客信息
		beego.NSNamespace("/passenger",
			beego.NSBefore(Auth),
			beego.NSInclude(
				&controllers.PassengerController{},
			),
		),
		//任务管理
		beego.NSNamespace("/task",
			beego.NSBefore(Auth),
			beego.NSInclude(
				&controllers.TaskController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

func Auth(ctx *context.Context) {
	//只检测OPTIONS以外的请求
	if !ctx.Input.Is("OPTIONS") {
		authString := ctx.Input.Header("Authorization")
		if authString == "" {
			AllowCross(ctx)
			return
		}
		kv := strings.Split(authString, " ")
		if len(kv) != 2 || kv[0] != "Bearer" {
			AllowCross(ctx)
			return
		}
		token := kv[1]
		jwt := utils.InitJwt()
		if !jwt.Checkd(token) {
			AllowCross(ctx)
			return
		}
	}
}

//错误返回
func AllowCross(ctx *context.Context) {
	ctx.Output.Header("Cache-Control", "no-store")
	ctx.Output.Header("Access-Control-Allow-Origin", "*")
	ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE,OPTIONS")
	ctx.Output.Header("Access-Control-Allow-Headers", "Authorization")
	ctx.Output.Header("WWW-Authenticate", `Bearer realm="`+beego.AppConfig.String("HostName")+`" error="Authorization" error_description="invalid Authorization"`)
	http.Error(ctx.ResponseWriter, "Unauthorized", 401)
}
