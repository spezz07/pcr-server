package role

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pcrweb/model"
	"pcrweb/response"
	"pcrweb/server"
)

func RoleList(ctx *gin.Context) {
	var pcrRolePageList model.PcrRolePageList
	ctx.ShouldBind(&pcrRolePageList)
	page, data, err := server.RoleServer.GetPageList(pcrRolePageList)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
		//response.Fail("内部错误", ctx)
		return
	} else {
		response.SuccessPage(data, "", page, ctx)
	}
}

func OKList(ctx *gin.Context) {
	response.Success(nil, "test", ctx)
}
