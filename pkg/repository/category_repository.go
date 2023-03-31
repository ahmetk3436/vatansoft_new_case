package repository

import (
	"errors"
	"strconv"
	"vatansoft/internal/storage"
	"vatansoft/pkg/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB    *gorm.DB
	Redis *storage.RedisClient
}

var (
	categoryTable = "categories"
)

func NewCategoryRepository(db *gorm.DB, redis *storage.RedisClient) *CategoryRepository {
	return &CategoryRepository{
		DB:    db,
		Redis: redis,
	}
}
func (r *CategoryRepository) CreateCategory(c echo.Context, category *model.Category) (*model.Category, error) {
	if category.Description == "" || category.Name == "" {
		return nil, errors.New("verilerde eksiklik mevcut")
	}
	if err := r.DB.Table(categoryTable).Create(&category).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	return category, nil
}

func (r *CategoryRepository) UpdateCategory(c echo.Context, id string, newCategory *model.Category) (*model.Category, error) {
	temporaryProduct := &model.Category{
		Name:        newCategory.Name,
		Description: newCategory.Description,
	}

	if err := r.DB.Table(categoryTable).Where("category_id = ?", id).Updates(temporaryProduct).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	// Convert the updated product to a ProductResponse object and return it
	return newCategory, nil
}
func (r *CategoryRepository) DeleteCategory(c echo.Context, id string) (*model.Category, error) {
	var category model.Category
	result := r.DB.Table(categoryTable).Where("category_id = ?", id).Scan(&category).Delete(&category)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New(result.Error.Error())
		}
		return nil, errors.New(result.Error.Error())
	}
	return &category, nil
}
func (r *CategoryRepository) GetAllCategories(c echo.Context) ([]*model.Category, error) {
	var categories []*model.Category
	if err := r.DB.Unscoped().Table(categoryTable).Find(&categories).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	return categories, nil
}

func (r *CategoryRepository) GetCategoryById(c echo.Context, id string) (*model.Category, error) {
	category := &model.Category{}
	if err := r.DB.Table(categoryTable).Where("category_id = ?", id).First(category).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	newId, _ := strconv.Atoi(id)
	category.CategoryID = uint(newId)
	return category, nil
}
