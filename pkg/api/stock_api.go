package api

import (
	"net/http"
	"vatansoft/pkg/service"

	"github.com/labstack/echo/v4"
)

type Api struct {
	stockService *service.StockService
}

func NewStockApi(service *service.StockService) *Api {
	return &Api{
		stockService: service,
	}
}

func (a *Api) CreateStockProductApi(e echo.Context) error {
	product, err := a.stockService.CreateStockProductService(e)
	if err != nil {
		return echo.ErrBadRequest
	}
	return e.JSON(http.StatusOK, product)
}
