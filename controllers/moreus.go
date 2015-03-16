package controllers

import (
	"github.com/astaxie/beego"
)

type MoreUsController struct {
	beego.Controller
}

func (this *MoreUsController) Get() {
	if this.isVisitValid() {
		this.Data["IsLogined"] = true
	}
	this.TplNames = "moreus.html"
}

func (this *MoreUsController) isVisitValid() bool {
	uname := GetAccountCookie(this.Ctx)
	s := this.GetSession(uname)
	if CheckAccount(this.Ctx) && s != nil {
		return true
	} else {
		return false
	}
}
