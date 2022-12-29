/*
 * @Author: zc
 * @Date: 2022-01-12 10:42:42
 * @LastEditTime: 2022-01-13 11:22:57
 * @LastEditors: zc
 * @Description:
 * @FilePath: \gohub\bootstrap\logger.go
 */
package bootstrap

import (
	"eshort/pkg/config"
	"eshort/pkg/logger"
)

func SetupLogger() {
	logger.InitLogger(
		config.GetString("log.filename"),
		config.GetInt("log.max_size"),
		config.GetInt("log.max_backup"),
		config.GetInt("log.max_age"),
		config.GetBool("log.compress"),
		config.GetString("log.type"),
		config.GetString("log.level"),
	)
}
