package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type CaptchaController struct {
	beego.Controller
}

func (this *CaptchaController) Get() {
	cid := this.Input().Get("crawlerid")
	itid := this.Input().Get("itemid")
	beego.Debug("cid:", cid)
	beego.Debug("itid:", itid)
	if cid == "" || itid == "" { //input reuslts...
		itemid, filename, crawlerid := GetUnsolvedCaptcha()
		this.Data["itemid"] = itemid
		this.Data["filename"] = filename
		this.Data["crawlerid"] = crawlerid
		this.TplNames = "captcha.html"
	} else { //crawler fetching
		result := GetSolvedCaptcha(itid, cid)
		beego.Debug("Found result:", result)
		this.Ctx.WriteString(result)
	}

}
func (this *CaptchaController) InputResult() {
	cid := this.Input().Get("crawlerid")
	itid := this.Input().Get("itemid")
	result := this.Input().Get("result")
	beego.Debug("cid:", cid)
	beego.Debug("itid:", itid)
	beego.Debug("input result:", result)
	SaveCaptchaResult(itid, cid, result)
	this.Redirect("/captcha", 302)
}

func (this *CaptchaController) Post() {
	crawlerId := this.Input().Get("crawlerid")
	beego.Debug("crawler id:", crawlerId)
	f, h, err := this.GetFile("captchafile")
	if err != nil {
		beego.Debug("get file err!", err)
	}

	filename := h.Filename
	beego.Debug("filename:", filename)
	path := `/tmp/captcha/` + filename //static filepath registered in main.go

	ss := strings.Split(filename, ".")
	beego.Debug(ss[0]) //itemid
	beego.Debug("save path:", path)
	f.Close()
	err = this.SaveToFile("captchafile", path)
	if err != nil {
		beego.Debug("save err:", err)
	}

	//insert into db
	AddNewCaptcha(ss[0], filename, crawlerId)
	this.Ctx.WriteString("upload end..")
}
