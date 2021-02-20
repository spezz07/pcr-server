package response

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"reflect"
)

const (
	success  = 200
	fail     = 400
	error    = 500
	ResOkStr = "success"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
type ResponsePage struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Pagination
}
type Pagination struct {
	Rows      int `json:"rows"`
	PageNo    int `json:"pageNo"`
	Total     int `json:"total"`
	TotalPage int `json:"totalPage"`
}

func resultSuccessPage(code int, data interface{}, msg string, page Pagination, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ResponsePage{
		Code:       code,
		Data:       data,
		Msg:        msg,
		Pagination: page,
	})
}
func resultSuccess(code int, data interface{}, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}
func resultFail(msg string, ctx *gin.Context) {
	if msg == "" {
		msg = "请求错误！"
	}
	ctx.JSON(http.StatusOK, Response{
		Code: error,
		Data: nil,
		Msg:  msg,
	})
}
func Success(data interface{}, msg string, ctx *gin.Context) {
	if msg == "" {
		msg = "success"
	}
	resultSuccess(success, data, msg, ctx)
}

func SuccessPage(data interface{}, msg string, page Pagination, ctx *gin.Context) {
	if msg == "" {
		msg = "success"
	}
	resultSuccessPage(success, data, msg, page, ctx)
}

func Fail(msg string, ctx *gin.Context) {
	resultFail(msg, ctx)
}

func FailCodeMessage(code int, msg string, ctx *gin.Context) {
	if msg == "" {
		msg = "请求错误！"
	}
	ctx.JSON(http.StatusOK, Response{
		Code: 400,
		Data: nil,
		Msg:  msg,
	})
}

func PaginationInfo(obj interface{}, total int) (limit int, offset int, page Pagination) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name == "Pagination" {
			structField := v.Field(i).Type()
			for k := 0; k < structField.NumField(); k++ {
				if structField.Field(k).Name == "PageNo" {
					page.PageNo = v.Field(i).Field(k).Interface().(int)
				}
				if structField.Field(k).Name == "Rows" {
					page.Rows = v.Field(i).Field(k).Interface().(int)
				}
				//fmt.Printf("%s %s = %v -tag:%s \n",
				//	structField.Field(k).Name,
				//	structField.Field(k).Type,
				//	v.Field(i).Field(k).Interface(),
				//	structField.Field(k).Tag)
			}
		}
	}
	if page.PageNo == 0 || page.Rows == 0 {
		page.PageNo = 1
		page.Rows = 10
	}
	page.Total = total
	page.TotalPage = int(math.Ceil(float64((page.Total + page.Rows - 1) / page.Rows)))
	limit = page.Rows
	offset = page.Rows * (page.PageNo - 1)
	return limit, offset, page
}
