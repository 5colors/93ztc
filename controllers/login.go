package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	if this.isVisitValid() {
		this.Redirect("/", 302)
		return
	}
	this.TplNames = "login.html"

}

func (this *LoginController) Logout() {
	beego.Debug("exiting...")
	this.Ctx.SetCookie("uname", "", "/")
	this.DelSession(GetAccountCookie(this.Ctx))
	this.Redirect("/", 302)
	return
}

func (this *LoginController) Post() {
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	if CheckUserPassword(uname, pwd) {
		this.Ctx.SetCookie("uname", uname, "/")
		this.SetSession(uname, int(1))
		this.Redirect("/dailyreport", 302)
	} else {
		beego.Debug("not valid user login")
		this.Data["NotValidUser"] = true
		this.TplNames = "login.html"
	}
}

func (this *LoginController) isVisitValid() bool {
	uname := GetAccountCookie(this.Ctx)
	s := this.GetSession(uname)
	if CheckAccount(this.Ctx) && s != nil {
		return true
	} else {
		return false
	}
}
