package sales

import (
	"context"
	"fmt"

	"github.com/avraam311/sales-tracker/internal/models"
)

func (s *Service) ReplaceSale(ctx context.Context, id uint, sale *models.SaleDTO) error {
	err := s.repo.ReplaceSale(ctx, id, sale)
	if err != nil {
		return fmt.Errorf("service/replace_sale.go - %w", err)
	}

	return nil
}
