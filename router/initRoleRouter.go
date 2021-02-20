package router

import (
	"github.com/gin-gonic/gin"
	"pcrweb/api/role"
)

func InitRoleRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("/role")
	//ApiRouter := Router.Group("/role").Use(middleware.VerifyToken())
	{
		ApiRouter.POST("/getList", role.RoleList)
		ApiRouter.GET("/test", role.OKList)
	}
}
