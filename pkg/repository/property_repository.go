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

type PropertyRepository struct {
	DB    *gorm.DB
	Redis *storage.RedisClient
}

var (
	PropertyTable = "product_properties"
)

func NewPropertyRepository(db *gorm.DB, redis *storage.RedisClient) *PropertyRepository {
	return &PropertyRepository{
		DB:    db,
		Redis: redis,
	}
}
func (r *PropertyRepository) CreateProperty(c echo.Context, Property *model.ProductProperty) (*model.ProductProperty, error) {
	if Property.Value == "" || Property.ProductID == 0 {
		return nil, errors.New("verilerde eksiklik mevcut")
	}
	if err := r.DB.Table(PropertyTable).Create(&Property).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	return Property, nil
}

func (r *PropertyRepository) UpdateProperty(c echo.Context, id string, newProperty *model.ProductProperty) (*model.ProductProperty, error) {

	temporaryProduct := &model.ProductProperty{
		Model:     newProperty.Model,
		Name:      newProperty.Name,
		ProductID: newProperty.ProductID,
		Value:     newProperty.Value,
	}

	if err := r.DB.Table(PropertyTable).Where("id = ?", id).Updates(temporaryProduct).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	r.Redis.Set("property"+id, newProperty, time.Minute)
	// Convert the updated product to a ProductResponse object and return it
	return newProperty, nil
}
func (r *PropertyRepository) DeleteProperty(c echo.Context, id string) (*model.ProductProperty, error) {
	var Property model.ProductProperty
	result := r.DB.Table(PropertyTable).Where("id = ?", id).Scan(&Property).Delete(id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New(result.Error.Error())
		}
		return nil, errors.New(result.Error.Error())
	}
	r.Redis.Delete("property" + id)
	return &Property, nil
}
func (r *PropertyRepository) GetAllPropertys(c echo.Context) ([]*model.ProductProperty, error) {
	data, err := r.Redis.Get("propertys")
	var redisData []*model.ProductProperty
	if err == nil {
		if len(data) > 0 {
			if err := json.Unmarshal(data, &redisData); err != nil {
				return nil, err
			}
			return redisData, nil
		}
	}
	var Propertys []*model.ProductProperty
	if err := r.DB.Unscoped().Table(PropertyTable).Find(&Propertys).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	if len(Propertys) == 0 {
		return nil, errors.New("sistemde ürün özelliği bulunmamaktadır")
	}
	r.Redis.Set("propertys", Propertys, time.Minute)
	return Propertys, nil
}

func (r *PropertyRepository) GetPropertyById(c echo.Context, id string) (*model.ProductProperty, error) {
	data, err := r.Redis.Get("property" + id)
	var redisData *model.ProductProperty
	if err == nil {
		if len(data) > 0 {
			if err := json.Unmarshal(data, &redisData); err != nil {
				return nil, err
			}
			return redisData, nil
		}
	}
	Property := &model.ProductProperty{}
	if err := r.DB.Table(PropertyTable).Where("id = ?", id).First(Property).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	newId, _ := strconv.Atoi(id)
	Property.ID = uint(newId)
	r.Redis.Set("property"+id, Property, time.Minute)
	return Property, nil
}
