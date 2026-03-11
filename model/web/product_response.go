package web

type ProductResponse struct {
	ID           uint                 `json:"id"`
	Name         string               `json:"name"`
	Type         string               `json:"type"`
	CompanyCode  uint                 `json:"company_code"`
	Description  string               `json:"description"`
	Images       string               `json:"images"`
	Available    bool                 `json:"available"`
	Company      CompanyResponse      `json:"company"`
	ProductPrice ProductPriceResponse `json:"product_price"`
}
