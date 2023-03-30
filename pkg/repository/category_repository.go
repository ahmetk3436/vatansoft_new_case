package repository

import (
	"strconv"
	"vatansoft/pkg/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

var (
	categoryTable = "categories"
)

func NewCategoryRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}
func (r *ProductRepository) CreateCategory(c echo.Context, category *model.Category) (*model.Category, error) {
	if category.Description == "" || category.Name == "" {
		return nil, echo.ErrBadRequest
	}

	if err := r.DB.Table(categoryTable).Create(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *ProductRepository) UpdateCategory(c echo.Context, id string, newCategory *model.Category) (*model.Category, error) {
	// Create a temporary ProductDTO object to store the updated values
	temporaryProduct := &model.Category{
		Model:       newCategory.Model,
		Name:        newCategory.Name,
		Description: newCategory.Description,
	}

	// Update the product with the given ID in the database
	if err := r.DB.Table(categoryTable).Where("id = ?", id).Updates(temporaryProduct).Error; err != nil {
		return nil, err
	}

	// Convert the updated product to a ProductResponse object and return it
	return newCategory, nil
}
func (r *ProductRepository) DeleteCategory(c echo.Context, id string) (*model.Category, error) {
	var category model.Category
	result := r.DB.Table(categoryTable).Where("id = ?", id).Scan(&category).Delete(&category)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &category, nil
}
func (r *ProductRepository) GetAllCategories(c echo.Context) ([]*model.Category, error) {
	var categories []*model.Category
	if err := r.DB.Unscoped().Table(categoryTable).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *ProductRepository) GetCategoryById(c echo.Context, id string) (*model.Category, error) {
	category := &model.Category{}
	if err := r.DB.Table(categoryTable).Where("id = ?", id).First(category).Error; err != nil {
		return nil, err
	}
	newId, _ := strconv.Atoi(id)
	category.ID = uint(newId)
	return category, nil
}
