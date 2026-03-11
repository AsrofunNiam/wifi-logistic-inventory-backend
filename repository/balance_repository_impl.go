package repository

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"gorm.io/gorm"
)

type BalanceRepositoryImpl struct {
}

func NewBalanceRepository() BalanceRepository {
	return &BalanceRepositoryImpl{}
}

func (repository *BalanceRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.Balances {
	balances := domain.Balances{}
	tx := db.Model(&domain.Balance{})

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	err = tx.Find(&balances).Error
	helper.PanicIfError(err)

	return balances
}

func (repository *BalanceRepositoryImpl) FindByID(db *gorm.DB, userID *uint) domain.Balance {
	var balance domain.Balance
	err := db.Where("user_id = ?", userID).First(&balance).Error
	helper.PanicIfError(err)
	return balance
}

func (repository *BalanceRepositoryImpl) Update(db *gorm.DB, balance *domain.Balance) *domain.Balance {
	err := db.Updates(&balance).First(&balance).Error
	helper.PanicIfError(err)
	return balance
}
