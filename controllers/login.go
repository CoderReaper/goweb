package controllers

import (
	"github.com/CoderReaper/goweb/util"
	"github.com/astaxie/beego"

	"github.com/CoderReaper/goweb/module"
)

//LoginController
type LoginController struct {
	beego.Controller
}

//Get
func (c *LoginController) Get() {
	c.TplName = "index.tpl"
	c.Data["login"] = "true"
	c.Data["msg"] = "unknown reason"
	c.Data["result"] = "fail"
	c.ServeJSON()
	//c.Ctx.WriteString("heloword")
}

//Post post
func (c *LoginController) Post() {
	var usstr = c.GetString("uestr")
	var password = c.GetString("password")
	var email = ""
	var name = ""
	var err error
	//是否是email
	if util.CheckEmail(usstr) {
		email = usstr
		name, err = module.GetNameByEmail(email)
	} else {
		name = usstr
		email, err = module.GetEmailByName(name)
	}
	if err != nil {
		beego.Emergency("user login fail %s", err.Error())
		c.Ctx.WriteString("login fail ")
		return
	}
	ret, err := module.Login(email, name, password)
	if err != nil || !ret {
		beego.Info("login fail email %s", email)
		c.Ctx.WriteString("login fail ")
		return
	}
	//生成token
	token, _ := util.CreateToken(email)
	//写cookie
	c.Ctx.SetCookie("email", email)
	c.Ctx.SetCookie("name", name)
	c.Ctx.SetCookie("token", token, 60*60*24*7)
	//b, _ := util.CheckToken(token, email)
	//println(b)
	c.Redirect("/user", 301)
}
