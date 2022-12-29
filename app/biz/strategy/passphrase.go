package strategy

import (
	"errors"
	"eshort/pkg/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"unicode/utf8"
)

//口令项目
type Passphrate struct {
}

func (p Passphrate) GenerateVali(str string) (bool, error) {
	//这里应当根据具体业务来配置验证，有可能是个url，也有可能仅是参数。
	return true, nil
}

func (p Passphrate) GenerateResult(key string, data gin.H) gin.H {
	data["data"] = strings.Replace(p.GetExt(), "$$$$", key, 1)
	return data
}

func (p Passphrate) AgentVali(str string) (bool, error) {
	sl := utf8.RuneCountInString(str)
	if sl < 3 {
		return false, errors.New("口令有误")
	}
	ext := p.GetExt()
	if sl < utf8.RuneCountInString(ext) {
		return false, errors.New("口令有误")
	}
	split := strings.Split(ext, "$$$$")
	if len(split) != 2 {
		return false, errors.New("口令系统配置有误")
	}
	start := split[0]
	fmt.Println("start:", start)
	se := string([]rune(str)[:utf8.RuneCountInString(start)])
	fmt.Println("se:", se)
	if se != start {
		return false, errors.New("口令有误")
	}
	return true, nil
}

func (p Passphrate) ExtracKey(str string) string {
	ext := p.GetExt()
	split := strings.Split(ext, "$$$$")
	start := split[0]
	end := split[1]
	str = strings.Replace(str, start, "", 1)
	str = strings.Replace(str, end, "", 1)
	return str
}

func (p Passphrate) AgentResult(key string, full_data string, data gin.H) (string, gin.H, string) {
	data["data"] = full_data
	return full_data, data, "response"

}

func (p Passphrate) GetExt() string {
	return config.GetString("eshort.ext")
}
