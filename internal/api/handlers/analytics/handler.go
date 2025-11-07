package analytics

import (
	"context"
	"time"

	"github.com/avraam311/sales-tracker/internal/models"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	GetAnalytics(context.Context, time.Time, time.Time) (*models.AnalyticsDB, error)
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
