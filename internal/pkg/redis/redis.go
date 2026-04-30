package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)
type RedisService interface {
	Set(ctx context.Context, key string, value string, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error 
}
type redisService struct {
	redisClient *redis.Client
}

func NewRedisService(redisUrl string) (RedisService, error) {

	redisClient, err := ConnectToRedisWithRetry(redisUrl, 5, 2*time.Second)

	if err != nil {
		return nil, err
	}

	return &redisService{redisClient: redisClient}, nil
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

func (r *redisService) Set(ctx context.Context, key string, value string, ttl time.Duration) error {

	err := r.redisClient.Set(ctx, key, value, ttl).Err();

	if err != nil {
		return fmt.Errorf("error setting value in redis - %w", err)
	}

	return nil;
}

func (r *redisService) Get(ctx context.Context, key string) (string, error) {

	value, err := r.redisClient.Get(ctx, key).Result();

	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", fmt.Errorf("error getting value from redis - %w", err)
	}

	return value, nil
 }

func (r *redisService) Delete(ctx context.Context, key string) (error) {

	err := r.redisClient.Del(ctx, key).Err();

	if err != nil {
		return fmt.Errorf("error deleting value from redis - %w", err)
	}

	return nil
}
