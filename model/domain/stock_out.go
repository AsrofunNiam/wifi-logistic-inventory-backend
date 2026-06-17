package domain

import (
	"time"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"gorm.io/gorm"
)

type StockOuts []StockOut
type StockOut struct {
	gorm.Model
	CreatedByID uint  `gorm:"default:null"`
	UpdatedByID uint  `gorm:"default:null"`
	DeletedByID *uint `gorm:"default:null"`

	// Required Fields
	Code        string    `gorm:"type:varchar(50);uniqueIndex;not null"`
	Date        time.Time `gorm:"type:date;not null"`
	ProductID   uint      `gorm:"not null"`
	Destination string    `gorm:"type:varchar(255);not null"`
	Quantity    int       `gorm:"not null"`
	Notes       string    `gorm:"type:text"`

	// Relations
	Product Product `gorm:"foreignKey:ProductID"`
	User    User    `gorm:"foreignKey:ID;references:CreatedByID"`
}

func (stockOut *StockOut) ToStockOutResponse() web.StockOutResponse {
	productCode := ""
	productUnit := ""
	if stockOut.Product.Code != "" {
		productCode = stockOut.Product.Code
	}
	if stockOut.Product.Unit != "" {
		productUnit = stockOut.Product.Unit
	}
	return web.StockOutResponse{
		ID:          stockOut.ID,
		Code:        stockOut.Code,
		Date:        stockOut.Date.Format("2006-01-02"),
		ProductID:   stockOut.ProductID,
		ProductName: stockOut.Product.Name,
		ProductCode: productCode,
		ProductUnit: productUnit,
		Destination: stockOut.Destination,
		Quantity:    stockOut.Quantity,
		Notes:       stockOut.Notes,
		CreatedByID: stockOut.CreatedByID,
		CreatedBy:   stockOut.User.FullName,
		CreatedAt:   stockOut.CreatedAt,
	}
}

func (stockOuts StockOuts) ToStockOutResponses() []web.StockOutResponse {
	stockOutResponses := []web.StockOutResponse{}
	for _, stockOut := range stockOuts {
		stockOutResponses = append(stockOutResponses, stockOut.ToStockOutResponse())
	}
	return stockOutResponses
}
