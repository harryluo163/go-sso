package controllers

import (
	"github.com/astaxie/beego"
	"go-sso/unit"
)

type MainController struct {
	beego.Controller
}

type SuccessController struct {
	beego.Controller
}

type ServiceValidateController struct {
	beego.Controller
}

func (this *MainController) Get() {
	token_head := this.Ctx.GetCookie("TOKEN")
	this.Data["service"] = this.GetString("service")
	if token_head == "" && unit.CheckToken(token_head) == "" {
		this.TplName = "index.html"
	} else {
		this.TplName = "Success.html"
	}

}
func (this *MainController) Post() {
	//获取service
	login_name := this.GetString("userName")
	password := this.GetString("password")
	service := this.GetString("service")
	//returnurl := this.GetString("returnurl")
	//验证用户名密码
	user, err := unit.AuthenticateUserForLogin(login_name, password)
	if user == nil {
		this.Data["json"] = map[string]interface{}{"success": -1, "message": err}
		this.ServeJSON()
		return
	}
	tokenString := unit.CreateToken(login_name)
	//请求中head中设置token
	this.Ctx.Output.Header("TOKEN", tokenString)
	//设置cookie 名称,值,时间,路径
	this.Ctx.SetCookie("TOKEN", tokenString, "3600", "/")
	this.Ctx.SetCookie("USERID", user.Id, "3600", "/")
	this.Data["json"] = map[string]interface{}{"success": 0, "msg": "登录成功", "user": user, "service": service}
	this.ServeJSON()
}
func (c *SuccessController) Get() {
	c.TplName = "Success.html"
}

func (this *ServiceValidateController) Get() {
	ticket := this.GetString("ticket")
	userName := unit.CheckToken(ticket)
	if userName != "" {
		this.Data["json"] = map[string]interface{}{"success": 0, "msg": "登录成功", "user": userName}
	} else {
		this.Data["json"] = map[string]interface{}{"success": -1, "msg": "用户未登录"}
	}
	this.ServeJSON()

}
