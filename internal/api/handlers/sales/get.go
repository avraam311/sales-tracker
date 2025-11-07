package sales

import (
	"fmt"
	"net/http"

	"github.com/avraam311/sales-tracker/internal/api/handlers"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) GetSales(c *ginext.Context) {
	sales, err := h.service.GetSales(c.Request.Context())
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to get sales")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.OK(c.Writer, sales)
}
