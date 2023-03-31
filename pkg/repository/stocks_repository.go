package repository

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"vatansoft/internal/storage"
	"vatansoft/pkg/model"
)

const (
	productTable           = "products"
	productCategoriesTable = "product_categories"
)

type ProductRepository struct {
	DB    *gorm.DB
	Redis *storage.RedisClient
}

func NewProductRepository(db *gorm.DB, redis *storage.RedisClient) *ProductRepository {
	return &ProductRepository{
		DB:    db,
		Redis: redis,
	}
}

func (r *ProductRepository) CreateProduct(c echo.Context, dto *model.ProductDTO) (*model.ProductResponse, error) {
	if dto.Name == "" || dto.Description == "" || dto.Price == 0 {
		return nil, errors.New("some fields in the data are empty")
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
		return nil, errors.New(err.Error())
	}
	r.Redis.Set(id, dto)
	// Convert the updated product to a ProductResponse object and return it
	return model.CreateProductResponseFromDTO(temporaryProduct), nil
}

func (r *ProductRepository) DeleteProduct(c echo.Context, id string) (*model.Product, error) {
	var product model.Product
	result := r.DB.Table(productTable).Where("id = ?", id).Scan(&product).Delete(product.ID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New(result.Error.Error())
		}
		return nil, errors.New(result.Error.Error())
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
		return nil, errors.New(err.Error())
	}
	return products, nil
}

func (r *ProductRepository) GetAllProducts(c echo.Context) ([]*model.Product, error) {
	var products []*model.Product
	if err := r.DB.Unscoped().Table(productTable).Find(&products).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	if len(products) == 0 {
		return nil, errors.New("sistemde ürün bulunmamaktadır")
	}
	r.Redis.Set("products", products)
	return products, nil
}

func (r *ProductRepository) GetProductByID(c echo.Context, id string) (*model.ProductDTO, error) {
	product := &model.Product{}
	if err := r.DB.Table(productTable).Where("id = ?", id).First(product).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	newID, _ := strconv.Atoi(id)
	product.ID = uint(newID)
	marshalProduct, marshalErr := product.MarshalBinary()
	if marshalErr != nil {
		return nil, errors.New(marshalErr.Error())
	}
	if err := r.Redis.Set(id, marshalProduct); err != nil {
		return nil, errors.New(err.Error())
	}

	return model.ToDTO(product), nil
}

func (r *ProductRepository) InsertCategoryForAllProducts(c echo.Context, category model.Category) ([]*model.ProductCategory, error) {
	// check if category exists
	if err := r.DB.Table(categoryTable).Find(&category).Error; err != nil {
		if err != nil {
			return nil, errors.New("category not found")
		}
		return nil, errors.New("database error")
	}

	// get all products and insert the new category for each product
	products, err := r.getAllProducts()
	if err != nil {
		return nil, err
	}
	var productCategories []*model.ProductCategory
	for _, p := range products {
		pc, err := r.addCategoryToProduct(p.ID, category.CategoryID)
		if err != nil {
			return nil, err
		}
		productCategories = append(productCategories, pc)
	}

	return productCategories, nil
}

func (r *ProductRepository) getAllProducts() ([]model.Product, error) {
	var products []model.Product
	if err := r.DB.Table(productTable).Find(&products).Error; err != nil {
		return nil, errors.New("database error")
	}
	return products, nil
}

func (r *ProductRepository) addCategoryToProduct(productID uint, categoryID uint) (*model.ProductCategory, error) {
	pc := &model.ProductCategory{
		ProductID:  productID,
		CategoryID: categoryID,
	}
	if err := r.DB.Table(productCategoriesTable).Save(pc).Error; err != nil {
		return nil, errors.New("database error")
	}

	// store the product category in redis
	if err := r.Redis.Set(strconv.Itoa(int(pc.ID)), pc).Error(); err != "" {
		return nil, errors.New("redis error")
	}

	return pc, nil
}

func (r *ProductRepository) DeleteCategoryForProductByID(c echo.Context, id, categoryID string) (*model.ProductCategory, error) {
	var productCategory model.ProductCategory
	if err := r.DB.Table(productCategoriesTable).Where("product_id = ? AND category_id = ?", id, categoryID).Find(&productCategory).Delete(productCategory.ID).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	return &productCategory, nil
}

func (r *ProductRepository) DeleteCategoriesForProductByID(c echo.Context, id string) ([]*model.ProductCategory, error) {
	var productCategories []*model.ProductCategory
	if err := r.DB.Table(productCategoriesTable).Where("product_id = ?", id).Find(&productCategories).Delete(&productCategories).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	return productCategories, nil
}
