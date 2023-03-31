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

type BillApi struct {
	BillService *service.BillService
}

func NewBillApi(service *service.BillService) *BillApi {
	return &BillApi{
		BillService: service,
	}
}

func (a *BillApi) CreateBillApi(e echo.Context) error {
	if e.Request().Header.Get("Content-Type") != "application/json" {
		storage.ConnectMongoSaveLog("JSON data is required", e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "JSON data is required"})
	}

	var bill model.Invoice
	if err := e.Bind(&bill); err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	billService, err := a.BillService.CreateBillService(e, &bill)
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return e.JSON(http.StatusOK, billService)
}

func (a *BillApi) UpdateBillApi(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "ID value is not correct !"})
	}

	var bill model.Invoice
	if err := e.Bind(&bill); err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	billService, err := a.BillService.UpdateBillService(e, strconv.Itoa(id), &bill)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			storage.ConnectMongoSaveLog(err.Error(), e)
			return e.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return e.JSON(http.StatusOK, billService)
}
func (a *BillApi) DeleteBillApi(e echo.Context) error {
	id := e.Param("id")

	bill, err := a.BillService.DeleteBillService(e, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			storage.ConnectMongoSaveLog(err.Error(), e)
			return e.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
	return e.JSON(http.StatusOK, bill)
}
func (a *BillApi) GetAllBillsApi(e echo.Context) error {
	product, err := a.BillService.GetAllBillsService(e)
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusOK, map[string]string{"message": err.Error()})
	}
	return e.JSON(http.StatusOK, product)
}

func (a *BillApi) GetBillByIdApi(e echo.Context) error {
	id := e.Param("id")

	product, err := a.BillService.GetBillByIdService(e, id)
	if err != nil {
		storage.ConnectMongoSaveLog(err.Error(), e)
		return e.JSON(http.StatusOK, map[string]string{"message": err.Error()})
	}
	return e.JSON(http.StatusOK, product)
}
