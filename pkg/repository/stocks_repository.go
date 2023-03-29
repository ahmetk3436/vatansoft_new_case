package repository

import (
	"vatansoft/pkg/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewStockRepository(db *gorm.DB) *Repository {
	return &Repository{}
}

func (r *Repository) CreateStockProduct(e echo.Context) (product *model.ProductResponse, err error) {
	requestBody := &model.ProductDTO{}
	if err := e.Bind(&requestBody); err != nil {
		return nil, err
	}
	r.DB.Create(requestBody)
	return model.ProductDTOToProductResponse(requestBody), nil
}
func (r *Repository) UpdateStockProduct(e echo.Context) (product *model.ProductResponse, err error) {
	id := e.Param("id")
	newProduct := &model.ProductDTO{}
	if err := e.Bind(&newProduct); err != nil {
		return nil, err
	}
	temporaryProduct := &model.ProductDTO{}
	temporaryProduct.Category = newProduct.Category
	temporaryProduct.Feature = newProduct.Feature
	temporaryProduct.Name = newProduct.Name
	temporaryProduct.UnitPrice = newProduct.UnitPrice
	temporaryProduct.StockAmount = newProduct.StockAmount
	r.DB.Table("Products").Where("id = ?", id).First(&newProduct)
	newProduct.ID = temporaryProduct.ID
	newProduct.Category = temporaryProduct.Category
	newProduct.Feature = temporaryProduct.Feature
	newProduct.Name = temporaryProduct.Name
	newProduct.UnitPrice = temporaryProduct.UnitPrice
	newProduct.StockAmount = temporaryProduct.StockAmount
	r.DB.Table("Products").Where("id = ?", id).Save(newProduct)
	return model.ProductDTOToProductResponse(newProduct), nil
}
func (r *Repository) DeleteStockProduct(e echo.Context) (product *model.Product, err error) {
	id := e.Param("id")
	result := r.DB.Delete(&product, id)
	if result.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		return nil, result.Error

	}
	return product, nil
}
func (r *Repository) FilterSearchStockProduct(e echo.Context) ([]*model.Product, error) {
	query := e.QueryParam("query")
	category := e.QueryParam("category")
	minPrice := e.QueryParam("min_price")
	maxPrice := e.QueryParam("max_price")
	var products []*model.Product
	db := r.DB.Model(&model.Product{})
	if query != "" {
		db = db.Where("name LIKE ?", "%"+query+"%")
	}
	if category != "" {
		db = db.Where("category = ?", category)
	}
	if minPrice != "" {
		db = db.Where("unit_price >= ?", minPrice)
	}
	if maxPrice != "" {
		db = db.Where("unit_price <= ?", maxPrice)
	}
	result := db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
func (r *Repository) GetAllStockProducts(e echo.Context) (products []*model.Product, err error) {
	result := r.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
func (r *Repository) GetStockProductById(e echo.Context) (productDTO *model.ProductDTO, err error) {
	id := e.Param("id")
	var product model.Product
	result := r.DB.Table("Products").Where("id = ?", id).Find(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return model.ProductToProductDTO(&product), nil
}
