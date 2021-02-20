package middleware

import (
	"github.com/gin-gonic/gin"
	"pcrweb/logger"
	"pcrweb/response"
	"runtime/debug"
)

func GinPanicHandle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.Error(err)
				logger.Log.Error(string(debug.Stack()))
				response.Fail("内部错误", ctx)
				return
			}
		}()
		//ctx.Next()
	}
}
