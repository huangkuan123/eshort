/*
 * @Author: zc
 * @Date: 2022-01-12 10:34:59
 * @LastEditTime: 2022-01-27 14:38:45
 * @LastEditors: zc
 * @Description:
 * @FilePath: \gohub\pkg\app\app.go
 */
package app

import (
	"eshort/pkg/config"
	"time"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}

func TimenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(chinaTimezone)
}

func URL(path string) string {
	return config.GetString("app.url") + path
}

func V1URL(path string) string {
	return URL("/v1/" + path)
}
