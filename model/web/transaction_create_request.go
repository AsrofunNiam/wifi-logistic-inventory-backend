package web

type TransactionCreateRequest struct {
	UserID          uint    `json:"user_id"`
	OnTheRoad       float64 `json:"on_the_road"`
	AdminFee        float64 `json:"admin_fee"`
	ProductID       uint    `json:"product_id" validate:"required"`
	TransactionType string  `json:"transaction_type"`
}
