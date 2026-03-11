package repository

import (
	domain "github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Login(db *gorm.DB, identity *string) domain.User
}
