package handler

import (
	"fmt"
	"vatansoft/internal/storage"
	"vatansoft/pkg/api"
	"vatansoft/pkg/repository"
	"vatansoft/pkg/service"
)

type Handler struct {
	Api         *api.Api
	CategoryApi *api.CategoryApi
	BillApi     *api.BillApi
	PropertyApi *api.PropertyApi
}

func NewStockHandler() *Handler {
	dbInstance := storage.GetDB()
	redisInstance, err := storage.NewRedisClient("45.12.81.218:6379", "toor")
	if err != nil {
		fmt.Println(err.Error())
	}
	//stock
	stockrepository := repository.NewProductRepository(dbInstance, redisInstance)
	stockservice := service.NewStockService(stockrepository)
	stockApi := api.NewStockApi(stockservice)
	//category
	categoryRepository := repository.NewCategoryRepository(dbInstance, redisInstance)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryApi := api.NewCategoryApi(categoryService)
	//bill
	billRepository := repository.NewBillRepository(dbInstance, redisInstance)
	billService := service.NewBillService(billRepository)
	billApi := api.NewBillApi(billService)
	//property
	propertyRepository := repository.NewPropertyRepository(dbInstance, redisInstance)
	propertyService := service.NewPropertyService(propertyRepository)
	propertyApi := api.NewPropertyApi(propertyService)
	return &Handler{
		Api:         stockApi,
		CategoryApi: categoryApi,
		BillApi:     billApi,
		PropertyApi: propertyApi,
	}
}
