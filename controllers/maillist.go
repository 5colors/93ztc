package controllers

import (
	"github.com/astaxie/beego"
)

type MailListController struct {
	beego.Controller
}

func (this *MailListController) Get() {
	if this.isVisitValid() {
		this.Data["IsLogined"] = true
	}
	this.TplNames = "moreus.html"
}

func (this *MailListController) Post() {
	if this.isVisitValid() {
		this.Data["IsLogined"] = true
	}
	email := this.Input().Get("email")
	beego.Debug("email for subscribe:", email)
	AddNewSubcribeEmail(email)
	this.Data["IsEmailReg"] = true
	this.TplNames = "moreus.html"
}

func (this *MailListController) isVisitValid() bool {
	uname := GetAccountCookie(this.Ctx)
	s := this.GetSession(uname)
	if CheckAccount(this.Ctx) && s != nil {
		return true
	} else {
		return false
	}
}
