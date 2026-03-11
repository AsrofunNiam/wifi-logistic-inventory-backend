package repository

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.Categories
	FindByID(db *gorm.DB, id *uint) domain.Category
	Create(db *gorm.DB, category *domain.Category) (*domain.Category, error)
	Update(db *gorm.DB, category *domain.Category) *domain.Category
	Delete(db *gorm.DB, id, deletedByID uint)
}
