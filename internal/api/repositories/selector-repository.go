package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type SelectorRepository struct {
	client *redis.Client
}

func NewSelectorRepository(client *redis.Client) *SelectorRepository {
	return &SelectorRepository{client: client}
}

func (r *SelectorRepository) GetBestEndpoint(ctx context.Context, key string) (string, error) {
	best, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return best, nil
}

func (r *SelectorRepository) SetBestEndpoint(ctx context.Context, key string, endpoint string) error {
	return r.client.Set(ctx, key, endpoint, 0).Err()
}
