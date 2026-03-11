package web

type UserResponse struct {
	ID          uint   `json:"id"`
	FullName    string `json:"full_name"`
	LegalName   string `json:"legal_name"`
	Role        string `json:"role"`
	NumberPhone string `json:"number_phone"`
	Email       string `json:"email"`
	Status      string `json:"status"`
}

type UserLoginRequest struct {
	AccessTokenLogin string `json:"access_token_login" validate:"required"`
}
