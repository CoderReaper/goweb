package main

import (
	_ "github.com/CoderReaper/goweb/routers"
	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogger(logs.AdapterFile, "{\"filename\": \"/beego.log\"}")
	beego.SetLevel(logs.LevelDebug)
	beego.SetLogFuncCall(true)
	beego.Run()
}
