package requests

type RegisterVali struct {
	//自定义验证规则。自定义错误信息。xx_msg为tag
	Phone           string `form:"phone" comment:"手机号"  validate:"len=11,phone"`
	RealName        string `form:"real_name" comment:"真实姓名"   validate:"required,min=2,max=16"`
	IdCard          string `form:"id_card" comment:"身份证号"   validate:"required,idcard"`
	QQ              string `form:"qq"  comment:"QQ号"  validate:"required,number,min=5,max=14"`
	Password        string `form:"password"  comment:"密码"  validate:"required,min=6"`
	ConfirmPassword string `form:"confirm_password"  comment:"重复密码"  validate:"required,eqfield=Password" eqfield_msg:"两次密码不相同"`
	IsRead          string `form:"is_read" comment:"用户协议"   validate:"required,number,eq=1" eq_msg:"请确认用户协议"`
}
