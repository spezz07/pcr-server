package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"pcrweb/config"
	"pcrweb/model"
	"pcrweb/response"
	"pcrweb/utils"
)


func VerifyToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			response.FailCodeMessage(402, "token不存在！", ctx)
			ctx.Abort()
			return
		}
		// redis获取token
		result, err := config.Redis.Get(context.Background(), token).Result()
		//uuid,err:=utils.VerifyToken(token, []byte(config.SecretKey))
		if result == "" {
			response.FailCodeMessage(402, "token失效，请重新登录！", ctx)
			ctx.Abort()
			return
		}
		if err != nil {
			fmt.Println(err.Error())
			_, verErr := utils.VerifyToken(token, []byte(config.SecretKey))
			if verErr.Error() == "Token is expired" {
				response.FailCodeMessage(402, "token过期，请重新登录！", ctx)
				ctx.Abort()
				return
			}
			response.FailCodeMessage(402, "token失效，请重新登录！", ctx)
			ctx.Abort()
			return
		}
		var user model.UserModel
		_, err = config.DB.Table("pcr_user").Where("user_id = ?", result).Get(&user)
		if user.Permission == "0" {
			response.FailCodeMessage(402, "无查询权限！", ctx)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
