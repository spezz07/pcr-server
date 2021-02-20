package impl

import "time"

type RedisMethods interface {
	RedisValueSetExp(key string, val string, expTime time.Duration) (err error)
	RedisValueGet(key string) (result string, err error)
}
