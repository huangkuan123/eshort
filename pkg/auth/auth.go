package auth

import (
	"errors"
	"eshort/app/models/admin"
	"eshort/pkg/session"
	"eshort/pkg/types"
	"gorm.io/gorm"
)

var session_name string = "uid"

// 认证管理包
// User 获取当前登录用户
func User() admin.Admin {
	uid := GetUID()
	if len(uid) > 0 {
		userData, err := admin.Get(uid)
		if err == nil {
			return userData
		}
	}
	return admin.Admin{}
}

func Login(userData admin.Admin) {
	session.Put(session_name, userData.GetStringID())
}

func Logout() {
	session.Forget(session_name)
}

func Check() bool {
	return len(GetUID()) > 0
}

func Attempt(phone string, password string) error {
	userData, err := admin.GetByPhone(phone)
	//fmt.Println("userData", userData)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("账号不存在或密码错误")
		}
		return errors.New("内部有误")
	}
	if !userData.ComparePassword(password) {
		return errors.New("账号不存在或密码错误")
	}
	//fmt.Println("开始写session")
	session.Put(session_name, userData.GetStringID())
	return nil
}

func GetUID() string {
	_uid := session.Get(session_name)
	uid, ok := _uid.(string)
	if ok && len(uid) > 0 {
		return uid
	}
	return ""
}

func GetUIDUint64() uint64 {
	return types.StringToUint64(GetUID())
}
