package cache

import (
	"account-transaction-api/internal/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Cache interface {
	SetNx(ctx context.Context, pattern, key string) bool
}

type Wrapper struct {
	client *redis.Client
	TTLs   map[string]int
}

func Init(cfg *config.Config) *Wrapper {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.CACHE.Host, cfg.CACHE.Port),
		Password: cfg.CACHE.Password,
		DB:       0, // use default DB
	})

	cache := &Wrapper{client: rdb, TTLs: map[string]int{
		TransactionLockKeyPattern: cfg.CACHE.TTL.Transactions},
	}

	return cache
}

func (c *Wrapper) SetNx(ctx context.Context, pattern, id string) bool {

	ttl := time.Duration(c.TTLs[pattern]) * time.Millisecond
	key := fmt.Sprintf(pattern, id)
	r := c.client.SetNX(ctx, key, key, ttl)

	return r.Val()
}
