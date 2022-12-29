package rsp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BaseResponse(c *gin.Context, status bool, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": msg,
	})
}

func SuccessMsg(c *gin.Context, msg ...string) {
	s := "操作成功"
	if len(msg) == 1 {
		s = msg[0]
	}
	BaseResponse(c, true, s)
	return
}

func ErrMsg(c *gin.Context, msg string) {
	BaseResponse(c, false, msg)
	return
}

// Abort500 响应 500，未传参 msg 时使用默认消息
func Abort500(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": "服务器内部错误，请稍后再试",
	})
}

func RepData(c *gin.Context, data interface{}, msg ...string) {
	s := "success"
	if len(msg) == 1 {
		s = msg[0]
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": s,
		"data":    data,
	})
}
