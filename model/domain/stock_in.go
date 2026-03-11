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
	User     User     `gorm:"foreignKey:CreatedByID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (stockIn *StockIn) ToStockInResponse() web.StockInResponse {
	productCode := ""
	productUnit := ""
	if stockIn.Product.Code != "" {
		productCode = stockIn.Product.Code
	}
	if stockIn.Product.Unit != "" {
		productUnit = stockIn.Product.Unit
	}
	return web.StockInResponse{
		ID:           stockIn.ID,
		Code:         stockIn.Code,
		Date:         stockIn.Date.Format("2006-01-02"),
		ProductID:    stockIn.ProductID,
		ProductName:  stockIn.Product.Name,
		ProductCode:  productCode,
		ProductUnit:  productUnit,
		SupplierID:   stockIn.SupplierID,
		SupplierName: stockIn.Supplier.Name,
		Quantity:     stockIn.Quantity,
		Notes:        stockIn.Notes,
		CreatedByID:  stockIn.CreatedByID,
		CreatedBy:    stockIn.User.FullName,
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
