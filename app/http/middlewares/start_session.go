package middlewares

import (
	"eshort/pkg/session"
	"github.com/gin-gonic/gin"
)

func StartSession() gin.HandlerFunc {
	//return func(writer http.ResponseWriter, request *http.Request) {
	//	session.StartSession(writer, request)
	//	next.ServeHTTP(writer, request)
	//}
	return func(c *gin.Context) {
		session.StartSession(c.Writer, c.Request)
	}
}
