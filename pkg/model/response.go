package model

// ProductResponse represents a product response object
type ProductResponse struct {
	Message    string      `json:"message"`
	ProductDTO *ProductDTO `json:"product"`
}

// ProductDTO represents a product data transfer object
type ProductDTO struct {
	ID          uint
	Name        string     `json:"name"`
	Price       float64    `json:"price"`
	Quantity    int        `json:"quantity"`
	Description string     `json:"description"`
	Categories  []Category `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	IsSold      string     `json:"sold"`
}
