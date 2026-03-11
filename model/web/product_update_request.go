package web

import "mime/multipart"

type ProductUpdateRequest struct {
	Name        string                  `json:"name"`
	Type        string                  `json:"type"`
	CompanyCode uint                    `json:"company_code"`
	Description string                  `json:"description"`
	ImageName   string                  `json:"images_name"`
	ImageFile   []*multipart.FileHeader `json:"image"`
	Available   bool                    `json:"available"`
}
