package server

import (
	"net/http"

	"github.com/avraam311/sales-tracker/internal/api/handlers/analytics"
	"github.com/avraam311/sales-tracker/internal/api/handlers/sales"
	"github.com/avraam311/sales-tracker/internal/api/middlewares"

	"github.com/wb-go/wbf/ginext"
)

func NewRouter(ginMode string, handlerSale *sales.Handler, handlerAn *analytics.Handler) *ginext.Engine {
	e := ginext.New(ginMode)

	e.Use(middlewares.CORSMiddleware())
	e.Use(ginext.Logger())
	e.Use(ginext.Recovery())

	api := e.Group("/sales-tracker/api")
	{
		api.POST("/items", handlerSale.CreateSale)
		api.GET("/items", handlerSale.GetSales)
		api.PUT("/items/:id", handlerSale.PutSale)
		api.DELETE("/items/:id", handlerSale.DeleteSale)

		api.GET("/analytics", handlerAn.GetAnalytics)
	}

	return e
}

func NewServer(addr string, router *ginext.Engine) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
