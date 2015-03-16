package controllers

import (
	"93ztc/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func AddNewCaptcha(itemid string, filename string, crawlerid string) {
	if itemid == "" || filename == "" || crawlerid == "" {
		return
	}

	o := orm.NewOrm()
	//remove old one if existed based on itemid and crawlerid
	num, err := o.QueryTable("captcha").Filter("item_id", itemid).Filter("crawler_id", crawlerid).Delete()
	beego.Debug("deleted: %s,%s,%s", num, err, itemid)
	captcha := new(models.Captcha)
	captcha.ItemId = itemid
	captcha.CaptchaFile = filename
	captcha.CrawlerId = crawlerid
	o.Insert(captcha)
}

func GetUnsolvedCaptcha() (string, string, string) { //itemid,filename,crawlerid
	o := orm.NewOrm()
	captcha := new(models.Captcha)
	var captchas []*models.Captcha
	num, _ := o.Raw("select * from captcha where result ='' limit 1").QueryRows(&captchas)
	if num > 0 {
		captcha = captchas[0]
		beego.Debug("itemid:%s,filename:%s,crawlerid:%s", captcha.ItemId, captcha.CaptchaFile, captcha.CrawlerId)
		return captcha.ItemId, captcha.CaptchaFile, captcha.CrawlerId
	} else {
		return "", "", ""
	}
}

func GetSolvedCaptcha(itemid string, crawlerid string) string {
	o := orm.NewOrm()
	captcha := new(models.Captcha)
	var captchas []*models.Captcha
	num, _ := o.Raw("select * from captcha where item_id=? and crawler_id=? ",
		itemid, crawlerid).QueryRows(&captchas)
	if num > 0 {
		captcha = captchas[0]
		beego.Debug("itemid:%s,result:%s,crawlerid:%s", captcha.ItemId, captcha.Result, captcha.CrawlerId)
		return captcha.Result
	} else {
		return ""
	}
}

func SaveCaptchaResult(itemid string, crawlerid string, result string) bool {
	o := orm.NewOrm()
	res, err := o.Raw("update captcha set result =? where item_id =? and crawler_id=? ",
		result, itemid, crawlerid).Exec()
	if err == nil {
		beego.Debug("updated result:", res)
		return true
	} else {
		return false
	}

}
