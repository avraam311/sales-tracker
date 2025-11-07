package analytics

import (
	"fmt"
	"net/http"
	"time"

	"github.com/avraam311/sales-tracker/internal/api/handlers"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) GetAnalytics(c *ginext.Context) {
	fromStr := c.Query("from")
	toStr := c.Query("to")

	var fromTime, toTime time.Time
	var err error
	if fromStr != "" {
		fromTime, err = time.Parse(time.RFC3339, fromStr)
		if err != nil {
			zlog.Logger.Error().Err(err).Str("from", fromStr).Msg("invalid 'from' date format")
			handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid 'from' date format: %v", err))
			return
		}
	}
	if toStr != "" {
		toTime, err = time.Parse(time.RFC3339, toStr)
		if err != nil {
			zlog.Logger.Error().Err(err).Str("to", toStr).Msg("invalid 'to' date format")
			handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid 'to' date format: %v", err))
			return
		}
	}
	if !fromTime.IsZero() && !toTime.IsZero() && fromTime.After(toTime) {
		err := fmt.Errorf("'from' date must be before 'to' date")
		zlog.Logger.Error().Err(err).Str("from", fromStr).Str("to", toStr).Msg("invalid date range")
		handlers.Fail(c.Writer, http.StatusBadRequest, err)
		return
	}
	analytics, err := h.service.GetAnalytis(c.Request.Context(), fromTime, toTime)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to get analytics")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.OK(c.Writer, analytics)
}
