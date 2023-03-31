package service

import (
	"errors"
	"vatansoft/pkg/model"
	"vatansoft/pkg/repository"

	"github.com/labstack/echo/v4"
)

type CategoryService struct {
	repository *repository.CategoryRepository
}

func NewCategoryService(repository *repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		repository: repository,
	}
}

func (c *CategoryService) CreateCategoryService(e echo.Context, dto *model.Category) (_ *model.Category, err error) {
	dto, err = c.repository.CreateCategory(e, dto)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return dto, nil
}

func (c *CategoryService) UpdateCategoryService(e echo.Context, id string, newCategory *model.Category) (*model.Category, error) {
	newCategory, updateErr := c.repository.UpdateCategory(e, id, newCategory)
	if updateErr != nil {
		return nil, errors.New(updateErr.Error())
	}
	return newCategory, nil
}

func (c *CategoryService) DeleteCategoryService(e echo.Context, id string) (category *model.Category, err error) {
	category, err = c.repository.DeleteCategory(e, id)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return category, nil
}
func (c *CategoryService) GetAllCategoriesService(e echo.Context) (categories []*model.Category, err error) {
	categories, err = c.repository.GetAllCategories(e)
	if err != nil && len(categories) == 0 {
		return nil, errors.New("sistemde ürün bulunmamaktadır")
	}
	return categories, nil
}

func (c *CategoryService) GetCategoryByIdService(e echo.Context, id string) (category *model.Category, err error) {
	category, err = c.repository.GetCategoryById(e, id)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return category, nil
}
