package base

import (
	"eshort/app/models/admin"
	"eshort/app/requests"
	"eshort/pkg/auth"
	"eshort/pkg/base_model"
	"eshort/pkg/eredis"
	"eshort/pkg/response"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	BaseController
}

// Register 注册
func (a *AuthController) Register(c *gin.Context) {
	var Author admin.Admin
	err := requests.Validata(c, &Author)
	if err != nil {
		rsp.ErrMsg(c, err.Error())
		return
	}
	_, err = admin.GetUserByPhone(Author.Phone)
	if nil == err {
		rsp.ErrMsg(c, "手机号已经注册")
		return
	}
	//新增用户
	err = Author.Create()
	if err != nil {
		rsp.ErrMsg(c, err.Error())
		return
	}
	rsp.SuccessMsg(c, "注册成功")
	return
}

// Login 登录
func (a *AuthController) Login(c *gin.Context) {
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	err := auth.Attempt(phone, password)
	if err != nil {
		rsp.ErrMsg(c, err.Error())
		return
	}
	rsp.SuccessMsg(c, "登陆成功")
}

// Logout 登出
func (a *AuthController) Logout(c *gin.Context) {
	uid := auth.GetUID()
	eredis.RedisClient.Del(c, "csrf_token:admin:"+uid)
	auth.Logout()
	rsp.SuccessMsg(c, "退出登陆成功")
}

// GetUser 获取当前登录用户信息
func (a *AuthController) GetUser(c *gin.Context) {
	user := auth.User()
	vo := admin.Vo{}
	base_model.ToVo(&vo, &user)
	rsp.RepData(c, vo)
}
