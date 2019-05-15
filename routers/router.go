// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"go-sso/controllers"
	"go-sso/unit"
	"strings"
)

func init() {

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
		cookie, err := ctx.Request.Cookie("TOKEN")
		fmt.Print(strings.Index(ctx.Request.RequestURI, "/login"))
		if strings.Index(ctx.Request.RequestURI, "/login") >= 0 || strings.Index(ctx.Request.RequestURI, "/static") >= 0 {

		} else if err != nil || unit.CheckToken(cookie.Value) == "" {
			ctx.Redirect(302, "/login")
		}

	})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{})                      //授权登录
	beego.Router("/Success", &controllers.SuccessController{})                 //登录成功
	beego.Router("/serviceValidate", &controllers.ServiceValidateController{}) //验证登录信息

	//另外一种写法
	//初始化 namespace
	//ns2 := beego.NewNamespace("/cas",
	//	beego.NSRouter("/", &controllers.MainController{}),
	//	beego.NSRouter("/login", &controllers.MainController{}),//授权登录
	//	beego.NSRouter("/Success", &controllers.SuccessController{}),
	//	beego.NSRouter("/serviceValidate", &controllers.SuccessController{}),//验证登录信息
	//)
	//beego.AddNamespace(ns2)
}
