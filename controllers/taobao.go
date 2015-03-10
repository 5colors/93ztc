package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strings"
)

func IsShopExisted(shopid string, shopname string) bool {
	url := "http://s.taobao.com/search?app=shopsearch&q=" + shopid
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(body))
	content := string(body)

	if response.StatusCode == 200 {
		//简单做一个查询验证。。
		beego.Debug("Shop found:", strings.Contains(content, shopid) && strings.Contains(content, `list-item`))
		return strings.Contains(content, shopid) && strings.Contains(content, `list-item`)

	} else {
		return false
	}

}
