package controllers

import (
	"93ztc/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

func AddNewUser(uname string, pwd string) {
	o := orm.NewOrm()
	user := new(models.User)
	user.Name = uname
	user.Password = pwd
	user.LevelActiveTime = time.Now().Unix()
	o.Insert(user)
}

func CheckUserExist(uname string) bool {
	o := orm.NewOrm()
	var maps []orm.Params
	num1, _ := o.Raw("select name from user where name=?", uname).Values(&maps)
	beego.Debug("checkuser exist:", num1)
	if num1 > 0 {
		return true
	} else {
		return false
	}
}

func CheckUserPassword(uname string, pwd string) bool {
	o := orm.NewOrm()
	var maps []orm.Params
	num1, _ := o.Raw("select name, password from user where name=? and password=?", uname, pwd).Values(&maps)
	beego.Debug("checkuser password:", num1)
	if num1 > 0 {
		return true
	} else {
		return false
	}
}

func ChangeUserPassword(uname string, newpwd string) bool {
	o := orm.NewOrm()
	user := new(models.User)
	var users []*models.User
	num1, _ := o.Raw("select * from user where name=?", uname).QueryRows(&users)
	if num1 > 0 {
		beego.Debug("Trying to change password=>", num1)
		user = users[0] //应该只有一个结果
		user.Password = newpwd
		beego.Debug("new Password:", user.Password)
		if num, err := o.Update(user); err == nil {
			beego.Debug("change pwd should be ok! ", num)
			return true
		} else {
			beego.Debug(err.Error())
			return false
		}
	}
	return false
}
