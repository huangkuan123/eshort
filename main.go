package main

import (
	"eshort/bootstrap"
	"eshort/config"
	"fmt"
)

func init() {
	fmt.Println("开始初始化环境配置")
	config.Initialize()
}

func main() {
	//fmt.Println(pc.Get("eshort.clash_retry"))
	fmt.Println("开始启动")
	bootstrap.Start()
}
