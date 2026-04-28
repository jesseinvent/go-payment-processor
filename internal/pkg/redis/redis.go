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

	redisClient, err := ConnectToRedisWithRetry(redisUrl, 5, 2*time.Second)

	if err != nil {
		return nil, err
	}

	return &RedisService{redisClient: redisClient}, nil
}

func ConnectToRedisWithRetry(redisUrl string, maxRetries int, retryInterval time.Duration) (*redis.Client, error) {
	var redisClient *redis.Client
	var err error

	for i := range maxRetries {
	 	redisClient = redis.NewClient(&redis.Options{
			Addr: redisUrl,
		})

		err = redisClient.Ping(context.Background()).Err()

		if err == nil {
			return redisClient, nil
		}

		log.Printf("Failed to connect to Redis (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(retryInterval)
	}

	return nil, fmt.Errorf("could not connect to Redis after %d attempts: %w", maxRetries, err)
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
		return "", fmt.Errorf("error getting value from redis - %s", err.Error())
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
