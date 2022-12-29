/*
 * @Author: zc
 * @Date: 2022-01-14 16:14:00
 * @LastEditTime: 2022-01-17 18:18:50
 * @LastEditors: zc
 * @Description:
 * @FilePath: \gohub\app\http\middlewares\recovery.go
 */
package middlewares

import (
	"eshort/pkg/logger"
	rsp "eshort/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, true)

				// 链接中断，无需记录堆栈信息
				var brokerPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						errStr := strings.ToLower(se.Error())
						if strings.Contains(errStr, "broken pipe") || strings.Contains(errStr, "connection reset by peer") {
							brokerPipe = true
						}
					}
				}
				if brokerPipe {
					logger.Error(c.Request.URL.Path,
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					c.Error(err.(error))
					c.Abort()
					return
				}
				//todo 按行打印输出
				logger.Error(
					"recovery from panic",
					//zap.Time("time", time.Now()),               // 记录时间
					zap.Any("error", err), // 记录错误信息
					//zap.String("request", string(httpRequest)), // 请求信息
					zap.Stack("stacktrace"), // 调用堆栈信息
				)
				rsp.Abort500(c)
			}
		}()
		c.Next()
	}
}
