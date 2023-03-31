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
}

func NewStockHandler() *Handler {
	dbInstance := storage.GetDB()
	redisInstance, err := storage.NewRedisClient("45.12.81.218:6379", "toor")
	if err != nil {
		fmt.Println(err.Error())
	}
	stockrepository := repository.NewProductRepository(dbInstance, redisInstance)
	stockservice := service.NewStockService(stockrepository)
	stockApi := api.NewStockApi(stockservice)
	categoryRepository := repository.NewCategoryRepository(dbInstance, redisInstance)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryApi := api.NewCategoryApi(categoryService)
	return &Handler{
		Api:         stockApi,
		CategoryApi: categoryApi,
	}
}
