package controllers

import (
	"93ztc/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

func AddNewWatchedShop(shopid string, shopname string, shoptype string, uname string) {
	if shopid == "" || shopname == "" || shoptype == "" || uname == "" {
		return
	}
	o := orm.NewOrm()
	watchedShop := new(models.WatchedShop)
	watchedShop.ShopId = shopid
	watchedShop.ShopName = shopname
	watchedShop.ShopType = shoptype
	watchedShop.Watcher = uname
	watchedShop.Active = 1
	watchedShop.CreateTime = time.Now().Unix()
	o.Insert(watchedShop)
	AddNewScheduledShop(shopid, shopname)
}

func AddNewWatcherShop(shopid string, shopname string, shoptype string, uname string) {
	if shopid == "" || shopname == "" || shoptype == "" || uname == "" {
		return
	}
	o := orm.NewOrm()
	watcherShop := new(models.WatcherShop)
	watcherShop.ShopId = shopid
	watcherShop.ShopName = shopname
	watcherShop.ShopType = shoptype
	watcherShop.UserName = uname
	watcherShop.Active = 1
	watcherShop.CreateTime = time.Now().Unix()
	o.Insert(watcherShop)
	AddNewScheduledShop(shopid, shopname)
}

func GetWatchedShop(uname string) (*models.WatchedShop, bool) {
	o := orm.NewOrm()
	watchedShop := new(models.WatchedShop)
	var shops []*models.WatchedShop
	num1, _ := o.Raw("select * from watched_shop where active=1 and watcher=?", uname).QueryRows(&shops)
	beego.Debug("found user watched shop:", num1)
	if num1 > 0 {
		beego.Debug("Got shops:", shops)
		watchedShop = shops[0] //正常是只有一个有效的监控店铺
		return watchedShop, true
	} else {
		return watchedShop, false
	}
	return watchedShop, false
}

func IsShopWatchedBefore(shopid, shopname, shoptype, watcher string) bool {
	o := orm.NewOrm()
	var shops []*models.WatchedShop
	num1, _ := o.Raw("select * from watched_shop where shop_id=? and shop_name=? and shop_type=? and watcher=?", shopid, shopname, shoptype, watcher).QueryRows(&shops)
	beego.Debug("finding old shop watched!")
	if num1 > 0 {
		return true
	} else {
		return false
	}
}

func DeactiveWatchedShop(watcher models.WatchedShop) bool {
	o := orm.NewOrm()
	watcher.Active = 0
	o.Update(&watcher)
	return true
}

func ActiveWatchedShop(shopid, shopname, shoptype, watcher string) bool {
	o := orm.NewOrm()
	var maps []orm.Params
	num1, _ := o.Raw("update watched_shop set active = 1 where shop_id=? and shop_name=? and shop_type=? and watcher =?", shopid, shopname, shoptype, watcher).Values(&maps)
	if num1 > 0 {
		return true
	}
	return false
}

func GetWatcherShop(uname string) (*models.WatcherShop, bool) {
	o := orm.NewOrm()
	watcherShop := new(models.WatcherShop)
	var shops []*models.WatcherShop
	num1, _ := o.Raw("select * from watcher_shop where active=1 and user_name=?", uname).QueryRows(&shops)
	beego.Debug("found user watcher shop:", num1)
	if num1 > 0 {
		beego.Debug("Got watcher shops:", shops)
		watcherShop = shops[0] //正常是只有一个有效的监控店铺
		return watcherShop, true
	} else {
		return watcherShop, false
	}
	return watcherShop, false
}

func IsShopWatcherBefore(shopid, shopname, shoptype, uname string) bool {
	o := orm.NewOrm()
	var shops []*models.WatcherShop
	num1, _ := o.Raw("select * from watcher_shop where shop_id=? and shop_name=? and shop_type=? and user_name=?", shopid, shopname, shoptype, uname).QueryRows(&shops)
	beego.Debug("finding old shop watcher~~~!")
	if num1 > 0 {
		return true
	} else {
		return false
	}
}

func DeactiveWatcherShop(watcher models.WatcherShop) bool {
	o := orm.NewOrm()
	watcher.Active = 0
	o.Update(&watcher)
	return true
}

func ActiveWatcherShop(shopid, shopname, shoptype, uname string) bool {
	o := orm.NewOrm()
	var maps []orm.Params
	num1, _ := o.Raw("update watcher_shop set active = 1 where shop_id=? and shop_name=? and shop_type=? and user_name =?", shopid, shopname, shoptype, uname).Values(&maps)
	if num1 > 0 {
		return true
	}
	return false
}

func GetScheduledShop(crawlerid string) ([]models.ScheduledShop, bool) {
	o := orm.NewOrm()
	var scheduledShops []models.ScheduledShop
	num, err := o.Raw("select * from scheduled_shop where crawler_id=?", crawlerid).QueryRows(&scheduledShops)
	if err == nil {
		beego.Debug("scheduled shops nums:", num)
		return scheduledShops, true
	} else {
		return nil, false
	}
}

//由于trigger不太好用，改为程序控制插入逻辑。。。
func IsScheduledShopExisted(shopid, shopname string) bool {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("select * from scheduled_shop where shop_id=? and shop_name=?", shopid, shopname).Values(&maps)
	if err == nil {
		beego.Debug("found shop:#", num, shopid, shopname)
		if num == 0 {
			return false //not existed
		} else {
			return true
		}

	} else {
		return false
	}
}

func AddNewScheduledShop(shopid, shopname string) bool {
	if IsScheduledShopExisted(shopid, shopname) {
		beego.Debug("scheduled shop existed!", shopid, shopname)
		return false
	}
	o := orm.NewOrm()
	scheduledShop := new(models.ScheduledShop)
	scheduledShop.ShopId = shopid
	scheduledShop.ShopName = shopname
	scheduledShop.CrawlerId = `1`
	scheduledShop.CreateTime = time.Now().Unix()
	o.Insert(scheduledShop)
	beego.Debug("scheduled shop inserted!", shopid, shopname)
	return true
}
