package main

import (
	_ "github.com/CoderReaper/goweb/routers"
	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

func main() {
	//开启日志
	beego.SetLogger(logs.AdapterFile, "{\"filename\": \"/home/gcloud/log/goweb/beego.log\"}")
	beego.SetLevel(logs.LevelDebug)
	beego.SetLogFuncCall(true)
	//session
	beego.BConfig.WebConfig.Session.SessionOn = true
	//beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"

	beego.Run()
}
