package redisclient

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/mohamedkaram400/go-crud-ops/config"
)

var Client *redis.Client

func Init() error {
	Client = redis.NewClient(&redis.Options{
		Addr: config.GetRedisHost(),
	})

	// Test connection
	ctx := context.Background()
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %v", err)
	}

	fmt.Println("✅ Connected to Redis successfully")
	return nil
}