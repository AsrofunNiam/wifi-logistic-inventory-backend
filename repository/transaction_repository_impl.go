package repository

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (repository *TransactionRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.Transactions {
	transactions := domain.Transactions{}
	tx := db.Model(&domain.Transaction{})

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	err = tx.Preload("Product").Preload("User").Find(&transactions).Error
	helper.PanicIfError(err)

	return transactions
}

func (repository *TransactionRepositoryImpl) Create(db *gorm.DB, transaction *domain.Transaction) *domain.Transaction {
	err := db.Create(&transaction).Error
	helper.PanicIfError(err)
	return transaction
}
