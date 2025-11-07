package sales

import (
	"context"
	"fmt"
)

func (r *Repository) DeleteSale(ctx context.Context, id uint) error {
	query := `
        DELETE FROM sale
        WHERE id = $1;
    `

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("repository/delete_sale.go - failed to delete sale: %w", err)
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return ErrSaleNotFound
	}

	return nil
}
