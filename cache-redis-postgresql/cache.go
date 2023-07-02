package main

import (
	"github.com/go-redis/redis/v8"
)

func NewCache(cfg *Config) (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.CACHE_URL,
		Password: "",
		DB:       0,
	})
	return
}
