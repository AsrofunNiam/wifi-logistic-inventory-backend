package web

type ProductCreateRequest struct {
	Code        string  `json:"code" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	CategoryID  uint    `json:"category_id"`
	SupplierID  string  `json:"supplier_id"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	MinStock    int     `json:"min_stock"`
	Unit        string  `json:"unit"`
	Price       float64 `json:"price"`
}
