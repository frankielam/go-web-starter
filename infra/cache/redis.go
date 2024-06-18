package cache

import (
	"context"
	"gowebstarter/infra/config"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	rdb       *redis.Client
	redisOnce sync.Once
)

func GetRedisClient() *redis.Client {
	redisOnce.Do(func() {
		conf := config.GetConf()
		rdb = redis.NewClient(&redis.Options{
			Addr:     conf.Redis.Addr,
			Password: conf.Redis.Password,
			DB:       conf.Redis.DB,
		})
		err := rdb.Ping(context.TODO()).Err()
		if err != nil {
			panic(err)
		}
		log.Println("redis connected")
	})
	return rdb
}
