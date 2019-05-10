package controllers

import (
	"github.com/astaxie/beego"

	"github.com/CoderReaper/goweb/module"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "index.tpl"
	//c.Ctx.WriteString("heloword")
}
func (c *RegisterController) Post() {
	var email = c.GetString("email")
	var pswd = c.GetString("password")
	module.Login(email, pswd)
	c.Ctx.WriteString("wait...")

}
