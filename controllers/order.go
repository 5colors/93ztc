package controllers

import (
	"github.com/astaxie/beego"
)

type OrderController struct {
	beego.Controller
}

func (this *OrderController) Get() {
	if !this.isVisitValid() {
		this.Redirect("/register", 302)
		return
	}
	orderType := this.Input().Get("type")
	beego.Debug("order type:", orderType)
	if orderType == "basic" {
		this.Data["OrderType"] = "总监智囊团"
		this.Data["DailyCost"] = 1.93
	} else if orderType == "pro" {
		this.Data["OrderType"] = "专家智囊团"
		this.Data["DailyCost"] = 4.93
	} else if orderType == "ceo" {
		this.Data["OrderType"] = "总裁智囊团"
		this.Data["DailyCost"] = 9.93
	} else if orderType == "vip" {
		this.Data["OrderType"] = "VIP 计划"
		this.Data["DailyCost"] = 99.93
	}
	this.Data["IsLogined"] = true
	this.TplNames = "orderdetail.html"
}

func (this *OrderController) isVisitValid() bool {
	uname := GetAccountCookie(this.Ctx)
	s := this.GetSession(uname)
	if CheckAccount(this.Ctx) && s != nil {
		return true
	} else {
		return false
	}
}
