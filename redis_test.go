package cgo

import (
	"testing"
)

func TestNewRedis(t *testing.T) {
	err := NewRedis().Ping().Err()
	if err != nil {
		t.Error(err)
	}
}

func TestRedisClient_Set(t *testing.T) {
	NewRedis()
	_, err := Redis.Set("test", "cccc", 0).Result()
	if err != nil {
		t.Error(err)
	}
}

func TestRedisClient_Get(t *testing.T) {
	NewRedis()
	str, err := Redis.Get("test").Result()
	if err != nil || str != "cccc" {
		t.Error(err)
	}
}

func TestRedisClient_Del(t *testing.T) {
	NewRedis()
	d, err := Redis.Del("test").Result()
	t.Log(d)
	if err != nil {
		t.Error(err)
	}
}
