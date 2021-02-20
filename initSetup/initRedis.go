package initSetup

import (
	"context"
	"github.com/go-redis/redis/v8"
	"pcrweb/config"
	"pcrweb/logger"
	"sync"
)

type RMethods struct {
}

var ctx = context.Background()
var onceR sync.Once

func SetupRedis() {
	onceR.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     config.RedisCfg.Addr,
			Password: config.RedisCfg.Password,
			DB:       0, // use default DB
		})
		pong, err := client.Ping(ctx).Result()
		if err != nil {
			logger.Log.Fatal(err)
		} else {
			logger.Log.Info("Redis服务器启动成功", pong)
			config.Redis = client
		}
	})

}
