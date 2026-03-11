package web

import "mime/multipart"

type ProductCreateRequest struct {
	Name        string                  `json:"name" validate:"required"`
	Type        string                  `json:"type" validate:"required"`
	CompanyCode uint                    `json:"company_code" validate:"required"`
	Description string                  `json:"description"`
	ImageFile   []*multipart.FileHeader `json:"image"`
	Available   bool                    `json:"available" validate:"required"`
}
