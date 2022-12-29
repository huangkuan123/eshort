package config

import "eshort/pkg/config"

func init() {
	config.Add("app", config.StrMap{
		// 应用名称，暂时没有使用到
		"name": config.Env("APP_NAME", "AuthorBook"),
		// 当前环境，用以区分多环境
		"env": config.Env("APP_ENV", "production"),
		// 是否进入调试模式
		"debug": config.Env("APP_DEBUG", false),
		// 应用服务端口
		"port": config.Env("APP_PORT", "8000"),
		//在 Cookie 中加密数据时使用
		"key": config.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),

		"url": config.Env("APP_URL", ""),
		// 会话的 Cookie 名称
		"session_name": config.Env("SESSION_NAME", "egin_session"),
		//csrf_token 的cookie名称
		"csrf_cookie_name": config.Env("CSRF_COOKIE_NAME", "csrf_token"),
		"csrf_exp":         config.Env("CSRF_EXP", 10),
		//适用本地开发http环境secure=false
		"csrf_secure":     config.Env("CSRF_SECURE", false),
		"csrf_cache_name": config.Env("CSRF_CACHE_NAME", "csrf_token:admin:"),
	})
}
