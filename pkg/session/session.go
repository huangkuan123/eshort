package session

import (
	"context"
	"eshort/pkg/config"
	"eshort/pkg/easylogger"
	"eshort/pkg/eredis"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v8"
	"net/http"
)

// Store gorilla sessions 的存储库
//var Store = sessions.NewCookieStore([]byte(config.GetString("app.key")))

var Store *redisstore.RedisStore

// Session 当前会话
var Session *sessions.Session

// Request 用以获取会话
var Request *http.Request

// Response 用以写入会话
var Response http.ResponseWriter

// StartSession 初始化会话，在中间件中调用
func StartSession(w http.ResponseWriter, r *http.Request) {
	var err error
	// Store.Get() 的第二个参数是 Cookie 的名称
	// gorilla/sessions 支持多会话，本项目我们只使用单一会话即可
	Store, err = redisstore.NewRedisStore(context.Background(), eredis.RedisClient)
	Store.KeyGen(func() (string, error) {
		return config.GetString("app.key"), err
	})
	easylogger.LogError(err, "session key 设置失败")
	Session, err = Store.Get(r, config.GetString("app.session_name"))
	easylogger.LogError(err, "开启redis,session失败")
	Request = r
	Response = w
}

// Put 写入键值对应的会话数据
func Put(key string, value interface{}) {
	Session.Values[key] = value
	Save()
}

// Get 获取会话数据，获取数据时请做类型检测
func Get(key string) interface{} {
	return Session.Values[key]
}

// Forget 删除某个会话项
func Forget(key string) {
	delete(Session.Values, key)
	Save()
}

// Flush 删除当前会话
func Flush() {
	Session.Options.MaxAge = -1
	Save()
}

// Save 保持会话
func Save() {
	// 非 HTTPS 的链接无法使用 Secure 和 HttpOnly，浏览器会报错
	// Session.Options.Secure = true
	Session.Options.HttpOnly = true
	err := Session.Save(Request, Response)
	easylogger.LogError(err, "保存session失败")
}
