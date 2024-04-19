package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}
type StandardClientAdapter struct {
	client *redis.Client
}

func (a *StandardClientAdapter) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return a.client.Set(ctx, key, value, expiration).Err()
}

func (a *StandardClientAdapter) Get(ctx context.Context, key string) (string, error) {
	return a.client.Get(ctx, key).Result()
}

type ClusterClientAdapter struct {
	clusterClient *redis.ClusterClient
}

func (a *ClusterClientAdapter) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return a.clusterClient.Set(ctx, key, value, expiration).Err()
}

func (a *ClusterClientAdapter) Get(ctx context.Context, key string) (string, error) {
	return a.clusterClient.Get(ctx, key).Result()
}

func RedisClientFactory(redisConfig interface{}) (RedisClient, error) {
	switch client := redisConfig.(type) {
	case *redis.Client:
		return &StandardClientAdapter{client: client}, nil
	case *redis.ClusterClient:
		return &ClusterClientAdapter{clusterClient: client}, nil
	default:
		return nil, fmt.Errorf("unsupported redis client type")
	}
}

func main() {
	var client RedisClient
	// standardClient := redis.NewClient(&redis.Options{ /* 配置参数 */ })
	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{ /* 配置参数 */ })

	// 使用工厂方法创建适配器
	// client, _ = RedisClientFactory(standardClient)
	client, _ = RedisClientFactory(clusterClient)

	// 使用适配器进行操作
	client.Set(context.Background(), "key1", "value1", 0)
	value, _ := client.Get(context.Background(), "key1")
	fmt.Println("Got value:", value)

}
