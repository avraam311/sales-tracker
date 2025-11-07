package analytics

import (
	"context"
	"fmt"

	"github.com/avraam311/sales-tracker/internal/models"
)

func (s *Service) GetAnalytics(ctx context.Context, from string, to string) (*models.AnalyticsDB, error) {
	analytics, err := s.repo.GetAnalytics(ctx, from, to)
	if err != nil {
		return analytics, fmt.Errorf("service/get_analytics.go - %w", err)
	}

	return analytics, nil
}
