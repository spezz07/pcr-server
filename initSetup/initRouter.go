package initSetup

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pcrweb/config"
	"pcrweb/logger"
	"pcrweb/middleware"
	route "pcrweb/router"
)

type Login struct {
	User     string `form:"user" json:"user"  binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func SetupRouter() *gin.Engine {
	env := config.Config.Sub("gin").Get("mode")
	fmt.Printf("env:%v", env)
	if env == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(middleware.Cors()).Use(middleware.ReqsLogs()).Use(middleware.GinPanicHandle())
	ApiGroup := router.Group("")
	route.InitUserRouter(ApiGroup)
	route.InitRoleRouter(ApiGroup)
	route.InitFightRouter(ApiGroup)
	router.Static("/static", "./static/")
	logger.Log.Info("路由初始化成功！")
	return router
}
