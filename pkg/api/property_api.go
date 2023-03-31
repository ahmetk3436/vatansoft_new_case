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

type PropertyApi struct {
	PropertyService *service.PropertyService
}

func NewPropertyApi(service *service.PropertyService) *PropertyApi {
	return &PropertyApi{
		PropertyService: service,
	}
}

func (a *PropertyApi) CreatePropertyApi(e echo.Context) error {
	if e.Request().Header.Get("Content-Type") != "application/json" {
		storage.ConnectMongoSaveLog("JSON data is required", e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "JSON data is required"})
	}

	var Property model.ProductProperty
	if err := e.Bind(&Property); err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	propertyService, err := a.PropertyService.CreatePropertyService(e, &Property)
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return e.JSON(http.StatusOK, propertyService)
}

func (a *PropertyApi) UpdatePropertyApi(e echo.Context) error {
	if e.Request().Header.Get("Content-Type") != "application/json" {
		storage.ConnectMongoSaveLog("JSON data is required", e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "JSON data is required"})
	}
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "ID value is not correct !"})
	}

	var property model.ProductProperty
	if err := e.Bind(&property); err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	propertyService, err := a.PropertyService.UpdatePropertyService(e, strconv.Itoa(id), &property)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			storage.ConnectMongoSaveLog(err.Error(), e)
			return e.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return e.JSON(http.StatusOK, propertyService)
}
func (a *PropertyApi) DeletePropertyApi(e echo.Context) error {

	id := e.Param("id")

	property, err := a.PropertyService.DeletePropertyService(e, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			storage.ConnectMongoSaveLog(err.Error(), e)
			return e.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
	return e.JSON(http.StatusOK, property)
}
func (a *PropertyApi) GetAllPropertysApi(e echo.Context) error {
	product, err := a.PropertyService.GetAllPropertysService(e)
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusOK, map[string]string{"message": err.Error()})
	}
	return e.JSON(http.StatusOK, product)
}

func (a *PropertyApi) GetPropertyByIdApi(e echo.Context) error {
	id := e.Param("id")

	property, err := a.PropertyService.GetPropertyByIdService(e, id)
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusOK, map[string]string{"message": err.Error()})
	}
	return e.JSON(http.StatusOK, property)
}
