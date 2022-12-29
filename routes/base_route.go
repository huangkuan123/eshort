package routes

import (
	"eshort/app/http/middlewares"
	"eshort/pkg/config"
	"eshort/pkg/easylogger"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoute() {
	route := gin.New()
	registerGlobalMiddleWare(route)
	SetupApiRoute(route)
	port := config.GetString("app.port")
	//err := route.Run(":" + port) #windows 可以放开此行，屏蔽gracehttp这行
	err := gracehttp.Serve(&http.Server{Addr: ":" + port, Handler: route})
	if err != nil {
		easylogger.Print(err.Error())
	}
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		middlewares.Logger(),
		middlewares.Recovery(),
		middlewares.StartSession(),
		middlewares.Cors(),
	)
}

// 适用于为组添加中间件
func registerGroupsMiddleWare(middleware []gin.HandlerFunc, groupRoutes ...*gin.RouterGroup) {
	for _, groupRoute := range groupRoutes {
		groupRoute.Use(middleware...)
	}
}

func registerGroupsRBACMiddleware(groupRoutes ...*gin.RouterGroup) {
	for _, groupRoute := range groupRoutes {
		//groupRoute.Use(middlewares.Auth(), middlewares.RbacAuth())
		groupRoute.Use(middlewares.Auth())
	}
}
