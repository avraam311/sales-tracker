package analytics

import (
	"context"

	"github.com/avraam311/sales-tracker/internal/models"
)

type Repository interface {
	GetAnalytics(context.Context, string, string) (*models.AnalyticsDB, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
