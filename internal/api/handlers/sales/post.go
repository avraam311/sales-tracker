package sales

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/avraam311/sales-tracker/internal/api/handlers"
	"github.com/avraam311/sales-tracker/internal/models"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) CreateSale(c *ginext.Context) {
	var sale models.SaleDTO
	if err := json.NewDecoder(c.Request.Body).Decode(&sale); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to decode request body")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid request body: %s", err.Error()))
		return
	}

	if err := h.validator.Struct(sale); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to validate request body")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("validation error: %s", err.Error()))
		return
	}

	id, err := h.service.CreateSale(c.Request.Context(), &sale)
	if err != nil {
		zlog.Logger.Error().Err(err).Interface("sale", sale).Msg("failed to create sale")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.Created(c.Writer, id)
}
