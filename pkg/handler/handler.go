package handler

import (
	"vatansoft/internal/db"
	"vatansoft/pkg/api"
	"vatansoft/pkg/repository"
	"vatansoft/pkg/service"
)

type StockHandler struct {
	Api *api.Api
}

func NewStockHandler() *StockHandler {
	dbInstance := db.GetDB()
	stockrepository := repository.NewProductRepository(dbInstance)
	stockservice := service.NewStockService(stockrepository)
	stockApi := api.NewStockApi(stockservice)
	return &StockHandler{
		Api: stockApi,
	}
}
