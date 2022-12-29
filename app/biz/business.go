package biz

import "github.com/gin-gonic/gin"

type Business interface {
	//GenerateVali 转换入参检测
	GenerateVali(str string) (bool, error)
	//GenerateResult 根据key转换为响应数据
	GenerateResult(key string, data gin.H) gin.H
	//访问入参检测，验证是否符合格式
	AgentVali(str string) (bool, error)
	//访问入参提取key，从格式中提取skey
	ExtracKey(str string) string
	//访问响应
	AgentResult(key string, domain string, data gin.H) (string, gin.H, string)
	//todo 还要加黑名单检测，不同的业务可能有不同的检测方法。如url，就不允许有违法网站的地址。
	//跳转和转换，可能分别由不同的黑名单逻辑和调用顺序。
	GetExt() string
}
