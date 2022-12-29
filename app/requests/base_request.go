package requests

import (
	"errors"
	"eshort/app/requests/validata_rules"
	"eshort/pkg/easylogger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var Trans ut.Translator
var Validator *validator.Validate
var AddRulesRegister map[string]AddRulesRegisterMeta

type AddRulesRegisterMeta struct {
	FuncName     string
	FuncErrorMsg string
	FuncHandle   validator.Func
}

// init 中文验证器初始化
func init() {
	translator := zh.New()
	universalTranslator := ut.New(translator)
	Validator = validator.New()
	Trans, _ = universalTranslator.GetTranslator("zh")
	err := zhTranslations.RegisterDefaultTranslations(Validator, Trans)
	easylogger.LogError(err, "验证器中文化失败")
	Validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("comment")
	})
	registerRules()
}

// @Validata 验证入口
// @Param data 验证结构体
func Validata(c *gin.Context, data interface{}) error {
	if err := c.Bind(data); err != nil {
		return errors.New(err.Error() + "请求解析错误，请确认请求格式是否正确")
	}
	if errs := Validator.Struct(data); errs != nil {
		validationErrors := errs.(validator.ValidationErrors)
		sliceErrs := []string{}
		errMsg := setupCustomErrorMsg(data)
		for _, e := range validationErrors {
			//使用validator.ValidationErrors类型里的Translate方法进行翻译
			tk := e.StructField() + ":" + e.ActualTag()
			if v, ok := errMsg[tk]; ok {
				sliceErrs = append(sliceErrs, v)
			} else if v, ok := AddRulesRegister[e.Tag()]; ok {
				sliceErrs = append(sliceErrs, v.FuncErrorMsg)
			} else {
				sliceErrs = append(sliceErrs, e.Translate(Trans))
			}
		}

		str := strings.Join(sliceErrs, ",")
		return errors.New(str)
	}
	return nil
}

func setupCustomErrorMsg(data interface{}) map[string]string {
	dType := reflect.TypeOf(data)
	//dValue := reflect.ValueOf(data)
	//dataValue := dValue.Elem()
	dataType := dType.Elem()
	errMsg := make(map[string]string)
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		fieldName := field.Name
		validateArr := strings.Split(field.Tag.Get("validate"), ",")
		for j := 0; j < len(validateArr); j++ {
			ruleName := strings.Split(validateArr[j], "=")[0]
			tMsg := field.Tag.Get(ruleName + "_msg")
			if len(tMsg) > 0 {
				errMsg[fieldName+":"+ruleName] = tMsg
			}
		}
	}
	return errMsg
}

func registerRules() {
	AddRulesRegister = make(map[string]AddRulesRegisterMeta)
	joinRule("phone", "手机号码不符合规则", validata_rules.ValiPhone)
	joinRule("idcard", "身份证号不符合规则", validata_rules.ValiIdCard)
	//joinRule("not_exists", "", validata_rules.NotExists)
	for _, meta := range AddRulesRegister {
		err := Validator.RegisterValidation(meta.FuncName, meta.FuncHandle)
		if err != nil {
			easylogger.LogError(err, meta.FuncName+"验证规则注册失败")
		}
	}
}

func joinRule(name string, msg string, fun validator.Func) {
	meta := AddRulesRegisterMeta{
		FuncName:     name,
		FuncErrorMsg: msg,
		FuncHandle:   fun,
	}
	AddRulesRegister[name] = meta
}
