package bootstrap

import (
	"eshort/pkg/eredis"
)

func SetUpRedis() {
	eredis.ConnectRedis()
}
