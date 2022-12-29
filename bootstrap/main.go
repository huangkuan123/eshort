package bootstrap

import (
	"eshort/routes"
)

func Start() {
	SetupLogger()
	SetUpDB()
	SetUpRedis()
	//SetUpRabbitMQ()
	//业务相关的初始化
	APPINIT()
	//路由启动要放到最后，不然打印不出来日志
	routes.SetupRoute()
}
