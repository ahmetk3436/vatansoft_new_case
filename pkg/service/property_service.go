package service

import (
	"errors"
	"vatansoft/pkg/model"
	"vatansoft/pkg/repository"

	"github.com/labstack/echo/v4"
)

type PropertyService struct {
	repository *repository.PropertyRepository
}

func NewPropertyService(repository *repository.PropertyRepository) *PropertyService {
	return &PropertyService{
		repository: repository,
	}
}

func (c *PropertyService) CreatePropertyService(e echo.Context, dto *model.ProductProperty) (_ *model.ProductProperty, err error) {
	dto, err = c.repository.CreateProperty(e, dto)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return dto, nil
}

func (c *PropertyService) UpdatePropertyService(e echo.Context, id string, newProperty *model.ProductProperty) (*model.ProductProperty, error) {
	newProperty, updateErr := c.repository.UpdateProperty(e, id, newProperty)
	if updateErr != nil {
		return nil, errors.New(updateErr.Error())
	}
	return newProperty, nil
}

func (c *PropertyService) DeletePropertyService(e echo.Context, id string) (property *model.ProductProperty, err error) {
	property, err = c.repository.DeleteProperty(e, id)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return property, nil
}
func (c *PropertyService) GetAllPropertysService(e echo.Context) (propertys []*model.ProductProperty, err error) {
	propertys, err = c.repository.GetAllPropertys(e)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return propertys, nil
}

func (c *PropertyService) GetPropertyByIdService(e echo.Context, id string) (property *model.ProductProperty, err error) {
	property, err = c.repository.GetPropertyById(e, id)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return property, nil
}
