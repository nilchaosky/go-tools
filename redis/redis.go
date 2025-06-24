package tools_redis

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/nilchaosky/go-tools/tools"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClient struct {
	client *redis.Client
}

type Option redis.Options

func GetClient(option *Option) *RedisClient {
	rdb := redis.NewClient((*redis.Options)(option))
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		return &RedisClient{client: nil}
	}
	return &RedisClient{client: rdb}
}

// Set 设置 key
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

// SetEx 设置 key 和 过期时间
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

// Get 获取 key
func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	if r.client == nil {
		return "", &NotClientError{}
	}

	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", &NotFoundError{label: key}
		}
		return "", err
	}
	return value, nil
}

// SetNX 如果 key 不存在才设置
func (r *RedisClient) SetNX(ctx context.Context, key string, value interface{}) error {
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

	err := r.client.SetNX(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// SetNEX 如果 key 不存在才设置 和 过期时间
func (r *RedisClient) SetNEX(ctx context.Context, key string, value interface{}, exp time.Duration) error {
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

	err := r.client.SetNX(ctx, key, value, exp).Err()
	if err != nil {
		return err
	}
	return nil
}

// SetXX 如果 key 存在才设置
func (r *RedisClient) SetXX(ctx context.Context, key string, value interface{}) error {
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

	err := r.client.SetXX(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// SetXEX 	如果 key 存在才设置 和 过期时间
func (r *RedisClient) SetXEX(ctx context.Context, key string, value interface{}, exp time.Duration) error {
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

	err := r.client.SetXX(ctx, key, value, exp).Err()
	if err != nil {
		return err
	}
	return nil
}
