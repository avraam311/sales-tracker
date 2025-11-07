package analytics

import (
	"context"
	"fmt"
	"time"

	"github.com/avraam311/sales-tracker/internal/models"
)

func (s *Service) GetAnalytics(ctx context.Context, from time.Time, to time.Time) (*models.AnalyticsDB, error) {
	analytics, err := s.repo.GetAnalytics(ctx, from, to)
	if err != nil {
		return analytics, fmt.Errorf("service/get_analytics.go - %w", err)
	}

	return analytics, nil
}
