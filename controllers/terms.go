package controllers

import (
	"github.com/astaxie/beego"
)

type TermsController struct {
	beego.Controller
}

func (this *TermsController) Get() {
	if this.isVisitValid() {
		this.Data["IsLogined"] = true
	}
	this.TplNames = "terms.html"
}

func (this *TermsController) isVisitValid() bool {
	uname := GetAccountCookie(this.Ctx)
	s := this.GetSession(uname)
	if CheckAccount(this.Ctx) && s != nil {
		return true
	} else {
		return false
	}
}
