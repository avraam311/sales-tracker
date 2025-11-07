package sales

import (
	"fmt"
	// "errors"
	"net/http"
	"strconv"

	"github.com/avraam311/sales-tracker/internal/api/handlers"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) DeleteSale(c *ginext.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to convert param id into int")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid id: %s", err.Error()))
		return
	}
	id := uint(idInt)

	err = h.service.DeleteSale(c.Request.Context(), id)
	if err != nil {
		// if errors.Is(err, sales.ErrSaleNotFound) {
		// 	zlog.Logger.Warn().Err(err).Msg("sale not found")
		// 	handlers.Fail(c.Writer, http.StatusNotFound, fmt.Errorf("sale not found"))
		// 	return
		// }

		zlog.Logger.Error().Err(err).Msg("failed to delete sale")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.OK(c.Writer, "sale deleted")
}
