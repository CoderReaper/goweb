package main

import (
	_ "hello/routers"
	"github.com/astaxie/beego"
	"hello/controllers"
)

type AddController struct {
	    beego.Controller	    
    }

func (this* AddController) Get(){
	  this.Ctx.WriteString("hello world")

}
func main() {
	beego.Router("/helo", &AddController{})
	beego.Router("/Hs",&controllers.HsController{})
	beego.Run()

}

