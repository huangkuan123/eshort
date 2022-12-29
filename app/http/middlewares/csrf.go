package middlewares

import (
	"eshort/pkg/auth"
	"eshort/pkg/csrf_token"
	rsp "eshort/pkg/response"
	"github.com/gin-gonic/gin"
)

func Csrf() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := auth.GetUID()
		if uid == "" {
			c.Abort()
			rsp.ErrMsg(c, "请重新登陆")
			return
		}
		if !csrf_token.Check(c, uid) {
			c.Abort()
			rsp.ErrMsg(c, "请求验证失败")
			return
		}
		c.Next()
	}
}
