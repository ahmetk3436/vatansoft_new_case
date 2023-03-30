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

func (a *Api) UpdateStockProductApi(e echo.Context) error {
	product, err := a.stockService.UpdateStockProductService(e)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, product)
}

func (a *Api) DeleteStockProductApi(e echo.Context) error {
	product, err := a.stockService.DeleteStockProductService(e)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, product)
}

func (a *Api) FilterSearchStockProductApi(e echo.Context) error {
	product, err := a.stockService.FilterSearchStockProductService(e)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, product)
}

func (a *Api) GetAllProductApi(e echo.Context) error {
	product, err := a.stockService.GetAllStockProductsService(e)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, product)
}

func (a *Api) GetStockProductByIdApi(e echo.Context) error {
	product, err := a.stockService.GetStockProductByIdService(e)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, product)
}
