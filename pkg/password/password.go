package password

import (
	"eshort/pkg/easylogger"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	easylogger.LogError(err, "hash生成出错")
	return string(bytes)
}

func CheckHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	easylogger.LogError(err, "检查hash出错")
	return err == nil
}

func IsHash(password string) bool {
	//使用bcrypt加密后的密码为60长度
	return len(password) == 60
}
