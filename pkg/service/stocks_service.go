package service

import (
	"errors"
	"vatansoft/pkg/model"
	"vatansoft/pkg/repository"

	"github.com/labstack/echo/v4"
)

type StockService struct {
	repository *repository.ProductRepository
}

func NewStockService(repository *repository.ProductRepository) *StockService {
	return &StockService{
		repository: repository,
	}
}

func (s *StockService) CreateStockProductService(e echo.Context, dto *model.ProductDTO) (product *model.ProductResponse, err error) {
	product, err = s.repository.CreateProduct(e, dto)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return product, nil
}

func (s *StockService) UpdateStockProductService(e echo.Context, id string, dto *model.ProductDTO) (product *model.ProductResponse, err error) {
	product, err = s.repository.UpdateProduct(e, id, dto)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return product, nil
}

func (s *StockService) DeleteStockProductService(e echo.Context, id string) (product *model.Product, err error) {
	product, err = s.repository.DeleteProduct(e, id)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return product, nil
}

func (s *StockService) FilterSearchStockProductService(e echo.Context, query, category, minPrice, maxPrice, isSold, isDeleted string) (product []*model.Product, err error) {
	product, err = s.repository.FilterSearchProducts(e, query, category, minPrice, maxPrice, isSold, isDeleted)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return product, nil
}

func (s *StockService) GetAllStockProductsService(e echo.Context) (products []*model.Product, err error) {
	products, err = s.repository.GetAllProducts(e)
	if err != nil && len(products) == 0 {
		return nil, errors.New("sistemde ürün bulunmamaktadır")
	}
	return products, nil
}

func (s *StockService) GetStockProductByIdService(e echo.Context, id string) (product *model.ProductDTO, err error) {
	product, err = s.repository.GetProductByID(e, id)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return product, nil
}
func (s *StockService) InsertCategoryForAllProductService(e echo.Context, category *model.Category) (products []*model.ProductCategory, err error) {
	products, err = s.repository.InsertCategoryForAllProducts(e, category)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return products, nil
}
func (s *StockService) DeleteCategoryForProductByIdService(e echo.Context, id, categoryId string) (product *model.ProductCategory, err error) {
	product, err = s.repository.DeleteCategoryForProductByID(e, id, categoryId)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return product, nil
}
func (s *StockService) DeleteCategoryForProductsByIdService(e echo.Context, id string) (product []*model.ProductCategory, err error) {
	product, err = s.repository.DeleteCategoriesForProductByID(e, id)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return product, nil
}
