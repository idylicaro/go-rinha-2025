package services

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/idylicaro/go-rinha-2025/internal/api/repositories"
)

type HealthResponse struct {
	Failing         bool `json:"failing"`
	MinResponseTime int  `json:"minResponseTime"`
}

type EndpointSelectorService struct {
	endpoints []string
	repo      *repositories.SelectorRepository
	redisKey  string
	interval  int
}

func NewEndpointSelectorService(endpoints []string, repo *repositories.SelectorRepository, redisKey string, interval int) *EndpointSelectorService {
	es := &EndpointSelectorService{
		endpoints: endpoints,
		repo:      repo,
		redisKey:  redisKey,
		interval:  interval,
	}
	go es.healthCheckLoop()
	return es
}

func (es *EndpointSelectorService) healthCheckLoop() {
	ctx := context.Background()
	for {
		var best string
		minTime := int(^uint(0) >> 1)
		for _, ep := range es.endpoints {
			resp, err := http.Get(ep + "/payments/service-health")
			if err != nil || resp.StatusCode != http.StatusOK {
				continue
			}
			var hr HealthResponse
			if err := json.NewDecoder(resp.Body).Decode(&hr); err != nil {
				resp.Body.Close()
				continue
			}
			resp.Body.Close()
			if !hr.Failing && hr.MinResponseTime < minTime {
				best = ep
				minTime = hr.MinResponseTime
			}
		}
		if best != "" {
			es.repo.SetBestEndpoint(ctx, es.redisKey, best)
		}
		time.Sleep(time.Duration(es.interval) * time.Second)
	}
}

func (es *EndpointSelectorService) GetBestEndpoint() string {
	ctx := context.Background()
	best, err := es.repo.GetBestEndpoint(ctx, es.redisKey)
	if err != nil {
		return ""
	}
	return best
}
