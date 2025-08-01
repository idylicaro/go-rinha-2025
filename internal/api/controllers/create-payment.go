package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/idylicaro/go-rinha-2025/internal/api/services"
	"github.com/idylicaro/go-rinha-2025/pkg/models"
)

type CreatePaymentController struct {
	service  *services.PaymentService
	selector *services.EndpointSelectorService
}

func NewCreatePaymentController(service *services.PaymentService, selector *services.EndpointSelectorService) *CreatePaymentController {
	return &CreatePaymentController{service: service, selector: selector}
}

func (c *CreatePaymentController) CreatePayment(w http.ResponseWriter, r *http.Request) {
	// Logic for creating a payment will go here
	var body models.PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	requestedAt := time.Now().UTC().Format(time.RFC3339)
	body.RequestedAt = &requestedAt

	bestEndpoint := c.selector.GetBestEndpoint()
	c.service.CreatePayment(r.Context(), bestEndpoint, body)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Payment created successfully"))
}
