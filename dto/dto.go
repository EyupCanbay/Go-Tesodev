package dto

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
}

type ServiceProduct struct {
	Name        string
	Price       float64
	Description string
}

type ProductResponse struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
