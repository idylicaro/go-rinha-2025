package services

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/idylicaro/go-rinha-2025/internal/api/repositories"
	"github.com/idylicaro/go-rinha-2025/pkg/models"
)

type PaymentService struct {
	repo *repositories.PaymentRepository
}

func NewPaymentService(repo *repositories.PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) CreatePayment(ctx context.Context, endpoint string, payment models.PaymentRequest) error {
	go func() {
		payload := map[string]interface{}{
			"correlationId": payment.CorrelationID,
			"amount":        payment.Amount,
		}

		var resp *http.Response
		var err error
		maxRetries := 3
		for i := 0; i < maxRetries; i++ {
			b, _ := json.Marshal(payload)
			resp, err = http.Post(endpoint+"/payments", "application/json", bytes.NewReader(b))
			if err == nil && resp.StatusCode == http.StatusOK {
				break
			}
			time.Sleep(time.Duration(100*(i+1)) * time.Millisecond) // exponential backoff
		}
		if err != nil || resp.StatusCode != http.StatusOK {
			// opcional: salvar falha para auditoria
			return
		}
		defer resp.Body.Close()

		audit := models.PaymentAudit{
			Endpoint: endpoint,
			Amount:   payment.Amount,
		}
		_ = s.repo.SaveAudit(ctx, audit)
	}()

	return nil
}
