package redis

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	redisClient *redis.Client
}


func NewRedis(redisUrl string) (*Redis, error) {

 	redisClient := redis.NewClient(&redis.Options{
		Addr: redisUrl,
	})

	err := redisClient.Ping(context.Background()).Err()

	if err != nil {
		return nil, fmt.Errorf("Failed to connect to Redis: %w", err)
	}

	return &Redis{redisClient: redisClient}, nil
}


func (r *Redis) Set(ctx context.Context, key string, value string, ttl time.Duration) (bool, error) {

	err := r.redisClient.Set(ctx, key, value, ttl).Err();

	if err != nil {
		log.Println("Error setting value in redis", err);
		return false, err;
	}

	return true, nil;
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {

	value, err := r.redisClient.Get(ctx, key).Result();

	if err == redis.Nil {
		return "", errors.New("Key does not exist")
	}

	return value, nil
 }

func (r *Redis) Delete(ctx context.Context, key string) (bool, error) {

	err := r.redisClient.Del(ctx, key).Err();

	if err != nil {
		log.Println("Error deleting in redis", err)
		return false, err
	}

	return true, nil
}
