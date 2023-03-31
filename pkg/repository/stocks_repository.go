package repository

import (
	"errors"
	"strconv"
	"vatansoft/internal/storage"
	"vatansoft/pkg/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB    *gorm.DB
	Redis *storage.RedisClient
}

var (
	productTable = "products"
)

func NewProductRepository(db *gorm.DB, redis *storage.RedisClient) *ProductRepository {
	return &ProductRepository{
		DB:    db,
		Redis: redis,
	}
}

func (r *ProductRepository) CreateProduct(c echo.Context, dto *model.ProductDTO) (*model.ProductResponse, error) {
	if dto.Name == "" || dto.Description == "" || dto.Price == 0 {
		return nil, errors.New("verideki bazı alanlar boş")
	}

	if err := r.DB.Table(productTable).Create(dto).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	id := strconv.Itoa(int(dto.ID))
	r.Redis.Set(id, dto)
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
	}

	// Update the product with the given ID in the database
	if err := r.DB.Table(productTable).Where("id = ?", id).Updates(temporaryProduct).Error; err != nil {
		return nil, err
	}
	r.Redis.Set(id, dto)
	// Convert the updated product to a ProductResponse object and return it
	return model.CreateProductResponseFromDTO(temporaryProduct), nil
}
func (r *ProductRepository) DeleteProduct(c echo.Context, id string) (*model.Product, error) {
	var product model.Product
	result := r.DB.Table(productTable).Where("id = ?", id).Scan(&product).Delete(&product)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		return nil, result.Error
	}
	r.Redis.Delete(id)
	return &product, nil
}

func (r *ProductRepository) FilterSearchProducts(c echo.Context, query, category, minPrice, maxPrice string) ([]*model.Product, error) {
	db := r.DB.Table(productTable).Model(&model.Product{})
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
	if err := db.Table(productTable).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetAllProducts(c echo.Context) ([]*model.Product, error) {
	var products []*model.Product
	if err := r.DB.Unscoped().Table(productTable).Find(&products).Error; err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, errors.New("sistemde ürün bulunmamaktadır")
	}
	r.Redis.Set("products", products)
	return products, nil
}

func (r *ProductRepository) GetProductById(c echo.Context, id string) (*model.ProductDTO, error) {
	product := &model.Product{}
	if err := r.DB.Table(productTable).Where("id = ?", id).First(product).Error; err != nil {
		return nil, err
	}
	newId, _ := strconv.Atoi(id)
	product.ID = uint(newId)
	r.Redis.Set(id, product)
	return model.ToDTO(product), nil
}
func (r *ProductRepository) InsertCategoryForAllProduct(c echo.Context, category model.Category) (*[]model.Product, error) {
	err := r.DB.Table(categoryTable).Find(&category)
	if err != nil {
		return nil, err.Error
	}
	var products []model.Product
	if err := r.DB.Table(productTable).Find(&products).Error; err != nil {
		return nil, err
	}
	for i := range products {
		products[i].Categories = append(products[i].Categories, category)
	}
	if updateErr := r.DB.Table(productTable).Updates(&products); err != nil {
		return nil, updateErr.Error
	}
	r.Redis.Set("products", products)
	return &products, nil
}
func (r *ProductRepository) DeleteCategoryForProductById(c echo.Context, id, categoryId string) (*model.ProductDTO, error) {
	var category model.Category
	err := r.DB.Table(categoryTable).Where("id = ?", categoryId).Scan(&category)
	if err != nil {
		return nil, err.Error
	}
	var product model.Product
	if err := r.DB.Table(productTable).Where("id = ?", id).Find(&product).Error; err != nil {
		return nil, err
	}
	for i, s := range product.Categories {
		if s.Name == category.Name && s.Description == category.Description {
			product.Categories = append(product.Categories[:i], product.Categories[i+1:]...)
			break
		}
	}

	if updateErr := r.DB.Table(productTable).Where("id = ?", id).Updates(&product); err != nil {
		return nil, updateErr.Error
	}
	r.Redis.Set(id, product)
	return model.ToDTO(&product), nil
}
func (r *ProductRepository) DeleteCategoryForProductsById(c echo.Context, id string) (*model.ProductDTO, error) {
	var product model.Product
	if err := r.DB.Table(productTable).Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	product.Categories = []model.Category{}

	if updateErr := r.DB.Save(&product).Error; updateErr != nil {
		return nil, updateErr
	}
	r.Redis.Set(id, product)
	return model.ToDTO(&product), nil
}
