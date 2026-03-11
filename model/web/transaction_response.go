package web

type TransactionResponse struct {
	ID              uint            `json:"id"`
	UserID          uint            `json:"user_id"`
	NumberContract  string          `json:"number_contract"`
	OnTheRoad       float64         `json:"on_the_road"`
	AdminFee        float64         `json:"admin_fee"`
	TotalValue      float64         `json:"total_value"`
	InterestAmount  float64         `json:"interest_amount"`
	Period          int             `json:"period"`
	TransactionType string          `json:"transaction_type"`
	User            UserResponse    `json:"user"`
	Product         ProductResponse `json:"product"`
}
