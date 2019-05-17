package util

import (
	"regexp"
	"sync"
)

var emailReg *regexp.Regexp
var userNameReg *regexp.Regexp
var one sync.Once

const (
	maxNameLen int = 20
	minNameLen int = 6

	maxPassWordLen int = 20
	minPassWordLen int = 6
)

func init() {
	one.Do(func() {
		emailReg = regexp.MustCompile(`[\w]+@[\w]+\.com`)
		userNameReg = regexp.MustCompile(`[\w]+`)
	})
}

//CheckEmail .
func CheckEmail(email string) bool {
	if emailReg.FindString(email) != email {
		return false
	}
	return true
}

//CheckUserName .
func CheckUserName(name string) bool {
	if len(name) > maxNameLen || len(name) <= minNameLen {
		return false
	}
	if userNameReg.FindString(name) != name {
		return false
	}
	return true
}

//CheckPassWord .
func CheckPassWord(passwd string) bool {
	if len(passwd) > maxPassWordLen || len(passwd) <= minPassWordLen {
		return false
	}
	return true
}
