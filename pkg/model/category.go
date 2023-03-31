package model

import "gorm.io/gorm"

type Category struct {
	CategoryID  uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null" json:"name"`
	Description string `json:"description"`
}
type ProductCategory struct {
	gorm.Model
	ProductID  uint `json:"ProductID"`
	CategoryID uint `json:"CategoryID"`
}
