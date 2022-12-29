package validata_rules

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

//所有自定义验证规则

func ValiPhone(fl validator.FieldLevel) bool {
	matched, _ := regexp.Match("^1[345789]{1}\\d{9}$", []byte(fl.Field().String()))
	return matched
}

func ValiIdCard(fl validator.FieldLevel) bool {
	matched, _ := regexp.Match("(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)", []byte(fl.Field().String()))
	return matched
}
