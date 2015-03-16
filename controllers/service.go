package controllers

import (
	"github.com/astaxie/beego"
)

type ServiceController struct {
	beego.Controller
}

func (this *ServiceController) Get() {
	if this.isVisitValid() {
		this.Data["IsLogined"] = true
	}
	this.TplNames = "service.html"
}

func (this *ServiceController) isVisitValid() bool {
	uname := GetAccountCookie(this.Ctx)
	s := this.GetSession(uname)
	if CheckAccount(this.Ctx) && s != nil {
		return true
	} else {
		return false
	}
}
