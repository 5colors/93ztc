package controllers

import (
	"github.com/astaxie/beego"
)

type AboutusController struct {
	beego.Controller
}

func (this *AboutusController) Get() {
	if this.isVisitValid() {
		this.Data["IsLogined"] = true
	}
	this.TplNames = "aboutus.html"
}

func (this *AboutusController) isVisitValid() bool {
	uname := GetAccountCookie(this.Ctx)
	s := this.GetSession(uname)
	if CheckAccount(this.Ctx) && s != nil {
		return true
	} else {
		return false
	}
}
