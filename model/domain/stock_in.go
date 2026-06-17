package domain

import (
	"time"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"gorm.io/gorm"
)

type StockIns []StockIn
type StockIn struct {
	gorm.Model
	CreatedByID uint  `gorm:"default:null"`
	UpdatedByID uint  `gorm:"default:null"`
	DeletedByID *uint `gorm:"default:null"`

	// Required Fields
	Code       string    `gorm:"type:varchar(50);uniqueIndex;not null"`
	Date       time.Time `gorm:"type:date;not null"`
	ProductID  uint      `gorm:"not null"`
	SupplierID uint      `gorm:"not null"`
	Quantity   int       `gorm:"not null"`
	Notes      string    `gorm:"type:text"`

	// Relations
	Product  Product  `gorm:"foreignKey:ProductID"`
	Supplier Supplier `gorm:"foreignKey:SupplierID"`
	User     User     `gorm:"foreignKey:ID;references:CreatedByID"`
}

func (stockIn *StockIn) ToStockInResponse() web.StockInResponse {
	productName := ""
	productCode := ""
	productUnit := ""

	supplierName := ""
	// createdBy := ""

	if stockIn.ProductID == stockIn.Product.ID {
		productName = stockIn.Product.Name
		productCode = stockIn.Product.Code
		productUnit = stockIn.Product.Unit
	}

	if stockIn.SupplierID == stockIn.Supplier.ID {
		supplierName = stockIn.Supplier.Name
	}

	createdBy := ""
	if stockIn.CreatedByID == stockIn.User.ID {
		createdBy = stockIn.User.FullName
	}

	return web.StockInResponse{
		ID:           stockIn.ID,
		Code:         stockIn.Code,
		Date:         stockIn.Date.Format("2006-01-02"),
		ProductID:    stockIn.ProductID,
		ProductName:  productName,
		ProductCode:  productCode,
		ProductUnit:  productUnit,
		SupplierID:   stockIn.SupplierID,
		SupplierName: supplierName,
		Quantity:     stockIn.Quantity,
		Notes:        stockIn.Notes,
		CreatedByID:  stockIn.CreatedByID,
		CreatedBy:    createdBy,
		CreatedAt:    stockIn.CreatedAt,
	}
}

func (stockIns StockIns) ToStockInResponses() []web.StockInResponse {
	stockInResponses := []web.StockInResponse{}
	for _, stockIn := range stockIns {
		stockInResponses = append(stockInResponses, stockIn.ToStockInResponse())
	}
	return stockInResponses
}
