package base

import (
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	BaseController
}

func (a *IndexController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
