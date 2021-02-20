package user

import (
	"github.com/gin-gonic/gin"
	"pcrweb/logger"
	"pcrweb/model"
	"pcrweb/response"
	"pcrweb/server"
)

func UserRegister(ctx *gin.Context) {
	var userInfo model.UserModel
	ctx.ShouldBind(&userInfo)
	if userInfo.Account == "" {
		response.Fail("用户名不能为空！", ctx)
		return
	}
	if userInfo.Password == "" {
		response.Fail("密码不能为空！", ctx)
	}
	err := server.UserSignUp(userInfo)
	if err != nil {
		logger.Log.Errorf("用户注册错误：%v", err)
		response.Fail(err.Error(), ctx)
	} else {
		response.Success(nil, "注册成功！", ctx)
	}
}

func UserWxLogin(ctx *gin.Context) {
	var userInfo model.UserWxModel
	_ = ctx.ShouldBind(&userInfo)
	if userInfo.Code == "" {
		response.Fail("code不能为空！", ctx)
	}
	token, err := server.UserWxLogin(userInfo)
	if err != nil {
		logger.Log.Errorf("用户注册错误：%v", err)
		response.Fail(err.Error(), ctx)
	} else {
		response.Success(token, "登录成功！", ctx)
	}
}

func UserLogin(ctx *gin.Context) {
	var userInfo model.UserModel
	ctx.ShouldBind(&userInfo)
	if userInfo.Account == "" {
		response.Fail("用户名不能为空！", ctx)
		return
	}
	if userInfo.Password == "" {
		response.Fail("密码不能为空！", ctx)
		return
	}
	token, err := server.Login(userInfo)
	if err != nil {
		logger.Log.Errorf("用户登录错误：%v", err)
		return
	}
	//tokenObj := make(map[string]string)
	//tokenObj["token"] = token
	response.Success(token, response.ResOkStr, ctx)
}

func UserList(ctx *gin.Context) {
	var userList model.UserList
	ctx.ShouldBind(&userList)
	page, data, err := server.GetUserList(userList)
	if err != nil {
		logger.Log.Errorf("用户列表请求错误：%v", err)
		response.Fail("内部错误", ctx)
		return
	} else {
		response.SuccessPage(data, "", page, ctx)
	}
}
