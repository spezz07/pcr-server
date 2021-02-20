package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/xormplus/xorm"
)

const (
	DataFormatStr = "2006-01-02 15:04:05"
	SecretKey     = ""
)

var (
	DB       *xorm.Engine
	Config   *viper.Viper
	Redis    *redis.Client
	WxAppCfg *wxApp
	RedisCfg *redisCfg
)
