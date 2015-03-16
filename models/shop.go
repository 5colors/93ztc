package models

type WatchedShop struct {
	Id         int64
	ShopId     string
	ShopName   string
	ShopLink   string `orm:"default('')"`
	ShopType   string
	Watcher    string //==wahtershop.username
	Active     int    `orm:"default(1)"`
	CreateTime int64
}

type WatcherShop struct {
	Id         int64
	ShopId     string
	ShopName   string
	ShopLink   string `orm:"default('')"`
	ShopType   string
	UserName   string
	Active     int `orm:"default(1)"`
	CreateTime int64
}

type ScheduledShop struct {
	Id         int64
	ShopId     string
	ShopName   string
	CrawlerId  string
	CreateTime int64
}
