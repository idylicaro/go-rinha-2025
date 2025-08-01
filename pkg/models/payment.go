package models

type PaymentRequest struct {
	CorrelationID string  `json:"correlationId"`
	Amount        float64 `json:"amount"`
	RequestedAt   *string `json:"requestedAt,omitempty"`
}

type PaymentAudit struct {
	Endpoint string  `json:"endpoint"`
	Amount   float64 `json:"amount"`
}

type PaymentSummary struct {
	Default  PaymentEndpointSummary `json:"default"`
	Fallback PaymentEndpointSummary `json:"fallback"`
}

type PaymentEndpointSummary struct {
	TotalRequests int64   `json:"totalRequests"`
	TotalAmount   float64 `json:"totalAmount"`
}
