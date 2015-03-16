package models

type User struct {
	Id              int64
	Name            string `orm:"Index"`
	Password        string
	Email           string `orm:"null"`
	Comment         string `orm:"null"` //注册表单中的信息来源
	Refer           string `orm:"null"` //推荐人
	Phone           string `orm:"null"`
	Active          int    `orm:"default(0)"` //是否激活？付过款了？
	Level           int    `orm:"default(0)"` //0, 总监级，1，专家级。。。类同淘宝等级
	LevelActiveTime int64  `orm:"auto_now_add;type(date)"`
}
