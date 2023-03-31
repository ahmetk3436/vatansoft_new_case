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

type CategoryApi struct {
	CategoryService *service.CategoryService
}

func NewCategoryApi(service *service.CategoryService) *CategoryApi {
	return &CategoryApi{
		CategoryService: service,
	}
}

func (a *CategoryApi) CreateCategoryApi(e echo.Context) error {
	if e.Request().Header.Get("Content-Type") != "application/json" {
		storage.ConnectMongoSaveLog("JSON data is required", e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "JSON data is required"})
	}

	var category model.Category
	if err := e.Bind(&category); err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	product, err := a.CategoryService.CreateCategoryService(e, &category)
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return e.JSON(http.StatusOK, product)
}

func (a *CategoryApi) UpdateCategoryApi(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "ID value is not correct !"})
	}

	var category model.Category
	if err := e.Bind(&category); err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	product, err := a.CategoryService.UpdateCategoryService(e, strconv.Itoa(id), &category)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			storage.ConnectMongoSaveLog(err.Error(), e)
			return e.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return e.JSON(http.StatusOK, product)
}
func (a *CategoryApi) DeleteCategoryApi(e echo.Context) error {
	id := e.Param("id")

	product, err := a.CategoryService.DeleteCategoryService(e, id)
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
func (a *CategoryApi) GetAllCategoriesApi(e echo.Context) error {
	product, err := a.CategoryService.GetAllCategoriesService(e)
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusOK, map[string]string{"message": err.Error()})
	}
	return e.JSON(http.StatusOK, product)
}

func (a *CategoryApi) GetCategoryByIdApi(e echo.Context) error {
	id := e.Param("id")

	product, err := a.CategoryService.GetCategoryByIdService(e, id)
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusOK, map[string]string{"message": err.Error()})
	}
	return e.JSON(http.StatusOK, product)
}
