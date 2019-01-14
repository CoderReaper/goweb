package main

import (
	"hello/controllers"
	_ "hello/routers"

	"github.com/astaxie/beego"
)

type AddController struct {
	beego.Controller
}

func (this *AddController) Get() {
	this.Ctx.WriteString("hello world")

}
func main() {
	beego.Router("/helo", &AddController{})
	beego.Router("/Hs", &controllers.HsController{})
	beego.Run()
}
