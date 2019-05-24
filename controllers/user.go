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
		rsp := HTTPRetMessage{
			Ret:    "fail",
			Reason: err.Error(),
			Data:   "please try later",
		}
		c.Data["json"] = &rsp
		c.ServeJSON()
		return
	}
	//失败
	if !b {
		c.Redirect("/login", 301)
		return
	}
	str := fmt.Sprintf("%s", name)
	rsp := HTTPRetMessage{
		Ret:    "ok",
		Reason: "",
		Data:   "hello " + str,
	}
	c.Data["json"] = &rsp
	c.ServeJSON()
}
