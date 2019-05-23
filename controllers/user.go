package controllers

import (
	"fmt"

	"github.com/CoderReaper/goweb/util"
	"github.com/astaxie/beego"
)

//UserController .
type UserController struct {
	beego.Controller
}

//Get .
func (c *UserController) Get() {
	c.TplName = "index.tpl"
	name := c.Ctx.GetCookie("name")
	email := c.Ctx.GetCookie("email")
	token := c.Ctx.GetCookie("token")
	//验证token
	b, err := util.CheckToken(token, email)
	if err != nil {
		c.Ctx.WriteString("get user check token fail ")
		return
	}
	//失败
	if !b {
		c.Redirect("/login", 301)
		return
	}
	str := fmt.Sprintf("%s", name)
	c.Ctx.WriteString("hello " + str)
}
