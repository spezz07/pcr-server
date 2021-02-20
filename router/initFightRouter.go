package router

import (
	"github.com/gin-gonic/gin"
	"pcrweb/api/fight"
	"pcrweb/middleware"
)

func InitFightRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("/fight").Use(middleware.VerifyToken())
	{
		ApiRouter.POST("/add", fight.FightAdd)
		ApiRouter.POST("/like", fight.FightLike)
		ApiRouter.POST("/getList", fight.FightList)
	}
}
