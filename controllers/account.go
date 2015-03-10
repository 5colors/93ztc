package controllers

import (
	"github.com/astaxie/beego"
)

type AccountController struct {
	beego.Controller
}

func (this *AccountController) Get() {
	uname := GetAccountCookie(this.Ctx)
	beego.Debug("uname checked in accountmgt:", CheckAccount(this.Ctx))
	if this.isVisitValid() {
		this.Data["IsLogined"] = true
		watchedShop, existed1 := GetWatchedShop(uname)
		if existed1 {
			this.Data["watchedShopId"] = watchedShop.ShopId
			this.Data["watchedShopName"] = watchedShop.ShopName
			this.Data["watchedShopType"] = watchedShop.ShopType
		}

		watcherShop, existed2 := GetWatcherShop(uname)
		if existed2 {
			this.Data["watcherShopId"] = watcherShop.ShopId
			this.Data["watcherShopName"] = watcherShop.ShopName
			this.Data["watcherShopType"] = watcherShop.ShopType
		}

		//需要显示反馈表格
		feedbacks, _ := GetReadableFeedbackList(uname)
		if len(feedbacks) > 0 {
			//show me the table
			beego.Debug("more feedbacks coming!", feedbacks)
			this.Data["feedbacks"] = feedbacks
		}

		if existed1 && existed2 {
			this.Data["IsShopSetted"] = true
		}

		this.TplNames = "account.html"
	} else {
		beego.Debug("session bad,redirecting home")
		this.Redirect("/", 302)
		//this.TplNames = "index.html"
		return
	}
}

func (this *AccountController) Post() {
	uname := GetAccountCookie(this.Ctx)
	if this.isVisitValid() {
		//目标设置部分
		watchedShop, _ := GetWatchedShop(uname)
		targetShopId := this.Input().Get("targetshopid")
		targetShopName := this.Input().Get("targetshopname")
		targetShopType := this.Input().Get("targetshoptype")
		beego.Debug("old shopid:", watchedShop.ShopId)
		beego.Debug("old shopname:", watchedShop.ShopName)
		beego.Debug("old shoptype:", watchedShop.ShopType)
		beego.Debug("old watcher:", uname)
		beego.Debug("got shopid:", targetShopId)
		beego.Debug("got shopname:", targetShopName)
		beego.Debug("got shoptype:", targetShopType)

		if watchedShop.ShopId == "" && watchedShop.ShopName == "" {
			//对输入进行查询校验
			if IsShopExisted(targetShopId, targetShopName) {
				AddNewWatchedShop(targetShopId, targetShopName, targetShopType, uname)
			} else {
				beego.Debug("input shop not exist in taobao")
			}

		} else if watchedShop.ShopId != targetShopId || watchedShop.ShopName != targetShopName || watchedShop.ShopType != targetShopType {
			//有变化，将旧的设置无效，同时插入新数据
			if IsShopWatchedBefore(targetShopId, targetShopName, targetShopType, uname) {
				DeactiveWatchedShop(*watchedShop)
				beego.Debug("old records deactivated!")
				ActiveWatchedShop(targetShopId, targetShopName, targetShopType, uname)
			} else {
				DeactiveWatchedShop(*watchedShop)
				beego.Debug("old records deactivated 2!")
				AddNewWatchedShop(targetShopId, targetShopName, targetShopType, uname)
			}
		}
		//自己店铺
		watcherShop, _ := GetWatcherShop(uname)
		myShopId := this.Input().Get("myshopid")
		myShopName := this.Input().Get("myshopname")
		myShopType := this.Input().Get("myshoptype")
		if watcherShop.ShopId == "" && watcherShop.ShopName == "" {
			AddNewWatcherShop(myShopId, myShopName, myShopType, uname)
		} else if watcherShop.ShopId != myShopId || watcherShop.ShopName != myShopName || watcherShop.ShopType != myShopType {
			if IsShopWatcherBefore(myShopId, myShopName, myShopType, uname) {
				DeactiveWatcherShop(*watcherShop)
				beego.Debug("old records deactivated!")
				ActiveWatcherShop(myShopId, myShopName, myShopType, uname)
			} else {
				DeactiveWatcherShop(*watcherShop)
				beego.Debug("old records deactivated 2!")
				AddNewWatcherShop(myShopId, myShopName, myShopType, uname)
			}
		}

		//提交反馈表格
		title := this.Input().Get("feedbacktitle")
		content := this.Input().Get("feedbackcontent")
		if title != "" && content != "" {
			AddNewFeedback(title, content, uname)
		}

		//修改密码
		oldpwd := this.Input().Get("oldpwd")
		newpwd := this.Input().Get("newpwd")
		if oldpwd != "" && newpwd != "" {
			if CheckUserPassword(uname, oldpwd) {
				if ChangeUserPassword(uname, newpwd) {
					beego.Debug("change pwd success!")
				} else {
					this.Data["IsOldPwdError"] = true
					beego.Debug("change pwd failed!")
				}
			}
		}
		this.Redirect("/account", 302)
		//this.TplNames = "account.html"
	} else {
		beego.Debug("bad post in account, redirecting home")
		this.Redirect("/", 302)
		//this.TplNames = "index.html"
		return
	}
}

func (this *AccountController) isVisitValid() bool {
	uname := GetAccountCookie(this.Ctx)
	s := this.GetSession(uname)
	if CheckAccount(this.Ctx) && s != nil {
		return true
	} else {
		return false
	}
}
