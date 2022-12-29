package csrf_token

import (
	"eshort/pkg/config"
	"eshort/pkg/eredis"
	"eshort/pkg/helpers"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	cookieName = config.GetString("csrf_cookie_name")
	cacheName  = config.GetString("csrf_cache_name")
	tokenLen   = 10
)

// Set 设置csrf_token
func Set(c *gin.Context, uid string) {
	token := makeToken()
	age := config.GetInt("csrf_exp")
	secure := config.GetBool("csrf_secure")
	c.SetCookie(cookieName, token, age, "/", "/", secure, true)
	eredis.RedisClient.Set(c, cacheName+uid, token, time.Duration(age)*time.Second)
}

func Check(c *gin.Context, uid string) bool {
	token := c.GetHeader("X-CSRF-Token")
	if token == "" {
		return false
	}
	cookie, err := c.Cookie(cookieName)
	if err != nil || token != cookie {
		return false
	}
	result, err := eredis.RedisClient.Get(c, cacheName+uid).Result()
	if err != nil || result != cookie {
		return false
	}
	eredis.RedisClient.Del(c, cacheName+uid)
	return true
}

func makeToken() string {
	return helpers.RandomString(tokenLen)
}
