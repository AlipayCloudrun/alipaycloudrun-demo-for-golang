package service

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

type redisService struct {
	client *redis.Client
}

func NewRedisService(client *redis.Client) IRedisService {
	return &redisService{
		client: client,
	}
}

func (r *redisService) Set(key string, value interface{}) error {
	if r.client == nil {
		return fmt.Errorf("redis连接为空，请检查环境变量")
	}
	log.Println("Set redis value")
	err := r.client.Set(key, value, 0).Err()

	if err != nil {
		return err
	}

	return nil
}

func (r *redisService) Get(key string) (interface{}, error) {
	if r.client == nil {
		return nil, fmt.Errorf("redis连接为空，请检查环境变量")
	}
	log.Println("get Redis value")
	result, err := r.client.Get(key).Result()
	if err == redis.Nil {
		result = "nil"
	} else if err != nil {
		return nil, err
	}

	return result, nil
}
