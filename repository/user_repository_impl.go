package repository

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	domain "github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Login(db *gorm.DB, identity *string) domain.User {
	var user domain.User
	err := db.Where("email = ? OR number_phone = ?", identity, identity).First(&user).Error
	helper.PanicIfError(err)
	return user
}
