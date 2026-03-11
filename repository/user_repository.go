package repository

import (
	domain "github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Login(db *gorm.DB, identity *string) domain.User
	FindAll(db *gorm.DB, filters *map[string]string) domain.Users
	FindByID(db *gorm.DB, id *uint) domain.User
	Create(db *gorm.DB, user *domain.User) (*domain.User, error)
	Update(db *gorm.DB, user *domain.User) *domain.User
	Delete(db *gorm.DB, id, deletedByID uint)
}
