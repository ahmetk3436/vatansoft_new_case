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

func (s *StockService) UpdateStockProductService(e echo.Context) (product *model.ProductResponse, err error) {
	product, err = s.repository.UpdateStockProduct(e)
	if err != nil {
		return nil, echo.ErrBadGateway
	}
	return product, nil
}

func (s *StockService) DeleteStockProductService(e echo.Context) (product *model.ProductResponse, err error) {
	product, err = s.repository.CreateStockProduct(e)
	if err != nil {
		return nil, echo.ErrBadGateway
	}
	return product, nil
}

func (s *StockService) FilterSearchStockProductService(e echo.Context) (product []*model.Product, err error) {
	product, err = s.repository.FilterSearchStockProduct(e)
	if err != nil {
		return nil, echo.ErrBadGateway
	}
	return product, nil
}

func (s *StockService) GetAllStockProductsService(e echo.Context) (products []*model.Product, err error) {
	products, err = s.repository.GetAllStockProducts(e)
	if err != nil {
		return nil, echo.ErrBadGateway
	}
	return products, nil
}

func (s *StockService) GetStockProductByIdService(e echo.Context) (product *model.ProductDTO, err error) {
	product, err = s.repository.GetStockProductById(e)
	if err != nil {
		return nil, echo.ErrBadGateway
	}
	return product, nil
}
