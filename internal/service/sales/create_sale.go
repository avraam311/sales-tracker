package sales

import (
	"context"
	"fmt"

	"github.com/avraam311/sales-tracker/internal/models"
)

func (s *Service) CreateSale(ctx context.Context, sale *models.SaleDTO) (uint, error) {
	id, err := s.repo.CreateSale(ctx, sale)
	if err != nil {
		return 0, fmt.Errorf("service/create_sale.go - %w", err)
	}

	return id, nil
}
