package web

type CategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type CategoryUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
