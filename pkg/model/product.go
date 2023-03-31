package model

import (
	"bytes"
	"encoding/gob"

	"gorm.io/gorm"
)

// Product represents a stock item
type Product struct {
	gorm.Model
	Name        string     `gorm:"not null" json:"name"`
	Description string     `json:"description"`
	Price       float64    `gorm:"not null" json:"price"`
	Quantity    int        `gorm:"not null" json:"quantity"`
	Categories  []Category `gorm:"many2many:product_categories"`
}

// ProductProperty represents a product property
type ProductProperty struct {
	gorm.Model
	Name      string `gorm:"not null" json:"name"`
	Value     string `json:"value"`
	ProductID uint   `json:"productID"`
}

// Invoice represents a product invoice
type Invoice struct {
	gorm.Model
	InvoiceNo  string  `gorm:"not null" json:"invoice_no"`
	ProductID  uint    `json:"productID"`
	Quantity   int     `gorm:"not null" json:"quantity"`
	TotalPrice float64 `gorm:"not null" json:"total_price"`
}

// ToDTO converts a product to a product DTO
func ToDTO(p *Product) *ProductDTO {
	return &ProductDTO{
		ID:          p.ID,
		Name:        p.Name,
		Price:       p.Price,
		Quantity:    p.Quantity,
		Description: p.Description,
		Categories:  p.Categories,
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
	}
}
func (p *Product) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(p); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
