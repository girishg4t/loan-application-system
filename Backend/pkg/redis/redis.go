package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Config represent redis configuration.
type Config struct {
	Database int
	Address  string
}

type RedisClient struct {
	Client *redis.Client
}

// NewClient creates new redis client.
func NewClient(ctx context.Context, cfg Config) (RedisClient, error) {
	var rdb = redis.NewClient(&redis.Options{
		Addr: cfg.Address,
		DB:   cfg.Database,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return RedisClient{}, fmt.Errorf("failed to create new redis client: %w", err)
	}

	return RedisClient{Client: rdb}, nil
}
