package repository

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"gorm.io/gorm"
)

type SupplierRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.Suppliers
	FindByID(db *gorm.DB, id *uint) domain.Supplier
	Create(db *gorm.DB, supplier *domain.Supplier) (*domain.Supplier, error)
	Update(db *gorm.DB, supplier *domain.Supplier) *domain.Supplier
	Delete(db *gorm.DB, id, deletedByID uint)
}
