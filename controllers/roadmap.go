package controllers

import (
	"github.com/astaxie/beego"
)

type RoadmapController struct {
	beego.Controller
}

func (this *RoadmapController) Get() {
	if this.isVisitValid() {
		this.Data["IsLogined"] = true
	}
	this.TplNames = "roadmap.html"
}

func (this *RoadmapController) isVisitValid() bool {
	uname := GetAccountCookie(this.Ctx)
	s := this.GetSession(uname)
	if CheckAccount(this.Ctx) && s != nil {
		return true
	} else {
		return false
	}
}
