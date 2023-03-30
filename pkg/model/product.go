package model

import (
	"gorm.io/gorm"
)

type Categories struct {
	Categories []string `json:"categories" gorm:"type:text;"`
}
type Product struct {
	gorm.Model
	Name        string   `json:"name"`
	UnitPrice   string   `json:"unitPrice"`
	StockAmount string   `json:"stockAmount"`
	Feature     string   `json:"feature"`
	Category    []string `json:"category" gorm:"type:text;"`
}

type ProductDTO struct {
	ID          uint     `gorm:"primaryKey,omitempty"`
	Name        string   `json:"name"`
	UnitPrice   string   `json:"unitPrice"`
	StockAmount string   `json:"stockAmount"`
	Feature     string   `json:"feature"`
	Category    []string `json:"category" gorm:"type:text;"`
}
type ProductResponse struct {
	Message    string      `json:"message"`
	ProductDTO *ProductDTO `json:"product"`
}
type DeletedProduct struct {
	gorm.Model
	Name        string `json:"name"`
	UnitPrice   string `json:"unitPrice"`
	StockAmount string `json:"stockAmount"`
	Feature     string `json:"feature"`
}

func ProductToProductDTO(product *Product) *ProductDTO {
	return &ProductDTO{

		Name:        product.Name,
		UnitPrice:   product.UnitPrice,
		StockAmount: product.StockAmount,
		Feature:     product.Feature,
		Category:    product.Category,
	}
}
func ProductDTOToProduct(product *Product) *Product {
	return &Product{
		Model:       gorm.Model{},
		Name:        product.Name,
		UnitPrice:   product.UnitPrice,
		StockAmount: product.StockAmount,
		Feature:     product.Feature,
		Category:    product.Category,
	}
}
func ProductToProductResponse(product *Product) *ProductResponse {
	productDTO := &ProductDTO{
		Name:        product.Name,
		UnitPrice:   product.UnitPrice,
		StockAmount: product.StockAmount,
		Feature:     product.Feature,
		Category:    product.Category,
	}
	return &ProductResponse{
		Message:    "Başarılı",
		ProductDTO: productDTO,
	}
}
func ProductDTOToProductResponse(productDTO *ProductDTO) *ProductResponse {
	return &ProductResponse{
		Message:    "Başarılı",
		ProductDTO: productDTO,
	}
}
