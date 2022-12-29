package routes

import (
	"eshort/app/http/controller"
	"github.com/gin-gonic/gin"
)

func SetupApiRoute(route *gin.Engine) {
	index := controller.Index{}
	route.POST("/generate", index.Generate)
	route.GET("/:key", index.Agent)
}
