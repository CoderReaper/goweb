package controllers

import (
	"github.com/astaxie/beego"
)

type HsController struct {
	beego.Controller
}

func (c *HsController) Get() {
	c.Data["MyAddress"] = "mail.126.com"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
