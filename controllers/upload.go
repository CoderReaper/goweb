package controllers

import (
	"github.com/CoderReaper/goweb/util"
	"github.com/astaxie/beego"
)

//UploadController .
type UploadController struct {
	beego.Controller
}

//Get .
func (c *UploadController) Get() {
	c.TplName = "upload.tpl"
}

//Post .
func (c *UploadController) Post() {
	c.TplName = "upload.tpl"
	token := c.Ctx.GetCookie("token")
	email := c.Ctx.GetCookie("email")
	b, err := util.CheckToken(token, email)
	if err != nil {
		c.Ctx.WriteString("user upload file fail token check fail ")
		return
	}
	if !b {
		c.Ctx.WriteString("user upload file fail token check fail2 ")
		return
	}

	f, fh, err := c.GetFile("upfile")
	if err != nil {
		beego.Warning("getfile err ", err)
	}
	defer f.Close()
	c.SaveToFile("upfile", "static/upload/"+fh.Filename)
	//email := c.Ctx.GetCookie("email")
	//name := c.Ctx.GetCookie("name")
}
