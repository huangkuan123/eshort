package bootstrap

import (
	"eshort/app/biz"
	"eshort/app/services/shortkey"
	_ "eshort/app/services/shortkey/factory"
	fc "eshort/app/services/shortkey/factory"
	"eshort/pkg/ehash"
	_ "eshort/pkg/ehash"
	"eshort/pkg/eviction/elru"
)

func APPINIT() {
	ehash.Setup()
	fc.Setup()
	shortkey.Setup()
	fc.Tidy()
	//按需启动监控数据池程序
	shortkey.SuperPool()
	//策略选择
	biz.SetBusiness()
	elru.NewLRU(200)
}
