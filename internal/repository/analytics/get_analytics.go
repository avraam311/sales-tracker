package analytics

import (
	"context"
	"fmt"
	"time"

	"github.com/avraam311/sales-tracker/internal/models"
)

func (r *Repository) GetAnalytics(ctx context.Context, from, to time.Time) (*models.AnalyticsDB, error) {
	query := `
        SELECT
            COALESCE(SUM(income), 0) as sum,
            COALESCE(AVG(income), 0) as avg,
            COUNT(*) as count,
            COALESCE(PERCENTILE_CONT(0.5) WITHIN GROUP (ORDER BY income), 0) AS median,
            COALESCE(PERCENTILE_CONT(0.9) WITHIN GROUP (ORDER BY income), 0) AS percentile_90
        FROM sale
        WHERE created_at >= $1 AND created_at <= $2;
    `

	var analytics models.AnalyticsDB
	err := r.db.QueryRowContext(ctx, query, from, to).Scan(
		&analytics.Sum,
		&analytics.Avg,
		&analytics.Count,
		&analytics.Median,
		&analytics.Percentile90,
	)
	if err != nil {
		return nil, fmt.Errorf("repository/get_analytics.go - failed to get analytics - %w", err)
	}

	return &analytics, nil
}
