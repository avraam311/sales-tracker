package sales

import (
	"context"
	"fmt"

	"github.com/avraam311/sales-tracker/internal/models"
)

func (r *Repository) ReplaceSale(ctx context.Context, id uint, sale *models.SaleDTO) error {
	query := `
        UPDATE sale
        SET item = $1, income = $2
        WHERE id = $3;
    `

	res, err := r.db.ExecContext(ctx, query, sale.Item, sale.Income, id)
	if err != nil {
		return fmt.Errorf("repository/replace_sale.go - failed to update sale: %w", err)
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return ErrSaleNotFound
	}

	return nil
}
