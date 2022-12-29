/*
 * @Author: zc
 * @Date: 2022-01-14 15:32:58
 * @LastEditTime: 2022-01-14 16:04:35
 * @LastEditors: zc
 * @Description:日志
 * @FilePath: \gohub\app\http\middlewares\logger.go
 */
package middlewares

import (
	"bytes"
	"eshort/pkg/helpers"
	"eshort/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		start := time.Now()
		c.Next()

		cost := time.Since(start)
		responseStatus := c.Writer.Status()

		logFields := []zap.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("request", c.Request.Method+" "+c.Request.URL.String()),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", helpers.MicrosecondsStr(cost)),
		}
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			requestBody, _ := c.GetRawData()
			// 请求内容
			logFields = append(logFields, zap.String("Request Body", string(requestBody)))
			// 响应内容
			logFields = append(logFields, zap.String("Response Body", w.body.String()))
		}
		if responseStatus > 400 && responseStatus <= 499 {
			logger.Warn("HTTP Warning "+cast.ToString(responseStatus), logFields...)
		} else if responseStatus >= 500 && responseStatus <= 599 {
			logger.Error("HTTP Error"+cast.ToString(responseStatus), logFields...)
		} else {
			logger.Debug("HTTP Access Log", logFields...)
		}
	}
}
