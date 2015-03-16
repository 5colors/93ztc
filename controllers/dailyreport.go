package controllers

import (
	"github.com/astaxie/beego"
)

type DailyReportController struct {
	beego.Controller
}

func (this *DailyReportController) Get() {
	if this.isVisitValid() {
		this.Data["IsLogined"] = true
		this.TplNames = "dailyreport.html"
	} else {
		this.Redirect("/", 302)
		return
	}

}

func (this *DailyReportController) isVisitValid() bool {
	uname := GetAccountCookie(this.Ctx)
	beego.Debug("uname in dailyreport:", uname)
	s := this.GetSession(uname)
	if CheckAccount(this.Ctx) && s != nil {
		return true
	} else {
		return false
	}
}
