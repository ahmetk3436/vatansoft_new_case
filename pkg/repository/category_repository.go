package repository

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
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
	r.Redis.Set("category"+id, newCategory, time.Minute)
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
	r.Redis.Delete("category" + id)
	return &category, nil
}
func (r *CategoryRepository) GetAllCategories(c echo.Context) ([]*model.Category, error) {
	data, err := r.Redis.Get("categories")
	var redisData []*model.Category
	if err == nil {
		if len(data) > 0 {
			if err := json.Unmarshal(data, &redisData); err != nil {
				return nil, err
			}
			return redisData, nil
		}
	}
	var categories []*model.Category
	if err := r.DB.Unscoped().Table(categoryTable).Find(&categories).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	r.Redis.Set("categories", categories, time.Minute)
	return categories, nil
}

func (r *CategoryRepository) GetCategoryById(c echo.Context, id string) (*model.Category, error) {
	data, err := r.Redis.Get("category" + id)
	var redisData *model.Category
	if err == nil {
		if len(data) > 0 {
			if err := json.Unmarshal(data, &redisData); err != nil {
				return nil, err
			}
			return redisData, nil
		}
	}
	category := &model.Category{}
	if err := r.DB.Table(categoryTable).Where("category_id = ?", id).First(category).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	newId, _ := strconv.Atoi(id)
	category.CategoryID = uint(newId)
	r.Redis.Set("category"+id, category, time.Minute)
	return category, nil
}
