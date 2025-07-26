package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"github.com/idylicaro/go-rinha-2025/internal/api/controllers"
	"github.com/idylicaro/go-rinha-2025/internal/api/repositories"
	"github.com/idylicaro/go-rinha-2025/internal/api/services"
)

func NewApiRouter(redisClient *redis.Client) *chi.Mux {
	api := chi.NewRouter()

	// Define routes
	api.Post("/payments", controllers.CreatePayment)
	api.Get("/payments-summary", controllers.GetPaymentsSummary)
	api.Get("/health", controllers.NewHealthController(services.NewHealthService(repositories.NewHealthRepository(redisClient))).HealthCheck)
	return api
}
