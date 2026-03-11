package web

import "time"

type SupplierResponse struct {
	ID        uint      `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Contact   string    `json:"contact"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SupplierCreateRequest struct {
	Code    string `json:"code" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Contact string `json:"contact"`
	Phone   string `json:"phone"`
	Email   string `json:"email" validate:"omitempty,email"`
	Address string `json:"address"`
	Status  string `json:"status"`
}

type SupplierUpdateRequest struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Contact string `json:"contact"`
	Phone   string `json:"phone"`
	Email   string `json:"email" validate:"omitempty,email"`
	Address string `json:"address"`
	Status  string `json:"status"`
}
