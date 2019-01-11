package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["12306_server/controllers:PassengerController"] = append(beego.GlobalControllerRouter["12306_server/controllers:PassengerController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["12306_server/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["12306_server/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["12306_server/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["12306_server/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "InitQuery",
			Router: `/init`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["12306_server/controllers:StationController"] = append(beego.GlobalControllerRouter["12306_server/controllers:StationController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["12306_server/controllers:TaskController"] = append(beego.GlobalControllerRouter["12306_server/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["12306_server/controllers:TaskController"] = append(beego.GlobalControllerRouter["12306_server/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["12306_server/controllers:TaskController"] = append(beego.GlobalControllerRouter["12306_server/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Log",
			Router: `/log/:id:int`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
