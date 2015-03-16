package controllers

import (
	"github.com/astaxie/beego/context"
)

//获取cookie中记录的用户名进行会话检查
func CheckAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value
	//对比之后返回true
	if uname == "" {
		return false
	} else {
		return true
	}
}

func GetAccountCookie(ctx *context.Context) string {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return ""
	} else {
		return ck.Value
	}

}
