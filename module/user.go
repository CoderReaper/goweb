package module

import (
	myredis "github.com/CoderReaper/goweb/third/redis"
	"github.com/astaxie/beego"
)

type UserDBModle struct {
}

//Login 登录
func Login(email string, pass string) bool {
	ret, err := myredis.Do("HMGET", "Users", email, "email", "password")
	if err != nil {
		beego.Warning("module user Login do fail %s", err.Error())
		return false
	}
	beego.Info("module user login succ %v", ret)
	return true
}

//Register 注册
func Register(email string, name string, passwd string) error {
	userinfo := []interface{}{email, "email", email, "name", name, "password", passwd}
	_, err := myredis.Do("HMSET", userinfo...)
	if err != nil {
		beego.Warning("module user register do fail %s", err.Error())
	}
	return err
}
