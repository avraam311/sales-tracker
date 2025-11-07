package sales

import (
	"context"
	"fmt"
)

func (s *Service) DeleteSale(ctx context.Context, id uint) error {
	err := s.repo.DeleteSale(ctx, id)
	if err != nil {
		return fmt.Errorf("service/delete_sale.go - %w", err)
	}

	return nil
}
