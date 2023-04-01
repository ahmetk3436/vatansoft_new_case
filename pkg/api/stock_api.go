package api

import (
	"errors"
	"net/http"
	"strconv"
	"vatansoft/internal/storage"
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
		storage.ConnectMongoSaveLog("JSON data is required", e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "JSON data is required"})
	}

	var productDTO model.ProductDTO
	if err := e.Bind(&productDTO); err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error() + "hata istekte"})
	}

	product, err := a.stockService.CreateStockProductService(e, &productDTO)
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error() + "hata repoda"})
	}

	return e.JSON(http.StatusOK, product)
}

func (a *Api) UpdateStockProductApi(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "ID value is not correct !"})
	}

	var productDTO model.ProductDTO
	if err := e.Bind(&productDTO); err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	product, err := a.stockService.UpdateStockProductService(e, strconv.Itoa(id), &productDTO)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			storage.ConnectMongoSaveLog(err.Error(), e)
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
			storage.ConnectMongoSaveLog(err.Error(), e)
			return e.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
	return e.JSON(http.StatusOK, product)
}

func (a *Api) FilterSearchStockProductApi(e echo.Context) error {
	q := e.QueryParam("query")
	category := e.QueryParam("category")
	minPrice := e.QueryParam("minPrice")
	maxPrice := e.QueryParam("maxPrice")
	isSold := e.QueryParam("sold")
	isDeleted := e.QueryParam("deleted")
	product, err := a.stockService.FilterSearchStockProductService(e, q, category, minPrice, maxPrice, isSold, isDeleted)
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusOK, map[string]string{"message": err.Error()})
	}
	return e.JSON(http.StatusOK, product)
}

func (a *Api) GetAllProductApi(e echo.Context) error {
	product, err := a.stockService.GetAllStockProductsService(e)
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusOK, map[string]string{"message": err.Error()})
	}
	return e.JSON(http.StatusOK, product)
}

func (a *Api) GetStockProductByIdApi(e echo.Context) error {
	id := e.Param("id")

	product, err := a.stockService.GetStockProductByIdService(e, id)
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusOK, map[string]string{"message": err.Error()})
	}
	return e.JSON(http.StatusOK, product)
}

func (a *Api) InsertCategoryForAllProductApi(e echo.Context) error {
	const contentType = "application/json"
	if e.Request().Header.Get("Content-Type") != contentType {
		msg := "JSON data is required"
		storage.ConnectMongoSaveLog(msg, e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": msg})
	}

	category := new(model.Category)
	if err := e.Bind(category); err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	product, err := a.stockService.InsertCategoryForAllProductService(e, category)
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusOK, map[string]string{"message": err.Error()})
	}

	return e.JSON(http.StatusOK, product)
}

func (a *Api) DeleteCategoryForProductByIdApi(e echo.Context) error {
	id := e.Param("id")
	categoryId := e.Param("categoryId")
	product, err := a.stockService.DeleteCategoryForProductByIdService(e, id, categoryId)
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusOK, map[string]string{"message": err.Error()})
	}
	return e.JSON(http.StatusOK, product)
}
func (a *Api) DeleteCategoryForProductsByIdApi(e echo.Context) error {
	id := e.Param("id")

	product, err := a.stockService.DeleteCategoryForProductsByIdService(e, id)
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusOK, err.Error())
	}
	return e.JSON(http.StatusOK, product)
}
