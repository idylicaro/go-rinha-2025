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
	// TODO: Create Factories for Controllers and Services
	// TODO: Use environment variables for endpoints
	createPaymentController := controllers.NewCreatePaymentController(
		services.NewPaymentService(repositories.NewPaymentRepository(redisClient)),
		services.NewEndpointSelectorService(
			[]string{"http://payment-processor-default:8080", "http://payment-processor-fallback:8080"},
			repositories.NewSelectorRepository(redisClient),
			"best_endpoint",
			5,
		),
	)
	api.Post("/payments", createPaymentController.CreatePayment)
	api.Get("/payments-summary", controllers.GetPaymentsSummary)
	api.Get("/health", controllers.NewHealthController(services.NewHealthService(repositories.NewHealthRepository(redisClient))).HealthCheck)
	return api
}
