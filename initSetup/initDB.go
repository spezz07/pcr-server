package initSetup

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/xormplus/xorm"
	"os"
	"pcrweb/config"
	"pcrweb/logger"
	"sync"
	"time"
)

var engine *xorm.Engine
var onceDb sync.Once

func SetupDB() {
	onceDb.Do(func() {
		//xorm reverse mysql root:admin@/go_admin?charset=utf8 models
		cfg := config.InitDBConfig(config.Config)
		dbString := cfg.Username + ":" + cfg.Password + "@(" + cfg.Path + ")/" + cfg.Dbname + "?" + cfg.Config
		var err error
		engine, err = xorm.NewEngine(cfg.Type, dbString)
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"ErrMes": err,
			}).Errorln("数据库启动异常！")
			os.Exit(1)
		} else {
			config.WxAppCfg = config.InitWxAppConfig(config.Config)
			config.RedisCfg = config.InitRedisConfig(config.Config)
			config.DB = engine
			config.DB.ShowSQL(true)
			//config.DB.SetMapper(core.GonicMapper{})
			err = config.DB.RegisterSqlMap(xorm.Xml("./mapper", ".xml"))
			err = config.DB.RegisterSqlTemplate(xorm.Pongo2("./stpl", ".stpl"))
			time.LoadLocation("Asia/Shanghai")
			if err != nil {
				logger.Log.Error("读取出错")
				logger.Log.Fatal(err.Error())
			}
			config.DB.Ping()

			if err != nil {
				logger.Log.Error("watch error")
				logger.Log.Fatal(err.Error())
			}
			logger.Log.Info("数据库启动成功")
		}
		//if db,err:=gorm.Open(cfg.Type, dbString);err != nil{
		//	logger.Log.WithFields(logrus.Fields{
		//			"ErrMes": err,
		//	}).Errorln("数据库启动异常！")
		//	os.Exit(1)
		//}else{
		//	//config.DB = db
		//	//config.DB.DB().SetMaxIdleConns(cfg.MaxIdleConns)
		//	//config.DB.DB().SetMaxOpenConns(cfg.MaxOpenConns)
		//	//config.DB.LogMode(cfg.LogMode)
		//	//config.DB.SingularTable(true)
		//	//logger.Log.Info("数据库启动成功")
		//	//config.DB.AutoMigrate(
		//	//	&model.UserModel{},
		//	//)
		//	logger.Log.Info("数据库注册成功")
		//}
	})

}
