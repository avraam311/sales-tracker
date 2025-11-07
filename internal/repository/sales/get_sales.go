package sales

import (
	"context"
	"fmt"

	"github.com/avraam311/sales-tracker/internal/models"
)

func (r *Repository) GetSales(ctx context.Context) ([]*models.SaleDB, error) {
	query := `
		SELECT id, item, income
		FROM sale
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository/get_sales.go - failed to get sales - %w", err)
	}
	defer rows.Close()

	var sales []*models.SaleDB
	for rows.Next() {
		var s models.SaleDB
		err := rows.Scan(&s.ID, &s.Item, &s.Income)
		if err != nil {
			return nil, fmt.Errorf("repository/get_sales.go - failed to scan sale row - %w", err)
		}
		sales = append(sales, &s)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("repository/get_sales.go - failed to iterate sales rows - %w", err)
	}

	return sales, nil
}
