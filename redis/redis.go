package tools_redis

import (
	"context"
	"encoding/json"
	"github.com/nilchaosky/go-tools"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClient struct {
	client *redis.Client
}

type Option redis.Options

func NewRedisClient(option *Option) *RedisClient {
	rdb := redis.NewClient((*redis.Options)(option))
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		return &RedisClient{client: nil}
	}
	return &RedisClient{client: rdb}
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}) error {
	if r.client == nil {
		return &NotClientError{}
	}
	if tools.IsStruct(value) {
		jsonBytes, err := json.Marshal(value)
		if err != nil {
			return err
		}
		value = string(jsonBytes)
	}

	err := r.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) SetEx(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	if r.client == nil {
		return &NotClientError{}
	}
	if tools.IsStruct(value) {
		jsonBytes, err := json.Marshal(value)
		if err != nil {
			return err
		}
		value = string(jsonBytes)
	}

	err := r.client.Set(ctx, key, value, exp).Err()
	if err != nil {
		return err
	}
	return nil
}
