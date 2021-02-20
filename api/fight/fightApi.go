package fight

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"pcrweb/config"
	"pcrweb/logger"
	"pcrweb/model"
	"pcrweb/response"
	"pcrweb/server"
	"pcrweb/utils"
)

func FightList(ctx *gin.Context) {
	var fightInfo model.FightResultPageInfo
	token := ctx.GetHeader("token")
	userId, err := config.Redis.Get(context.Background(), token).Result()
	ctx.ShouldBind(&fightInfo)
	if userId == "" {
		response.Fail("token失效，请重新登录！", ctx)
		return
	}
	fightInfo.UserId = userId
	page, data, err := server.FightServer.GetPageList(fightInfo)
	if err != nil {
		logger.Log.Errorf("竞技场列表错误：%v", err)
		//panic(err.Error())
		response.Fail(err.Error(), ctx)
		return
	} else {
		response.SuccessPage(data, "success", page, ctx)
	}
}

func FightAdd(ctx *gin.Context) {
	var fightInfo model.FightResultInfo
	ctx.ShouldBind(&fightInfo)
	token := ctx.GetHeader("token")
	uuid, err := utils.VerifyToken(token, []byte(config.SecretKey))
	fightInfo.UserId = uuid
	_, err = server.FightServer.Add(fightInfo)
	if err != nil {
		fmt.Println(err.Error())
		//panic(err.Error())
		response.Fail(err.Error(), ctx)
		return
	} else {
		response.Success(nil, "操作成功！", ctx)
	}
}

func FightLike(ctx *gin.Context) {
	var fightLike model.FightLike
	ctx.ShouldBind(&fightLike)
	if fightLike.Status == 0 {
		response.Fail("状态错误！", ctx)
		return
	}
	if fightLike.FightResultId == 0 {
		response.Fail("id不能为空！", ctx)
		return
	}
	token := ctx.GetHeader("token")
	uuid, err := utils.VerifyToken(token, []byte(config.SecretKey))
	fightLike.UserId = uuid
	err = server.FightServer.ApproveHandle(fightLike)
	if err != nil {
		fmt.Println(err.Error())
		//panic(err.Error())
		response.Fail(err.Error(), ctx)
		return
	} else {
		response.Success(nil, "操作成功！", ctx)
	}
}
