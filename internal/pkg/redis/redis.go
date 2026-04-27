package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	redisClient *redis.Client
}

func NewRedisService(redisUrl string) (*RedisService, error) {

 	redisClient := redis.NewClient(&redis.Options{
		Addr: redisUrl,
	})

	err := redisClient.Ping(context.Background()).Err()

	if err != nil {
		return nil, fmt.Errorf("Failed to connect to Redis: %w", err)
	}

	return &RedisService{redisClient: redisClient}, nil
}

func (r *RedisService) Set(ctx context.Context, key string, value string, ttl time.Duration) (bool, error) {

	err := r.redisClient.Set(ctx, key, value, ttl).Err();

	if err != nil {
		log.Println("Error setting value in redis", err);
		return false, err;
	}

	return true, nil;
}

func (r *RedisService) Get(ctx context.Context, key string) (string, error) {

	value, err := r.redisClient.Get(ctx, key).Result();

	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", fmt.Errorf("Error getting value from redis - %s", err.Error())
	}

	return value, nil
 }

func (r *RedisService) Delete(ctx context.Context, key string) (bool, error) {

	err := r.redisClient.Del(ctx, key).Err();

	if err != nil {
		log.Println("Error deleting in redis", err)
		return false, err
	}

	return true, nil
}
