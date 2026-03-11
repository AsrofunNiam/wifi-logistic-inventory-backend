package domain

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"gorm.io/gorm"
)

type Balances []Balance
type Balance struct {
	gorm.Model
	CreatedByID uint    `gorm:""`
	UpdatedByID uint    `gorm:""`
	DeletedByID uint    `gorm:""`
	UserID      uint    `gorm:"not null"` // Required Fields
	Value       float64 `gorm:"type:decimal(10,2);not null" `
	Available   bool    `gorm:"default:true"`

	// Relations
	User User `gorm:"foreignKey:UserID"`
}

func (balance *Balance) ToBalanceResponse() web.BalanceResponse {
	return web.BalanceResponse{
		// Required Fields
		ID:        balance.ID,
		UserID:    balance.UserID,
		User:      balance.User.ToUserResponse(),
		Value:     balance.Value,
		Available: balance.Available,

		// Relations
		// CompanyResponse: balance.Company.ToCompanyResponse(),
	}
}

func (balances Balances) ToBalanceResponses() []web.BalanceResponse {
	balanceResponses := []web.BalanceResponse{}
	for _, balance := range balances {
		balanceResponses = append(balanceResponses, balance.ToBalanceResponse())
	}
	return balanceResponses
}
