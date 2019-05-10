package controllers

import (
	"github.com/astaxie/beego"
)

type LogOutController struct {
	beego.Controller
}

func (c *LogOutController) Get() {
	//	c.Data["Website"] = "beego.me"
	//	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
