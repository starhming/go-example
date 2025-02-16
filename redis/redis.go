package redis

import "github.com/redis/go-redis/v9"

var redisCli *redis.Client

func init() {
	redisCli = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
}

type SetScorer interface {
	SetScore(val float64)
}
