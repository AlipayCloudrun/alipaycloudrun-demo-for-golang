package app

import (
	"alipaycloudrun-demo-for-go/util"
	"fmt"
	"github.com/go-redis/redis"
)

func InitRedis() *redis.Client {
	// redis域名
	host := util.GetEnvDefault("REDIS_HOST", "127.0.0.1")
	// redis端口号
	port := util.GetEnvDefault("REDIS_PORT", "6379")
	// redis密码
	pwd := util.GetEnvDefault("REDIS_PASSWORD", "")

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: pwd, // no password set
		DB:       0,   // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Printf("redis open error, err:%+v", err.Error())
		return nil
	}

	return rdb
}
