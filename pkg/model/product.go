package model

import (
	"gorm.io/gorm"
)

// Product represents a stock item
type Product struct {
	gorm.Model
	Name        string            `gorm:"not null" json:"name"`
	Description string            `json:"description"`
	Price       float64           `gorm:"not null" json:"price"`
	Quantity    int               `gorm:"not null" json:"quantity"`
	CategoryID  uint              `gorm:"index"`
	Category    Category          `json:"category" gorm:"foreignkey:CategoryID"`
	Properties  []ProductProperty `json:"properties"`
}

// Category represents a stock category
type Category struct {
	gorm.Model
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Products    []Product `json:"products"`
}
type ProductCategory struct {
	gorm.Model
	ProductID  uint `json:"productID"`
	CategoryID uint `json:"categoryID"`
}

// ProductProperty represents a product property
type ProductProperty struct {
	gorm.Model
	Name      string  `gorm:"not null" json:"name"`
	Value     string  `json:"value"`
	ProductID uint    `gorm:"index"`
	Product   Product `json:"product" gorm:"foreignkey:ProductID"`
}

// Invoice represents a product invoice
type Invoice struct {
	gorm.Model
	InvoiceNo  string  `gorm:"not null" json:"invoice_no"`
	ProductID  uint    `gorm:"index"`
	Product    Product `json:"product" gorm:"foreignkey:ProductID"`
	Quantity   int     `gorm:"not null" json:"quantity"`
	TotalPrice float64 `gorm:"not null" json:"total_price"`
}

// ProductDTO represents a product data transfer object
type ProductDTO struct {
	ID          uint
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Quantity    int      `json:"quantity"`
	Description string   `json:"description"`
	CategoryID  uint     `gorm:"index"`
	Category    Category `json:"category" gorm:"foreignkey:CategoryID"`
}

// ProductResponse represents a product response object
type ProductResponse struct {
	Message    string      `json:"message"`
	ProductDTO *ProductDTO `json:"product"`
}

// ToDTO converts a product to a product DTO
func ToDTO(p *Product) *ProductDTO {
	return &ProductDTO{
		ID:          p.ID,
		Name:        p.Name,
		Price:       p.Price,
		Quantity:    p.Quantity,
		Description: p.Description,
		CategoryID:  p.CategoryID,
	}
}

// Creates a product response object from a product
func CreateProductResponse(p *Product) *ProductResponse {
	return &ProductResponse{
		Message:    "Success",
		ProductDTO: ToDTO(p),
	}
}

// Creates a product response object from a product DTO
func CreateProductResponseFromDTO(p *ProductDTO) *ProductResponse {
	return &ProductResponse{
		Message:    "Success",
		ProductDTO: p,
	}
}

// ToProduct converts a product DTO to a product
func ToProduct(p *ProductDTO) *Product {
	return &Product{
		Model:       gorm.Model{},
		Name:        p.Name,
		Price:       p.Price,
		Quantity:    p.Quantity,
		Description: p.Description,
		CategoryID:  p.CategoryID,
	}
}
