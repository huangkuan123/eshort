package admin

import "eshort/pkg/base_model"

type Vo struct {
	base_model.BaseModel
	Phone    string ``
	RealName string ``
	IdCard   string ``
	QQ       string ``
	Email    string `column:"email"`
}
