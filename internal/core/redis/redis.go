package redis

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jesseinvent/go-payment-processor/internal/configs"
)


var redisClient = redis.NewClient(&redis.Options{
	Addr: configs.LoadConfigs().REDIS_URL,
})

func Set(ctx context.Context, key string, value string, ttl time.Duration) (bool, error) {

	err := redisClient.Set(ctx, key, value, ttl).Err();

	if err != nil {
		log.Println("Error setting value in redis", err);
		return false, err;
	}

	return true, nil;
}

func Get(ctx context.Context, key string) (string, error) {

	value, err := redisClient.Get(ctx, key).Result();

	if err == redis.Nil {
		return "", errors.New("Key does not exist")
	}

	return value, nil
 }

func Delete(ctx context.Context, key string) (bool, error) {

	err := redisClient.Del(ctx, key).Err();

	if err != nil {
		log.Println("Error deleting in redis", err)
		return false, err
	}

	return true, nil
}
