package web

type ProductUpdateRequest struct {
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	CategoryID  uint    `json:"category_id"`
	SupplierID  uint    `json:"supplier_id"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	MinStock    int     `json:"min_stock"`
	Unit        string  `json:"unit"`
	Price       float64 `json:"price"`
}
