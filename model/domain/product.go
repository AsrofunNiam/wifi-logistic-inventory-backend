package domain

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"gorm.io/gorm"
)

type Products []Product
type Product struct {
	gorm.Model
	CreatedByID uint `gorm:"default:null"`
	UpdatedByID uint `gorm:"default:null"`
	DeletedByID uint `gorm:"default:null"`

	// Required Fields
	Code        string  `gorm:"type:varchar(50);uniqueIndex;not null"`
	Name        string  `gorm:"type:varchar(255);not null"`
	CategoryID  uint    `gorm:"default:null"`
	SupplierID  uint    `gorm:"default:null"`
	Type        string  `gorm:"type:text"`
	CompanyCode uint    `gorm:"not null"`
	Description string  `gorm:"type:text"`
	Images      string  `gorm:"type:text"`
	Available   bool    `gorm:"default:true"`
	Stock       int     `gorm:"default:0"`
	MinStock    int     `gorm:"default:0"`
	Unit        string  `gorm:"type:varchar(50);default:'Unit'"`
	Price       float64 `gorm:"type:decimal(20,2);default:0"`

	// Relations
	Company      Company      `gorm:"foreignKey:CompanyCode;references:ID"`
	ProductPrice ProductPrice `gorm:"foreignKey:ProductID;references:ID"`
	Category     Category     `gorm:"foreignKey:CategoryID"`
	Supplier     Supplier     `gorm:"foreignKey:SupplierID"`
}

func (product *Product) ToProductResponse() web.ProductResponse {
	return web.ProductResponse{
		// Required Fields
		ID:          product.ID,
		Code:        product.Code,
		Name:        product.Name,
		CategoryID:  product.CategoryID,
		SupplierID:  product.SupplierID,
		Type:        product.Type,
		CompanyCode: product.CompanyCode,
		Description: product.Description,
		Images:      product.Images,
		Available:   product.Available,
		Stock:       product.Stock,
		MinStock:    product.MinStock,
		Unit:        product.Unit,
		Price:       product.Price,

		// Relations
		Company:      product.Company.ToCompanyResponse(),
		ProductPrice: product.ProductPrice.ToProductPriceResponse(),
	}
}

func (users Products) ToProductResponses() []web.ProductResponse {
	productResponses := []web.ProductResponse{}
	for _, user := range users {
		productResponses = append(productResponses, user.ToProductResponse())
	}
	return productResponses
}
