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

func (repository *UserRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.Users {
	users := domain.Users{}
	tx := db.Model(&domain.User{})

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	err = tx.Find(&users).Error
	helper.PanicIfError(err)

	return users
}

func (repository *UserRepositoryImpl) FindByID(db *gorm.DB, id *uint) domain.User {
	var user domain.User

	err := db.First(&user, id).Error
	helper.PanicIfError(err)
	return user
}

func (repository *UserRepositoryImpl) Create(db *gorm.DB, user *domain.User) (*domain.User, error) {
	err := db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *UserRepositoryImpl) Update(db *gorm.DB, user *domain.User) *domain.User {
	err := db.Updates(&user).First(&user).Error
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(db *gorm.DB, id, deletedByID uint) {
	err := db.First(&domain.User{}, id).Error
	helper.PanicIfError(err)

	// soft delete
	err = db.Updates(&domain.User{
		Model:       gorm.Model{ID: uint(id)},
		DeletedByID: &deletedByID,
	}).Delete(&domain.User{}, id).Error

	helper.PanicIfError(err)
}
