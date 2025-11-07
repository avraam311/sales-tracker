package analytics

import (
	"context"
	"time"

	"github.com/avraam311/sales-tracker/internal/models"
)

type Repository interface {
	GetAnalytics(context.Context, time.Time, time.Time) (*models.AnalyticsDB, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
