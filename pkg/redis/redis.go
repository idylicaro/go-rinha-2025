package redis

import (
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient() *redis.Client {
	db := 0
	if dbStr := os.Getenv("REDIS_DB"); dbStr != "" {
		if val, err := strconv.Atoi(dbStr); err == nil {
			db = val
		}
	}

	poolSize := 3 // default value
	if poolSizeStr := os.Getenv("REDIS_POOL_SIZE"); poolSizeStr != "" {
		if val, err := strconv.Atoi(poolSizeStr); err == nil {
			poolSize = val
		}
	}

	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"), // ex: "redis:6379"
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,       // default is 0
		PoolSize: poolSize, // PoolSize = CPUs * 4 → 0.6 * 4 ≈ 2~3
	})
}
