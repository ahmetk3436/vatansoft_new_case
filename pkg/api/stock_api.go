package api

import (
	"errors"
	"net/http"
	"strconv"
	"vatansoft/internal/db"
	"vatansoft/pkg/model"
	"vatansoft/pkg/service"

	"gorm.io/gorm"

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
	if e.Request().Header.Get("Content-Type") != "application/json" {
		db.ConnectMongoSaveLog("JSON data is required", e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "JSON data is required"})
	}

	var productDTO model.ProductDTO
	if err := e.Bind(&productDTO); err != nil {
		db.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	product, err := a.stockService.CreateStockProductService(e, &productDTO)
	if err != nil {
		db.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return e.JSON(http.StatusOK, product)
}

func (a *Api) UpdateStockProductApi(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		db.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "ID value is not correct !"})
	}

	var productDTO model.ProductDTO
	if err := e.Bind(&productDTO); err != nil {
		db.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	product, err := a.stockService.UpdateStockProductService(e, strconv.Itoa(id), &productDTO)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.ConnectMongoSaveLog(err.Error(), e)
			return e.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return e.JSON(http.StatusOK, product)
}

func (a *Api) DeleteStockProductApi(e echo.Context) error {
	id := e.Param("id")

	product, err := a.stockService.DeleteStockProductService(e, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.ConnectMongoSaveLog(err.Error(), e)
			return e.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		db.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
	return e.JSON(http.StatusOK, product)
}

func (a *Api) FilterSearchStockProductApi(e echo.Context) error {
	q := e.QueryParam("query")
	category := e.QueryParam("category")
	sortBy := e.QueryParam("minPrice")
	sortOrder := e.QueryParam("maxPrice")

	product, err := a.stockService.FilterSearchStockProductService(e, q, category, sortBy, sortOrder)
	if err != nil {
		db.ConnectMongoSaveLog(err.Error(), e)
		return err
	}
	return e.JSON(http.StatusOK, product)
}

func (a *Api) GetAllProductApi(e echo.Context) error {
	product, err := a.stockService.GetAllStockProductsService(e)
	if err != nil {
		db.ConnectMongoSaveLog(err.Error(), e)
		return err
	}
	return e.JSON(http.StatusOK, product)
}

func (a *Api) GetStockProductByIdApi(e echo.Context) error {
	id := e.Param("id")

	product, err := a.stockService.GetStockProductByIdService(e, id)
	if err != nil {
		db.ConnectMongoSaveLog(err.Error(), e)
		return err
	}
	return e.JSON(http.StatusOK, product)
}
