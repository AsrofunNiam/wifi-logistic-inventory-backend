package web

type ProductResponse struct {
	ID           uint    `json:"id"`
	Code         string  `json:"code"`
	Name         string  `json:"name"`
	CategoryID   uint    `json:"category_id"`
	CategoryName string  `json:"category_name"`
	SupplierID   uint    `json:"supplier_id"`
	SupplierName string  `json:"supplier_name"`
	Description  string  `json:"description"`
	Stock        int     `json:"stock"`
	MinStock     int     `json:"min_stock"`
	Unit         string  `json:"unit"`
	Price        float64 `json:"price"`
}
