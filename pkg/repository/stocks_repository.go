package repository

import (
	"strconv"
	"vatansoft/pkg/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

var (
	table = "products"
)

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (r *ProductRepository) CreateProduct(c echo.Context, dto *model.ProductDTO) (*model.ProductResponse, error) {
	if dto.Name == "" || dto.Description == "" || dto.Price != 0 || dto.CategoryID != 0 {
		return nil, echo.ErrBadRequest
	}

	if err := r.DB.Table(table).Create(dto).Error; err != nil {
		return nil, err
	}

	return model.CreateProductResponseFromDTO(dto), nil
}

func (r *ProductRepository) UpdateProduct(c echo.Context, id string, dto *model.ProductDTO) (*model.ProductResponse, error) {
	// Create a temporary ProductDTO object to store the updated values
	temporaryProduct := &model.ProductDTO{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Quantity:    dto.Quantity,
		CategoryID:  dto.CategoryID,
	}

	// Update the product with the given ID in the database
	if err := r.DB.Table(table).Where("id = ?", id).Updates(temporaryProduct).Error; err != nil {
		return nil, err
	}

	// Convert the updated product to a ProductResponse object and return it
	return model.CreateProductResponseFromDTO(temporaryProduct), nil
}
func (r *ProductRepository) DeleteProduct(c echo.Context, id string) (*model.Product, error) {
	var product model.Product
	result := r.DB.Table(table).Where("id = ?", id).Scan(&product).Delete(&product)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &product, nil
}

func (r *ProductRepository) FilterSearchProducts(c echo.Context, query, category, minPrice, maxPrice string) ([]*model.Product, error) {
	db := r.DB.Table(table).Model(&model.Product{})
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
	var products []*model.Product
	if err := db.Table(table).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetAllProducts(c echo.Context) ([]*model.ProductDTO, error) {
	var products []*model.ProductDTO
	if err := r.DB.Unscoped().Table(table).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetProductById(c echo.Context, id string) (*model.ProductDTO, error) {
	product := &model.Product{}
	if err := r.DB.Table(table).Where("id = ?", id).First(product).Error; err != nil {
		return nil, err
	}
	newId, _ := strconv.Atoi(id)
	product.ID = uint(newId)
	return model.ToDTO(product), nil
}
