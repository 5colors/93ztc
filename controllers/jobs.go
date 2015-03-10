package controllers

import (
	"github.com/astaxie/beego"
)

type JobsController struct {
	beego.Controller
}

func (this *JobsController) Get() {
	if this.isVisitValid() {
		this.Data["IsLogined"] = true
	}
	this.TplNames = "jobs.html"
}

func (this *JobsController) isVisitValid() bool {
	uname := GetAccountCookie(this.Ctx)
	s := this.GetSession(uname)
	if CheckAccount(this.Ctx) && s != nil {
		return true
	} else {
		return false
	}
}
