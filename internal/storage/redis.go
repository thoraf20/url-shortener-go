package storage

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/thoraf20/url-shortner/internal/models"
)

type RedisStore struct  {
	client *redis.Client
}

func NewRedisStore(addr string) (*RedisStore, error) {
	client := redis.NewClient(&redis.Options{ Addr: addr })
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &RedisStore{ client: client}, nil
}

func (s *RedisStore) SaveURL(ctx context.Context, url *models.URL) error {
	data, err := json.Marshal(url)
	if err != nil {
		return err
	}
	key := url.TenantID + ":" + url.ShortCode
	return s.client.Set(ctx, key, data, 0).Err()
}

func (s *RedisStore) GetURL(ctx context.Context, tenantID, shortCode string) (*models.URL, error) {
	key := tenantID + ":" + shortCode
	data, err := s.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var url models.URL
	if err := json.Unmarshal(data, &url); err != nil {
		return nil, err
	}

	return &url, nil
}