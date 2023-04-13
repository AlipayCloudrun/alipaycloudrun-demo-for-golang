package service

import (
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
	log.Println("Set redis value")
	err := r.client.Set(key, value, 0).Err()

	if err != nil {
		return err
	}

	return nil
}

func (r *redisService) Get(key string) (interface{}, error) {
	log.Println("get Redis value")
	result, err := r.client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	return result, nil
}
