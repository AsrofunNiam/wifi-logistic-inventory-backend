package web

type UserResponse struct {
	// Fields
	ID           uint    `json:"id"`
	FullName     string  `json:"full_name"`
	LegalName    string  `json:"legal_name"`
	PlaceOfBirth string  `json:"place_of_birth"`
	DateOfBirth  string  `json:"date_of_birth"`
	Salary       float64 `json:"salary"`
	Role         string  `json:"role"`
	NumberPhone  string  `json:"number_phone"`
	Email        string  `json:"email"`
}

type UserLoginRequest struct {
	AccessTokenLogin string `json:"access_token_login" validate:"required"`
}
