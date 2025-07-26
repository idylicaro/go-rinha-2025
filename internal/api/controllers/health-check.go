package controllers

import (
	"net/http"

	"github.com/idylicaro/go-rinha-2025/internal/api/services"
)

type HealthController struct {
	service *services.HealthService
}

func NewHealthController(service *services.HealthService) *HealthController {
	return &HealthController{service: service}
}

func (hc *HealthController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if !hc.service.CheckRedis(r.Context()) {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("redis unavailable"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
