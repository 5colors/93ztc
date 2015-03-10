package controllers

import (
	"github.com/astaxie/beego"
)

type MonthlyReportController struct {
	beego.Controller
}

func (this *MonthlyReportController) Get() {
	if this.isVisitValid() {
		this.Data["IsLogined"] = true
		this.TplNames = "monthlyreport.html"
	} else {
		this.Redirect("/", 302)
		return
	}

}

func (this *MonthlyReportController) isVisitValid() bool {
	uname := GetAccountCookie(this.Ctx)
	s := this.GetSession(uname)
	if CheckAccount(this.Ctx) && s != nil {
		return true
	} else {
		return false
	}
}
