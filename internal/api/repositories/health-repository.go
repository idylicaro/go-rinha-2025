package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type HealthRepository struct {
	client *redis.Client
}

func NewHealthRepository(client *redis.Client) *HealthRepository {
	return &HealthRepository{client: client}
}

func (r *HealthRepository) Ping(ctx context.Context) bool {
	return r.client.Ping(ctx).Err() == nil
}
