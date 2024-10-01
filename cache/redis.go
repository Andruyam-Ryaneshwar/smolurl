package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"smolurl/config"
)

var ctx = context.Background()
var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisAddr,
		Password: config.AppConfig.RedisPwd,
		DB:       0,
	})
}

func GetURL(shortUrl string) (string, error) {
	return RedisClient.Get(ctx, shortUrl).Result()
}

func SetURL(shortUrl string, originalUrl string) error {
	return RedisClient.Set(ctx, shortUrl, originalUrl, 0).Err()
}
