package domain

import (
	"time"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"gorm.io/gorm"
)

type ProductPrices []ProductPrice
type ProductPrice struct {
	gorm.Model
	CreatedByID uint      `gorm:""`
	UpdatedByID uint      `gorm:""`
	DeletedByID uint      `gorm:""`
	ProductID   uint      `gorm:"not null"`
	CurrencyID  uint      `gorm:"not null"`
	Price       float64   `gorm:"type:decimal(20,2);not null"`
	StartDate   time.Time `gorm:"type:date;not null"`
	EndDate     time.Time `gorm:"type:date;not null"`

	// Relations
	Currency Currency `gorm:"foreignKey:CurrencyID;references:ID"`
}

func (productPrice *ProductPrice) ToProductPriceResponse() web.ProductPriceResponse {
	return web.ProductPriceResponse{
		// Required Fields
		ID: productPrice.ID,
		// Fields
		ProductID: productPrice.ProductID,
		Price:     productPrice.Price,
		StartDate: productPrice.StartDate.Format("2006-01-02"),
		EndDate:   productPrice.EndDate.Format("2006-01-02"),

		// Relations
		Currency: productPrice.Currency.ToCurrencyResponse(),
	}
}

func (productPrices ProductPrices) ToProductPriceResponses() []web.ProductPriceResponse {
	productPriceResponses := []web.ProductPriceResponse{}
	for _, productPrice := range productPrices {
		productPriceResponses = append(productPriceResponses, productPrice.ToProductPriceResponse())
	}
	return productPriceResponses
}
