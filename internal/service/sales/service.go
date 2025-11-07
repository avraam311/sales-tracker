package sales

import (
	"context"

	"github.com/avraam311/sales-tracker/internal/models"
)

type Repository interface {
	CreateSale(context.Context, *models.SaleDTO) (uint, error)
	GetSales(context.Context) ([]*models.SaleDB, error)
	ReplaceSale(context.Context, uint, *models.SaleDTO) error
	DeleteSale(context.Context, uint) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
