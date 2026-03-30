package sales

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/avraam311/sales-tracker/internal/api/handlers"
	"github.com/avraam311/sales-tracker/internal/repository/sales"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) PutSale(c *ginext.Context) {
	id, err := h.parseID(c)
	if err != nil {
		handlers.Fail(c.Writer, http.StatusBadRequest, err)
		return
	}

	sale, err := h.decodeAndValidateSale(c)
	if err != nil {
		handlers.Fail(c.Writer, http.StatusBadRequest, err)
		return
	}

	err = h.service.ReplaceSale(c.Request.Context(), id, sale)
	if err != nil {
		if errors.Is(err, sales.ErrSaleNotFound) {
			zlog.Logger.Warn().Err(err).Msg("sale not found")
			handlers.Fail(c.Writer, http.StatusNotFound, fmt.Errorf("sale not found"))
			return
		}

		zlog.Logger.Error().Err(err).Interface("sale", sale).Msg("failed to replace sale")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.Created(c.Writer, "sale replaced")
}
