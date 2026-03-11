package web

import "time"

type StockInResponse struct {
	ID           uint      `json:"id"`
	Code         string    `json:"code"`
	Date         string    `json:"date"`
	ProductID    uint      `json:"product_id"`
	ProductName  string    `json:"product_name"`
	ProductCode  string    `json:"product_code"`
	ProductUnit  string    `json:"product_unit"`
	SupplierID   uint      `json:"supplier_id"`
	SupplierName string    `json:"supplier_name"`
	Quantity     int       `json:"quantity"`
	Notes        string    `json:"notes"`
	CreatedByID  uint      `json:"created_by_id"`
	CreatedBy    string    `json:"created_by"`
	CreatedAt    time.Time `json:"created_at"`
}

type StockInCreateRequest struct {
	Code       string `json:"code" validate:"required"`
	Date       string `json:"date" validate:"required"`
	ProductID  uint   `json:"product_id" validate:"required"`
	SupplierID uint   `json:"supplier_id" validate:"required"`
	Quantity   int    `json:"quantity" validate:"required,gt=0"`
	Notes      string `json:"notes"`
}

type StockInUpdateRequest struct {
	Code       string `json:"code"`
	Date       string `json:"date"`
	ProductID  uint   `json:"product_id"`
	SupplierID uint   `json:"supplier_id"`
	Quantity   int    `json:"quantity"`
	Notes      string `json:"notes"`
}
