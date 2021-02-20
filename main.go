package main

import (
	"net/http"
	"pcrweb/initSetup"
	"pcrweb/logger"
	"time"
)

func main() {

	initSetup.SetupLog()
	router := initSetup.SetupRouter()
	initSetup.SetupDB()
	//initSetup.SetupRedis()
	service := &http.Server{
		Addr:           ":8001",
		Handler:        router,
		ReadTimeout:    50 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//server.RoleImgUpload()
	logger.Log.Fatal(service.ListenAndServe())
	logger.Log.Info("服务启动成功！")
}
