package controllers

import (
	"github.com/astaxie/beego"
)

type ScheduledShopController struct {
	beego.Controller
}

func (this *ScheduledShopController) Get() {
	//this.Data["json"] = map[string]string{"ObjectId": "hello"}
	crawlerid := this.Input().Get("crawlerid")
	if crawlerid != "" {
		shops, ok := GetScheduledShop(crawlerid)
		if ok && shops != nil {
			var str string
			str = `[`
			for i := 0; i < len(shops); i++ {
				str = str + `{"shopid":"` + shops[i].ShopId + `","shopname":"` + shops[i].ShopName + `"},`
			}
			str = str + `]`
			this.Data["json"] = map[string]interface{}{"crawlerid": crawlerid, "shops": str}
			this.ServeJson()
			beego.Debug("scheduled shops got...")
		} else {
			beego.Debug("Error getting scheduled shops...")
			this.Ctx.WriteString("error!")
		}

	} else {
		this.Ctx.WriteString("need id!")
		beego.Debug("empty crawler id...")
	}

}
