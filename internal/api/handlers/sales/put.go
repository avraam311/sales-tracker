package sales

import (
	"encoding/json"
	// "errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/avraam311/sales-tracker/internal/api/handlers"
	"github.com/avraam311/sales-tracker/internal/models"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) PutSale(c *ginext.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to convert param id into int")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid id: %s", err.Error()))
		return
	}
	id := uint(idInt)

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

	err = h.service.ReplaceSale(c.Request.Context(), id, &sale)
	if err != nil {
		// if errors.Is(err, sales.ErrSaleNotFound) {
		// 	zlog.Logger.Warn().Err(err).Msg("sale not found")
		// 	handlers.Fail(c.Writer, http.StatusNotFound, fmt.Errorf("sale not found"))
		// 	return
		// }

		zlog.Logger.Error().Err(err).Interface("sale", sale).Msg("failed to replace sale")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.Created(c.Writer, "sale replaced")
}
