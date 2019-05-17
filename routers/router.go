package routers

import (
	"github.com/CoderReaper/goweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.RootController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogOutController{})
	beego.Router("/register", &controllers.RegisterController{})
}
