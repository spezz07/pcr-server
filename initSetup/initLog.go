package initSetup

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"path"
	"pcrweb/config"
	"pcrweb/logger"
	"time"
)

func SetupLog() {
	logInit := log.New()
	logInit.SetReportCaller(true)
	logInit.SetFormatter(&log.TextFormatter{
		TimestampFormat: config.DataFormatStr + ".000",
		FullTimestamp:   true,
	})
	//log.SetOutput(os.Stdout)
	baseLogPath := path.Join("./", "logs_")
	writer, err := rotatelogs.New(
		baseLogPath+"%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(baseLogPath),
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("system logger error. %+v", errors.WithStack(err))
	}
	writeMap := lfshook.WriterMap{
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.PanicLevel: writer,
	}
	lfHook := lfshook.NewHook(writeMap, &log.TextFormatter{
		TimestampFormat: config.DataFormatStr + ".000",
		FullTimestamp:   true,
	})
	logInit.AddHook(lfHook)
	logger.Log = logInit
	logger.Log.Info("---日志初始化成功---")
}
