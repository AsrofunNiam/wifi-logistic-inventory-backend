package web

import "time"

type StockOutResponse struct {
	ID          uint      `json:"id"`
	Code        string    `json:"code"`
	Date        string    `json:"date"`
	ProductID   uint      `json:"product_id"`
	ProductName string    `json:"product_name"`
	ProductCode string    `json:"product_code"`
	ProductUnit string    `json:"product_unit"`
	Destination string    `json:"destination"`
	Quantity    int       `json:"quantity"`
	Notes       string    `json:"notes"`
	CreatedByID uint      `json:"created_by_id"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type StockOutCreateRequest struct {
	Code        string `json:"code"`
	Date        string `json:"date" validate:"required"`
	ProductID   uint   `json:"product_id" validate:"required"`
	Destination string `json:"destination" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required,gt=0"`
	Notes       string `json:"notes"`
}

type StockOutUpdateRequest struct {
	Code        string `json:"code"`
	Date        string `json:"date"`
	ProductID   uint   `json:"product_id"`
	Destination string `json:"destination"`
	Quantity    int    `json:"quantity"`
	Notes       string `json:"notes"`
}
