package config

import "eshort/pkg/config"

func init() {
	config.Add("eshort", config.StrMap{
		"key":         config.Env("ESHORT_KEY", "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"),
		"pool_max":    config.Env("ESHORT_KEY_CACHE_POOL_MAX", 300),
		"app_type":    config.Env("ESHORT_APP_TYPE", "shorturl"),
		"ext":         config.Env("ESHORT_APP_EXT", "localhost"),
		"grow":        config.Env("ESHORT_KEY_GROW", 0.6),
		"grow_type":   config.Env("ESHORT_GROW_TYPE", "default"),
		"clash_retry": config.Env("ESHORT_CLASH_RETRY", 3),
	})
}
