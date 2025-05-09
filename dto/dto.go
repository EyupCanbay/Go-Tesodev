package dto

type ProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
}

type ServiceProduct struct {
	Name        string
	Price       float64
	Description string
}

type SuccessResponse struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type ErrorResponse struct {
	StatusCode int    `json:"status"`
	Message    bool   `json:"message"`
	Error      string `json:"error"`
}
