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

func (h *Handler) DeleteSale(c *ginext.Context) {
	id, err := h.parseID(c)
	if err != nil {
		handlers.Fail(c.Writer, http.StatusBadRequest, err)
		return
	}

	err = h.service.DeleteSale(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, sales.ErrSaleNotFound) {
			zlog.Logger.Warn().Err(err).Msg("sale not found")
			handlers.Fail(c.Writer, http.StatusNotFound, fmt.Errorf("sale not found"))
			return
		}

		zlog.Logger.Error().Err(err).Msg("failed to delete sale")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.OK(c.Writer, "sale deleted")
}
