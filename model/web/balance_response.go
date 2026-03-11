package web

type BalanceResponse struct {
	ID        uint         `json:"id"`
	UserID    uint         `json:"user_id"`
	User      UserResponse `json:"user"`
	Value     float64      `json:"value"`
	Available bool         `json:"available"`
}
