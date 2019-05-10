package controllers

import (
	"github.com/astaxie/beego"

	"github.com/CoderReaper/goweb/module"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "index.tpl"
	//c.Ctx.WriteString("heloword")
}
func (c *LoginController) Post() {
	var email = c.GetString("email")
	var pswd = c.GetString("password")
	module.Login(email, pswd)
	c.Ctx.WriteString("wait...")

}
