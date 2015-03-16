package main

import (
	"93ztc/models"
	_ "93ztc/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/93ztc?charset=utf8")
	orm.RegisterModel(new(models.User), new(models.WatchedShop), new(models.WatcherShop), new(models.Feedback),
		new(models.ScheduledShop), new(models.MailList), new(models.Captcha))
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", true, true)
	beego.SetStaticPath("/uploadcaptcha", "/tmp/captcha")
	beego.SessionOn = true
	beego.Run()
}
