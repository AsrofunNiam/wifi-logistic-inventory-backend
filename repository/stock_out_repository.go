package repository

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"gorm.io/gorm"
)

type StockOutRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.StockOuts
	FindByID(db *gorm.DB, id *uint) domain.StockOut
	Create(db *gorm.DB, stockOut *domain.StockOut) (*domain.StockOut, error)
	Update(db *gorm.DB, stockOut *domain.StockOut) *domain.StockOut
	Delete(db *gorm.DB, id, deletedByID uint)
}
