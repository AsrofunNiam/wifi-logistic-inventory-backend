package repository

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"gorm.io/gorm"
)

type StockInRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.StockIns
	FindByID(db *gorm.DB, id *uint) domain.StockIn
	Create(db *gorm.DB, stockIn *domain.StockIn) (*domain.StockIn, error)
	Update(db *gorm.DB, stockIn *domain.StockIn) *domain.StockIn
	Delete(db *gorm.DB, id, deletedByID uint)
}
