package strategy

import (
	"errors"
	"eshort/pkg/config"
	"github.com/gin-gonic/gin"
	"net/url"
)

//短链接项目
type ShortURL struct {
}

func (a ShortURL) GenerateVali(str string) (bool, error) {
	if len(str) < 3 {
		return false, errors.New("地址有误")
	}
	curl, err := url.ParseRequestURI(str)
	if err != nil || curl.Scheme == "" {
		return false, errors.New("地址有误")
	}
	return true, nil
}

func (a ShortURL) GenerateResult(key string, data gin.H) gin.H {
	data["url"] = a.GetExt() + key
	return data
}

func (a ShortURL) AgentVali(str string) (bool, error) {
	lstr := len(str)
	if lstr < 3 || lstr > 10 {
		return false, errors.New("地址有误" + string(lstr))
	}
	return true, nil
}

func (a ShortURL) AgentResult(key string, full_data string, data gin.H) (result string, H gin.H, rtype string) {
	return full_data, data, "redirect"
}

func (a ShortURL) ExtracKey(str string) string {
	return str
}

func (a ShortURL) GetExt() string {
	return config.GetString("eshort.ext")
}
