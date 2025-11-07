package sales

import (
	"context"

	"github.com/avraam311/sales-tracker/internal/models"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	CreateSale(context.Context, *models.SaleDTO) (uint, error)
	GetSales(context.Context) ([]*models.SaleDB, error)
	ReplaceSale(context.Context, uint, *models.SaleDTO) error
	DeleteSale(context.Context, uint) error
}

type Handler struct {
	service   Service
	validator *validator.Validate
}

func NewHandler(service Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}
