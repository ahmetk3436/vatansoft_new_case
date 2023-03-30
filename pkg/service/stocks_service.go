package service

import (
	"vatansoft/pkg/model"
	"vatansoft/pkg/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type StockService struct {
	repository *repository.ProductRepository
	DB         *gorm.DB
}

func NewStockService(repository *repository.ProductRepository) *StockService {
	return &StockService{
		repository: repository,
	}
}

func (s *StockService) CreateStockProductService(e echo.Context, dto *model.ProductDTO) (product *model.ProductResponse, err error) {
	product, err = s.repository.CreateProduct(e, dto)
	if err != nil {
		return nil, echo.ErrBadGateway
	}
	return product, nil
}

func (s *StockService) UpdateStockProductService(e echo.Context, id string, dto *model.ProductDTO) (product *model.ProductResponse, err error) {
	product, err = s.repository.UpdateProduct(e, id, dto)
	if err != nil {
		return nil, echo.ErrBadGateway
	}
	return product, nil
}

func (s *StockService) DeleteStockProductService(e echo.Context, id string) (product *model.Product, err error) {
	product, err = s.repository.DeleteProduct(e, id)
	if err != nil {
		return nil, echo.ErrBadGateway
	}
	return product, nil
}

func (s *StockService) FilterSearchStockProductService(e echo.Context, query, category, minPrice, maxPrice string) (product []*model.Product, err error) {
	product, err = s.repository.FilterSearchProducts(e, query, category, minPrice, maxPrice)
	if err != nil {
		return nil, echo.ErrBadGateway
	}
	return product, nil
}

func (s *StockService) GetAllStockProductsService(e echo.Context) (products []*model.ProductDTO, err error) {
	products, err = s.repository.GetAllProducts(e)
	if err != nil {
		return nil, echo.ErrBadGateway
	}
	return products, nil
}

func (s *StockService) GetStockProductByIdService(e echo.Context, id string) (product *model.ProductDTO, err error) {
	product, err = s.repository.GetProductById(e, id)
	if err != nil {
		return nil, echo.ErrBadGateway
	}
	return product, nil
}
