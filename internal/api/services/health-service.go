package services

import (
	"context"

	"github.com/idylicaro/go-rinha-2025/internal/api/repositories"
)

type HealthService struct {
	repo *repositories.HealthRepository
}

func NewHealthService(repo *repositories.HealthRepository) *HealthService {
	return &HealthService{repo: repo}
}

func (s *HealthService) CheckRedis(ctx context.Context) bool {
	return s.repo.Ping(ctx)
}
