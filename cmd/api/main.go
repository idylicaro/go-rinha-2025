package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/idylicaro/go-rinha-2025/internal/api"
	"github.com/idylicaro/go-rinha-2025/pkg/redis"
)

func main() {
	// Initialize the API router
	redisClient := redis.NewRedisClient()
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	}))
	r.Mount("/api", api.NewApiRouter(redisClient))

	// Start the HTTP server
	http.ListenAndServe(":8080", r)
}
