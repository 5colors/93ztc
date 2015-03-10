package controllers

import (
	"github.com/astaxie/beego"
)

type ArticlesController struct {
	beego.Controller
}

func (this *ArticlesController) Get() {
	if this.isVisitValid() {
		this.Data["IsLogined"] = true
	}
	this.TplNames = "articles.html"
}

func (this *ArticlesController) isVisitValid() bool {
	uname := GetAccountCookie(this.Ctx)
	s := this.GetSession(uname)
	if CheckAccount(this.Ctx) && s != nil {
		return true
	} else {
		return false
	}
}
