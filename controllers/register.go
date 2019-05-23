package controllers

import (
	"github.com/CoderReaper/goweb/module"
	"github.com/CoderReaper/goweb/util"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "index.tpl"
	//c.Ctx.WriteString("heloword")
}

//Post 注册
func (c *RegisterController) Post() {
	//检测email
	var email = c.GetString("email")

	if !util.CheckEmail(email) || !module.EmailIsRegister(email) {
		//c.Ctx.WriteString("'result':'fake','reason':'invalid email'}")
		c.Data["msg"] = "invalid email"
		c.Data["result"] = "fail"
		c.ServeJSON()
		return
	}

	//检测username
	var name = c.GetString("name")
	if !util.CheckUserName(name) || !module.UserNameIsRegister(name) {
		c.Data["msg"] = "invalid name"
		c.Data["result"] = "fail"
		c.ServeJSON()
		return
	}
	//检测password
	var pswd = c.GetString("password")
	if !util.CheckPassWord(pswd) {
		c.Data["msg"] = "invalid password"
		c.Data["result"] = "fail"
		c.ServeJSON()
		return
	}

	//可以注册
	password, err := bcrypt.GenerateFromPassword([]byte(pswd+name), 4)
	if err != nil {
		c.Data["msg"] = "unknown reason"
		c.Data["result"] = "fail"
		c.ServeJSON()
		return
	}
	//db操作
	println(password)
	encodePassWord := string(password[:])
	err = module.Register(email, name, encodePassWord)
	if err != nil {
		c.Data["msg"] = "unknown reason"
		c.Data["result"] = "fail"
		c.ServeJSON()
		return
	}

	//登录成功
	c.Redirect("/user", 301)
}
