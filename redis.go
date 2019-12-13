package cgo

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type redisClient struct {
	Client *redis.Client
}

var Redis redisClient

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

	Redis.Client = client
	return client
}

func (r *redisClient) Set(key string, val interface{}, exp time.Duration) *redis.StatusCmd {
	return r.Client.Set(r.AddPrefix(key), val, exp)
}

func (r *redisClient) Get(key string) *redis.StringCmd {
	return r.Client.Get(r.AddPrefix(key))
}

func (r *redisClient) Del(keys ...string) *redis.IntCmd {
	var key []string
	for _, v := range keys {
		key = append(key, r.AddPrefix(v))
	}
	return r.Client.Del(key...)
}

func (r *redisClient) HSet(key string, field string, val interface{}) *redis.BoolCmd {
	return r.Client.HSet(r.AddPrefix(key), field, val)
}

func (r *redisClient) HDel(key string, fields ...string) *redis.IntCmd {
	return r.Client.HDel(r.AddPrefix(key), fields...)
}

// 为缓存key添加前缀
func (r *redisClient) AddPrefix(key string) string {
	var prefix string
	if Config.App.Name == "" {
		prefix = "cgo"
	} else {
		prefix = Config.App.Name
	}
	return fmt.Sprintf("%s_%s", prefix, key)
}
