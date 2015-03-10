package controllers

import (
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplNames = "register.html"
}

func (this *RegisterController) Post() {
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	source := this.Input().Get("source")
	beego.Debug("uanme in register:", uname)
	beego.Debug("pwd in register:", pwd)
	beego.Debug("source in register:", source)
	if CheckUserExist(uname) != true {
		AddNewUser(uname, pwd)
		//保存会话并重定向
		this.SetSession(uname, int(1))
		this.Ctx.SetCookie("uname", uname, "/")
		this.Redirect("/account#accinfo", 302)
		return
	} else {
		this.Data["DupUser"] = true
		this.TplNames = "register.html"
	}
}
