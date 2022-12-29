/*
 * @Author: zc
 * @Date: 2022-01-12 10:38:28
 * @LastEditTime: 2022-01-12 10:40:51
 * @LastEditors: zc
 * @Description:
 * @FilePath: \gohub\config\log.go
 */
package config

import "eshort/pkg/config"

func init() {
	config.Add("log", config.StrMap{
		"mysql": map[string]interface{}{
			"level":      config.Env("LOG_LEVEL", "debug"),
			"type":       config.Env("LOG_TYPE", "single"),
			"filename":   config.Env("LOG_NAME", "storage/logs/logs.log"),
			"max_size":   config.Env("LOG_MAX_SIZE", 64),
			"max_backup": config.Env("LOG_MAX_BACKUP", 5),
			"max_age":    config.Env("LOG_MAX_AGE", 30),
			"compress":   config.Env("LOG_COMPRESS", false),
		},
	})
}
