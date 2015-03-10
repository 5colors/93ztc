package routers

import (
	"93ztc/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/service", &controllers.ServiceController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/account", &controllers.AccountController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/dailyreport", &controllers.DailyReportController{})
	beego.Router("/products", &controllers.ProductsController{})
	beego.Router("/faq", &controllers.FaqController{})
	beego.Router("/aboutus", &controllers.AboutusController{})
	beego.Router("/terms", &controllers.TermsController{})
	beego.Router("/moreus", &controllers.MoreUsController{})
	beego.Router("/jobs", &controllers.JobsController{})
	beego.Router("/order", &controllers.OrderController{})
	beego.Router("/scheduledshop", &controllers.ScheduledShopController{})
	beego.Router("/articles", &controllers.ArticlesController{})
	beego.Router("/roadmap", &controllers.RoadmapController{})
	beego.Router("/maillist", &controllers.MailListController{})
}
