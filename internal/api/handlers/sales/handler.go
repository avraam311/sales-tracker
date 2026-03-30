package sales

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/avraam311/sales-tracker/internal/models"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"

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

func (h *Handler) parseID(c *ginext.Context) (uint, error) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to convert param id into int")
		return 0, fmt.Errorf("invalid id: %s", err.Error())
	}
	return uint(idInt), nil
}

func (h *Handler) decodeAndValidateSale(c *ginext.Context) (*models.SaleDTO, error) {
	var sale models.SaleDTO
	if err := json.NewDecoder(c.Request.Body).Decode(&sale); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to decode request body")
		return nil, fmt.Errorf("invalid request body: %s", err.Error())
	}
	if err := h.validator.Struct(sale); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to validate request body")
		return nil, fmt.Errorf("validation error: %s", err.Error())
	}
	return &sale, nil
}

func NewHandler(service Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}
