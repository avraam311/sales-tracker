package sales

import (
	"context"
	"fmt"

	"github.com/avraam311/sales-tracker/internal/models"
)

func (r *Repository) CreateSale(ctx context.Context, sale *models.SaleDTO) (uint, error) {
	query := `
		INSERT INTO sale (item, income)
		VALUES ($1, $2)
		RETURNING id;
	`

	var id uint
	err := r.db.QueryRowContext(ctx, query, sale.Item, sale.Income).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("repository/create_sale.go - failed to create sale")
	}

	return id, nil
}
