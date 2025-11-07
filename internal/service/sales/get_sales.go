package sales

import (
	"context"
	"fmt"

	"github.com/avraam311/sales-tracker/internal/models"
)

func (s *Service) GetSales(ctx context.Context) ([]*models.SaleDB, error) {
	sales, err := s.repo.GetSales(ctx)
	if err != nil {
		return sales, fmt.Errorf("service/get_sales.go - %w", err)
	}

	return sales, nil
}
