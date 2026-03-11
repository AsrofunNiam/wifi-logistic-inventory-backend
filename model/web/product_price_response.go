package web

type ProductPriceResponse struct {
	ID        uint             `json:"id"`
	ProductID uint             `json:"product_id"`
	Price     float64          `json:"price"`
	StartDate string           `json:"start_date"`
	EndDate   string           `json:"end_date"`
	Currency  CurrencyResponse `json:"currency"`
}
