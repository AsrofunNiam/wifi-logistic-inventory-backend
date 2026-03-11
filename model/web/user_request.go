package web

type UserCreateRequest struct {
	FullName  string `json:"full_name" validate:"required"`
	LegalName string `json:"legal_name"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Role      string `json:"role" validate:"required"`
	Status    string `json:"status"`
}

type UserUpdateRequest struct {
	FullName  string `json:"full_name"`
	LegalName string `json:"legal_name"`
	Email     string `json:"email" validate:"omitempty,email"`
	Role      string `json:"role"`
	Status    string `json:"status"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}
