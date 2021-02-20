package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"pcrweb/config"
)

func ReqsLogs() gin.HandlerFunc {
	logInit := log.New()
	logInit.SetReportCaller(true)
	return func(ctx *gin.Context) {
		logInit.SetFormatter(&log.TextFormatter{
			TimestampFormat: config.DataFormatStr + ".000",
			FullTimestamp:   true,
		})
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		//data, err := ctx.GetRawData()
		s := make(map[string]interface{})
		_ = jsoniter.Unmarshal(body, &s)
		data, _ := jsoniter.Marshal(s)
		//sR:= json.RawMessage([]byte(str))
		//if err != nil {
		//	logger.Log.Warn(err.Error())
		//}
		//ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		logInit.WithFields(log.Fields{
			"url":     ctx.Request.URL,
			"methods": ctx.Request.Method,
			"IP":      ctx.ClientIP(),
			"req":     ctx.Request.PostForm.Encode(),
			"data":    string(data),
		}).Info("请求参数:")
		ctx.Next()
	}
}
