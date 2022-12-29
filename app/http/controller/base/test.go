package base

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Test struct {
	BaseController
}

func (a Test) Test(c *gin.Context) {
	//在这里测试一些小功能
	c.JSON(http.StatusOK, gin.H{
		"message": 1,
	})
}
