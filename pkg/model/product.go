package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string   `json:"name"`
	Category    []string `json:"category"`
	UnitPrice   string   `json:"unitPrice"`
	StockAmount string   `json:"stockAmount"`
	Feature     string   `json:"feature"`
}
type ProductDTO struct {
	ID          uint     `gorm:"primaryKey,omitempty"`
	Name        string   `json:"name"`
	Category    []string `json:"category"`
	UnitPrice   string   `json:"unitPrice"`
	StockAmount string   `json:"stockAmount"`
	Feature     string   `json:"feature"`
}
type ProductResponse struct {
	Message    string      `json:"message"`
	ProductDTO *ProductDTO `json:"product"`
}
type DeletedProduct struct {
	gorm.Model
	Name        string   `json:"name"`
	Category    []string `json:"category"`
	UnitPrice   string   `json:"unitPrice"`
	StockAmount string   `json:"stockAmount"`
	Feature     string   `json:"feature"`
}

func ProductToProductDTO(product *Product) *ProductDTO {
	return &ProductDTO{

		Name:        product.Name,
		Category:    product.Category,
		UnitPrice:   product.UnitPrice,
		StockAmount: product.StockAmount,
		Feature:     product.Feature,
	}
}
func ProductDTOToProduct(product *Product) *Product {
	return &Product{
		Model:       gorm.Model{},
		Name:        product.Name,
		Category:    product.Category,
		UnitPrice:   product.UnitPrice,
		StockAmount: product.StockAmount,
		Feature:     product.Feature,
	}
}
func ProductToProductResponse(product *Product) *ProductResponse {
	productDTO := &ProductDTO{
		Name:        product.Name,
		Category:    product.Category,
		UnitPrice:   product.UnitPrice,
		StockAmount: product.StockAmount,
		Feature:     product.Feature,
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
