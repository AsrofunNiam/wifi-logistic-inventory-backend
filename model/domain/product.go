package domain

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"gorm.io/gorm"
)

type Products []Product
type Product struct {
	gorm.Model
	CreatedByID uint  `gorm:"default:null"`
	UpdatedByID uint  `gorm:"default:null"`
	DeletedByID *uint `gorm:"default:null"`

	// Required Fields
	Code        string  `gorm:"type:varchar(50);uniqueIndex;not null"`
	Name        string  `gorm:"type:varchar(255);not null"`
	CategoryID  uint    `gorm:"default:null"`
	SupplierID  uint    `gorm:"default:null"`
	Description string  `gorm:"type:text"`
	Stock       int     `gorm:"default:0"`
	MinStock    int     `gorm:"default:0"`
	Unit        string  `gorm:"type:varchar(50);default:'Unit'"`
	Price       float64 `gorm:"type:decimal(20,2);default:0"`

	// Relations
	Category Category `gorm:"foreignKey:CategoryID"`
	Supplier Supplier `gorm:"foreignKey:SupplierID"`
}

func (product *Product) ToProductResponse() web.ProductResponse {
	categoryName := ""
	if product.Category.Name != "" {
		categoryName = product.Category.Name
	}
	supplierName := ""
	if product.Supplier.Name != "" {
		supplierName = product.Supplier.Name
	}
	return web.ProductResponse{
		ID:           product.ID,
		Code:         product.Code,
		Name:         product.Name,
		CategoryID:   product.CategoryID,
		CategoryName: categoryName,
		SupplierID:   product.SupplierID,
		SupplierName: supplierName,
		Description:  product.Description,
		Stock:        product.Stock,
		MinStock:     product.MinStock,
		Unit:         product.Unit,
		Price:        product.Price,
	}
}

func (products Products) ToProductResponses() []web.ProductResponse {
	productResponses := []web.ProductResponse{}
	for _, product := range products {
		productResponses = append(productResponses, product.ToProductResponse())
	}
	return productResponses
}
