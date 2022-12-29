package admin

import (
	"eshort/pkg/easylogger"
	"eshort/pkg/model"
	"eshort/pkg/password"
	"eshort/pkg/types"
)

func All() ([]Admin, error) {
	var users []Admin
	err := model.DB.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func GetUserByPhone(phone string) (Admin, error) {
	var user Admin
	err := model.DB.Where("phone=?", phone).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func Get(idstr string) (Admin, error) {
	var user Admin
	id := types.StringToUint64(idstr)
	err := model.DB.First(&user, id).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetByEmail(email string) (Admin, error) {
	var user Admin
	err := model.DB.Where("email=?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetByPhone(phone string) (Admin, error) {
	var user Admin
	err := model.DB.Where("phone=?", phone).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (author *Admin) ComparePassword(str string) bool {
	return password.CheckHash(str, author.Password)
}

func (author *Admin) Create() (err error) {
	err = model.DB.Create(&author).Error
	if err != nil {
		easylogger.LogError(err, "添加用户失败")
		return err
	}
	return nil
}
