package controllers

import (
	"github.com/CoderReaper/goweb/module"
	"github.com/CoderReaper/goweb/util"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

//HTTPRetMessage call back
type HTTPRetMessage struct {
	Ret    string `json:"ret"`
	Reason string `json:"reason"`
	Data   string `json:"data"`
}

//RegisterController .
type RegisterController struct {
	beego.Controller
}

//Get .
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
		rsp := HTTPRetMessage{
			Ret:    "fail",
			Reason: "invalid email",
			Data:   "check your email",
		}
		c.Data["json"] = &rsp
		c.ServeJSON()
		return
	}

	//检测username
	var name = c.GetString("name")
	if !util.CheckUserName(name) || !module.UserNameIsRegister(name) {
		rsp := HTTPRetMessage{
			Ret:    "fail",
			Reason: "invalid name",
			Data:   "check your name",
		}
		c.Data["json"] = &rsp
		c.ServeJSON()
		return
	}
	//检测password
	var pswd = c.GetString("password")
	if !util.CheckPassWord(pswd) {
		rsp := HTTPRetMessage{
			Ret:    "fail",
			Reason: "invalid password",
			Data:   "check your password",
		}
		c.Data["json"] = &rsp
		c.ServeJSON()
		return
	}

	//可以注册
	password, err := bcrypt.GenerateFromPassword([]byte(pswd+name), 4)
	if err != nil {
		rsp := HTTPRetMessage{
			Ret:    "fail",
			Reason: "unknown reason",
			Data:   "please try later ",
		}
		c.Data["json"] = &rsp
		c.ServeJSON()
		return
	}
	//db操作
	println(password)
	encodePassWord := string(password[:])
	err = module.Register(email, name, encodePassWord)
	if err != nil {
		rsp := HTTPRetMessage{
			Ret:    "fail",
			Reason: "unknown reason",
			Data:   "please try later ",
		}
		c.Data["json"] = &rsp
		c.ServeJSON()
		return
	}

	//登录成功
	c.Redirect("/user", 301)
}
