package domain

import "github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"

type Currencies []Currency
type Currency struct {
	CreatedByID uint   `gorm:""`
	UpdatedByID uint   `gorm:""`
	DeletedByID uint   `gorm:""`
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(15);not null"`
}

func (currency *Currency) ToCurrencyResponse() web.CurrencyResponse {
	return web.CurrencyResponse{
		// Required Fields
		ID:   currency.ID,
		Name: currency.Name,
	}
}

func (users Currencies) ToCurrencyResponses() []web.CurrencyResponse {
	currencyResponses := []web.CurrencyResponse{}
	for _, user := range users {
		currencyResponses = append(currencyResponses, user.ToCurrencyResponse())
	}
	return currencyResponses
}
