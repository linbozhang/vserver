// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"vitoserver/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/v1/login",&controllers.LoginController{})
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/app_info",
			beego.NSInclude(
				&controllers.AppInfoController{},
			),
		),

		beego.NSNamespace("/file_info",
			beego.NSInclude(
				&controllers.FileInfoController{},
			),
		),

		beego.NSNamespace("/user_app_info",
			beego.NSInclude(
				&controllers.UserAppInfoController{},
			),
		),

		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
