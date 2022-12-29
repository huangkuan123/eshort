package middlewares

import (
	"eshort/pkg/auth"
	rsp "eshort/pkg/response"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !auth.Check() {
			c.Abort()
			rsp.ErrMsg(c, "请重新登陆")
			return
		}
		c.Next()
	}
}
