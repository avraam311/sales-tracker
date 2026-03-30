package analytics

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/avraam311/sales-tracker/internal/models"
)

func (r *Repository) GetAnalytics(ctx context.Context, from, to time.Time, groupBy string) ([]*models.AnalyticsDB, error) {
	if from.IsZero() {
		from = time.Unix(0, 0)
	}
	if to.IsZero() {
		to = time.Now()
	}

	var query string
	var args []interface{}
	var scanFunc func(rows *sql.Rows) ([]*models.AnalyticsDB, error)

	switch groupBy {
	case "day":
		query = `
			SELECT
				DATE(created_at) as period,
				COALESCE(SUM(income), 0) as sum,
				COALESCE(AVG(income), 0) as avg,
				COUNT(*) as count,
				COALESCE(PERCENTILE_CONT(0.5) WITHIN GROUP (ORDER BY income), 0) AS median,
				COALESCE(PERCENTILE_CONT(0.9) WITHIN GROUP (ORDER BY income), 0) AS percentile_90
			FROM sale
			WHERE created_at >= $1 AND created_at <= $2
			GROUP BY DATE(created_at)
			ORDER BY period;
		`
		args = []interface{}{from, to}
		scanFunc = func(rows *sql.Rows) ([]*models.AnalyticsDB, error) {
			defer rows.Close()
			var analytics []*models.AnalyticsDB
			for rows.Next() {
				var a models.AnalyticsDB
				if err := rows.Scan(&a.Sum, &a.Avg, &a.Count, &a.Median, &a.Percentile90); err != nil {
					return nil, fmt.Errorf("scan analytics day: %w", err)
				}
				analytics = append(analytics, &a)
			}
			return analytics, rows.Err()
		}
	case "week":
		query = `
			SELECT
				DATE_TRUNC('week', created_at) as period,
				COALESCE(SUM(income), 0) as sum,
				COALESCE(AVG(income), 0) as avg,
				COUNT(*) as count,
				COALESCE(PERCENTILE_CONT(0.5) WITHIN GROUP (ORDER BY income), 0) AS median,
				COALESCE(PERCENTILE_CONT(0.9) WITHIN GROUP (ORDER BY income), 0) AS percentile_90
			FROM sale
			WHERE created_at >= $1 AND created_at <= $2
			GROUP BY DATE_TRUNC('week', created_at)
			ORDER BY period;
		`
		args = []interface{}{from, to}
		scanFunc = func(rows *sql.Rows) ([]*models.AnalyticsDB, error) {
			defer rows.Close()
			var analytics []*models.AnalyticsDB
			for rows.Next() {
				var a models.AnalyticsDB
				if err := rows.Scan(&a.Sum, &a.Avg, &a.Count, &a.Median, &a.Percentile90); err != nil {
					return nil, fmt.Errorf("scan analytics week: %w", err)
				}
				analytics = append(analytics, &a)
			}
			return analytics, rows.Err()
		}
	case "category":
		query = `
			SELECT
				item as period,
				COALESCE(SUM(income), 0) as sum,
				COALESCE(AVG(income), 0) as avg,
				COUNT(*) as count,
				COALESCE(PERCENTILE_CONT(0.5) WITHIN GROUP (ORDER BY income), 0) AS median,
				COALESCE(PERCENTILE_CONT(0.9) WITHIN GROUP (ORDER BY income), 0) AS percentile_90
			FROM sale
			WHERE created_at >= $1 AND created_at <= $2
			GROUP BY item
			ORDER BY sum DESC;
		`
		args = []interface{}{from, to}
		scanFunc = func(rows *sql.Rows) ([]*models.AnalyticsDB, error) {
			defer rows.Close()
			var analytics []*models.AnalyticsDB
			for rows.Next() {
				var a models.AnalyticsDB
				if err := rows.Scan(&a.Sum, &a.Avg, &a.Count, &a.Median, &a.Percentile90); err != nil {
					return nil, fmt.Errorf("scan analytics category: %w", err)
				}
				analytics = append(analytics, &a)
			}
			return analytics, rows.Err()
		}
	default:
		query = `
			SELECT
				COALESCE(SUM(income), 0) as sum,
				COALESCE(AVG(income), 0) as avg,
				COUNT(*) as count,
				COALESCE(PERCENTILE_CONT(0.5) WITHIN GROUP (ORDER BY income), 0) AS median,
				COALESCE(PERCENTILE_CONT(0.9) WITHIN GROUP (ORDER BY income), 0) AS percentile_90
			FROM sale
			WHERE created_at >= $1 AND created_at <= $2;
		`
		args = []interface{}{from, to}
		rows, err := r.db.QueryContext(ctx, query, args...)
		if err != nil {
			return nil, fmt.Errorf("repository/get_analytics.go - failed to get analytics - %w", err)
		}
		defer rows.Close()
		var analytics models.AnalyticsDB
		if !rows.Next() {
			return []*models.AnalyticsDB{{}}, nil
		}
		err = rows.Scan(&analytics.Sum, &analytics.Avg, &analytics.Count, &analytics.Median, &analytics.Percentile90)
		if err != nil {
			return nil, fmt.Errorf("repository/get_analytics.go - failed to scan analytics - %w", err)
		}
		return []*models.AnalyticsDB{&analytics}, nil
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("repository/get_analytics.go - failed to get analytics - %w", err)
	}
	defer rows.Close()

	return scanFunc(rows)
}
