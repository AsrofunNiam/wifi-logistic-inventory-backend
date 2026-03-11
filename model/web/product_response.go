package web

type ProductResponse struct {
	ID           uint                 `json:"id"`
	Code         string               `json:"code"`
	Name         string               `json:"name"`
	CategoryID   uint                 `json:"category_id"`
	SupplierID   uint                 `json:"supplier_id"`
	Type         string               `json:"type"`
	CompanyCode  uint                 `json:"company_code"`
	Description  string               `json:"description"`
	Images       string               `json:"images"`
	Available    bool                 `json:"available"`
	Stock        int                  `json:"stock"`
	MinStock     int                  `json:"min_stock"`
	Unit         string               `json:"unit"`
	Price        float64              `json:"price"`
	Company      CompanyResponse      `json:"company"`
	ProductPrice ProductPriceResponse `json:"product_price"`
}
