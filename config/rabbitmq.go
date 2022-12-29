package config

import "eshort/pkg/config"

func init() {
	config.Add("rabbitmq", config.StrMap{
		"host":     config.Env("RABBITMQ_HOST", "localhost"),
		"port":     config.Env("RABBITMQ_PORT", "5672"),
		"user":     config.Env("RABBITMQ_USERNAME", "guest"),
		"password": config.Env("RABBITMQ_PASSWORD", "guest"),
		"vhost":    config.Env("RABBITMQ_VHOST", "eurl"),
	})
}
