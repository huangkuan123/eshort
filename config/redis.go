package config

import "eshort/pkg/config"

func init() {
	config.Add("redis", config.StrMap{
		"host":      config.Env("REDIS_HOST", "127.0.0.1"),
		"port":      config.Env("REDIS_PORT", "6379"),
		"password":  config.Env("REDIS_PASSWORD", ""),
		"select_db": config.Env("REDIS_SELECT_DB", "0"),
	})
}
