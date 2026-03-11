package repository

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.Products
	FindByID(db *gorm.DB, id *uint) domain.Product
	Create(db *gorm.DB, Product *domain.Product) (*domain.Product, error)
	Update(db *gorm.DB, Product *domain.Product) *domain.Product
	Delete(db *gorm.DB, id, deletedByID uint)
}
