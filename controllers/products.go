package controllers

import (
	"github.com/astaxie/beego"
)

type ProductsController struct {
	beego.Controller
}

func (this *ProductsController) Get() {
	if this.isVisitValid() {
		this.Data["IsLogined"] = true
	}
	this.TplNames = "products.html"
}

func (this *ProductsController) isVisitValid() bool {
	uname := GetAccountCookie(this.Ctx)
	s := this.GetSession(uname)
	if CheckAccount(this.Ctx) && s != nil {
		return true
	} else {
		return false
	}
}
