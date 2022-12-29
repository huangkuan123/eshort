package admin

import (
	"eshort/pkg/base_model"
	"eshort/pkg/password"
	"gorm.io/gorm"
)

// Admin 需要自动入库，修改表名，不能为Author
type Admin struct {
	base_model.BaseModel
	//自定义验证规则。自定义错误信息。xx_msg为tag column指定列名 gorm:"-" —— 设置 GORM 在读写时略过此字段
	Phone    string `form:"phone" comment:"手机号"  validate:"len=11,phone"`
	RealName string `form:"real_name" comment:"真实姓名"   validate:"required,min=2,max=16"`
	IdCard   string `form:"id_card" comment:"身份证号"   validate:"required,idcard"`
	QQ       string `form:"qq"  comment:"QQ号"  validate:"required,number,min=5,max=14"`
	Password string `form:"password"  comment:"密码"  validate:"required,min=6"`
	IsRead   string `form:"is_read" gorm:"-" comment:"用户协议"   validate:"required,number,eq=1" eq_msg:"请确认用户协议"`
	Email    string `column:"email"`
	//model忽略字段放最后，用于valida验证
	ConfirmPassword string `gorm:"-" form:"confirm_password"  comment:"重复密码"  validate:"required,eqfield=Password" eqfield_msg:"两次密码不相同"`
}

// BeforeSave 可以覆盖BeforeCreate和BeforeUpdate
func (author *Admin) BeforeSave(tx *gorm.DB) (err error) {
	if !password.IsHash(author.Password) {
		author.Password = password.Hash(author.Password)
	}
	return nil
}
