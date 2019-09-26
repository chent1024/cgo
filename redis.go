package cgo

import (
	"github.com/go-redis/redis"
)

func NewRedis() *redis.Client {
	var cfg redis.Options
	conf := Config.Redis
	cfg = redis.Options{
		Network:  conf.Network,
		Addr:     conf.Address,
		Password: conf.Password,
		DB:       conf.Db,
	}
	client := redis.NewClient(&cfg)
	_, err := client.Ping().Result()
	if err != nil {
		return nil
	}

	return client
}
