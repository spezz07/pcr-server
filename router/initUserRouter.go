package router

import (
	"github.com/gin-gonic/gin"
	"pcrweb/api/user"
	"pcrweb/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("/user")
	{
		ApiRouter.POST("/signIn", user.UserRegister)
		ApiRouter.POST("/login", user.UserLogin)
		ApiRouter.POST("/wxlogin", user.UserWxLogin)
		ApiRouter.POST("/pdf",user.Updf)
		ApiRouter.POST("/getList", user.UserList).Use(middleware.VerifyToken())
	}
}
