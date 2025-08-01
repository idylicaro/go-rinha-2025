package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/idylicaro/go-rinha-2025/pkg/models"
)

type PaymentRepository struct {
	client *redis.Client
}

func NewPaymentRepository(client *redis.Client) *PaymentRepository {
	return &PaymentRepository{client: client}
}

func (r *PaymentRepository) SaveAudit(ctx context.Context, audit models.PaymentAudit) error {
	currentAmount, err := r.client.Get(ctx, audit.Endpoint).Float64()
	if currentAmount == 0 && err == redis.Nil {
		r.client.Set(ctx, audit.Endpoint, audit.Amount, 0)
	} else {
		newAmount := currentAmount + audit.Amount
		r.client.Set(ctx, audit.Endpoint, newAmount, 0)
	}

	countKey := audit.Endpoint + ":count"
	_, err = r.client.Incr(ctx, countKey).Result()
	return err
}

func (r *PaymentRepository) GetSummary(ctx context.Context, endpoint string) (*models.PaymentEndpointSummary, error) {
	total, err := r.client.Get(ctx, endpoint).Float64()
	if err == redis.Nil {
		total = 0
	} else if err != nil {
		return nil, err
	}

	count, err := r.client.Get(ctx, endpoint+":count").Int64()
	if err == redis.Nil {
		count = 0
	} else if err != nil {
		return nil, err
	}

	return &models.PaymentEndpointSummary{
		TotalRequests: count,
		TotalAmount:   total,
	}, nil
}
