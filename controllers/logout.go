package controllers

import (
	"github.com/astaxie/beego"
)

type LogOutController struct {
	beego.Controller
}

//Get .
func (c *LogOutController) Get() {
	c.TplName = "index.tpl"
}

//Post .
func (c *LogOutController) Post() {
	c.TplName = "index.tpl"
	logout := c.GetString("logout")
	token := c.Ctx.GetCookie("token")
	if logout != "" && token != "" {
		c.Ctx.SetCookie("token", "")
	}
	c.Redirect("/login", 301)
}
