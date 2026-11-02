package service

import (
	"context"

	"study-smart-backend/internal/repository"
)

// HealthService contains business logic for health checks.
type HealthService struct {
	repo *repository.HealthRepository
}

// NewHealthService creates a new HealthService.
func NewHealthService(repo *repository.HealthRepository) *HealthService {
	return &HealthService{repo: repo}
}

// Check verifies the health of all dependencies.
func (s *HealthService) Check(ctx context.Context) (map[string]string, error) {
	status := map[string]string{
		"api": "up",
	}

	if err := s.repo.Ping(ctx); err != nil {
		status["database"] = "down"
		return status, err
	}

	status["database"] = "up"
	return status, nil
}
