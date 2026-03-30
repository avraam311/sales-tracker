package sales

import (
	"fmt"
	"net/http"

	"github.com/avraam311/sales-tracker/internal/api/handlers"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) CreateSale(c *ginext.Context) {
	sale, err := h.decodeAndValidateSale(c)
	if err != nil {
		handlers.Fail(c.Writer, http.StatusBadRequest, err)
		return
	}

	id, err := h.service.CreateSale(c.Request.Context(), sale)
	if err != nil {
		zlog.Logger.Error().Err(err).Interface("sale", sale).Msg("failed to create sale")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.Created(c.Writer, id)
}
