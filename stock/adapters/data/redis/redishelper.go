package redis

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type redisHelper struct {
	client *redis.Client
}

func (rh redisHelper) Ping(ctx context.Context) (bool) {
	val, err := rh.client.Ping(ctx).Result()
	if err != nil {
		return false
	}
	if val != "PONG" {
		return true
	}
	return false
}

func (rh redisHelper) GetIntValue(ctx context.Context, id string) (int, error) {
	val, err := rh.client.Get(ctx, id).Result()
	if err != nil {
		return 0, err
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return intVal, nil
}
