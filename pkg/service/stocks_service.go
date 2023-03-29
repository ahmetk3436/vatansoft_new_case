package service

import (
	"vatansoft/pkg/model"
	"vatansoft/pkg/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type StockService struct {
	repository *repository.Repository
	DB         *gorm.DB
}

func NewStockService(repository *repository.Repository) *StockService {
	return &StockService{
		repository: repository,
	}
}

func (s *StockService) CreateStockProductService(e echo.Context) (product *model.ProductResponse, err error) {
	product, err = s.repository.CreateStockProduct(e)
	if err != nil {
		return nil, echo.ErrBadGateway
	}
	return product, nil
}
