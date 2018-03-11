// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/gnenux/pic-bed/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/image",
			beego.NSInclude(
				&controllers.ImgaeController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.Router("/", &controllers.HomeController{})
	beego.InsertFilter("/static/*", beego.BeforeStatic, func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
		ctx.Output.Header("Access-Control-Allow-Methods", "GET, HEAD, OPTIONS")
	})
}
