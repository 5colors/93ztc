package controllers

import (
	"github.com/astaxie/beego"
)

type FaqController struct {
	beego.Controller
}

func (this *FaqController) Get() {
	if this.isVisitValid() {
		this.Data["IsLogined"] = true
	}
	this.TplNames = "faq.html"
}

func (this *FaqController) isVisitValid() bool {
	uname := GetAccountCookie(this.Ctx)
	s := this.GetSession(uname)
	if CheckAccount(this.Ctx) && s != nil {
		return true
	} else {
		return false
	}
}
