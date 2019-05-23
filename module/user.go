package module

import (
	"encoding/json"
	"time"

	myredis "github.com/CoderReaper/goweb/third/redis"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

//UserData .
type UserData struct {
	Email    string `json:"email"`
	Uname    string `json:"uname"`
	Regtime  int32  `json:"regtime"`
	Password string `json:"password"`
}

//UserDBModle .
type UserDBModle struct {
}

//Login 登录
func Login(email string, name string, pass string) (bool, error) {
	ret, err := myredis.Do("HGET", myredis.TUser, email)
	if err != nil {
		beego.Warning("module user Login do fail %s", err.Error())
		return false, err
	}
	var userdata UserData
	err = json.Unmarshal(ret.([]byte), &userdata)
	if err != nil {
		beego.Critical("module user Login unmarshal fail %s", err.Error())
		return false, err
	}
	//验证密码
	err = bcrypt.CompareHashAndPassword([]byte(userdata.Password), []byte(pass+name))
	if err != nil {
		beego.Warning("module user Login password fail")
		return false, nil
	}
	//succ
	beego.Info("module user login succ %v", ret)
	return true, nil
}

//Register 注册
func Register(email string, name string, passwd string) error {
	//先加锁
	lock := myredis.RedisSyncLock{
		Lockkey: "key_user_register",
		Expr:    10,
	}
	//连接
	conn := myredis.GetConn()
	//更新t_email表
	emailInfo := []interface{}{myredis.TMail, email, name}
	_, err := conn.Do("HSET", emailInfo...)
	//更新t_uname表
	userInfo := []interface{}{myredis.TUserName, name, email}
	_, err = conn.Do("HSET", userInfo...)
	//释放锁
	lock.Unlock(conn)
	//更新t_user表
	userdata := UserData{
		Email:    email,
		Uname:    name,
		Password: passwd,
		Regtime:  int32(time.Now().Unix()),
	}
	redisdata, err := json.Marshal(userdata)
	if err != nil {
		beego.Emergency("user mashal data fail %v", userdata)
		return err
	}
	_, err = conn.Do("HSET", myredis.TUser, email, redisdata)
	if err != nil {
		beego.Warning("module user register do fail %s", err.Error())
	}
	defer func() {
		myredis.PutBack(conn)
	}()
	beego.Info("user %s email %s register succ", name, email)
	return err
}

//EmailIsRegister 检测email
func EmailIsRegister(email string) bool {
	ret, err := myredis.Do(myredis.HGET, myredis.TMail, email)
	if err != nil {
		beego.Emergency("module EmailIsRegister fatal %s ", err.Error())
		return false
	}
	return ret == nil
}

//UserNameIsRegister 检测用户名
func UserNameIsRegister(uname string) bool {
	ret, err := myredis.Do(myredis.HGET, myredis.TUserName, uname)
	if err != nil {
		beego.Emergency("module UserNameIsRegister fatal %s ", err.Error())
		return false
	}
	return ret == nil
}

//GetEmailByName 根据用户名获取邮箱
func GetEmailByName(name string) (string, error) {
	ret, err := myredis.Do(myredis.HGET, myredis.TUserName, name)
	if err != nil {
		beego.Emergency("module GetEmailByName fatal %s ", err.Error())
		return "", err
	}
	return string(ret.([]byte)), nil
}

//GetNameByEmail 根据邮箱获取用户名
func GetNameByEmail(email string) (string, error) {
	ret, err := myredis.Do(myredis.HGET, myredis.TMail, email)
	if err != nil {
		beego.Emergency("module GetNameByEmail fatal %s ", err.Error())
		return "", err
	}
	return string(ret.([]byte)), nil
}
