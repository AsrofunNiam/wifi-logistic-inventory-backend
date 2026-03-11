package repository

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.Categories {
	categories := domain.Categories{}
	tx := db.Model(&domain.Category{})

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	err = tx.Find(&categories).Error
	helper.PanicIfError(err)

	return categories
}

func (repository *CategoryRepositoryImpl) FindByID(db *gorm.DB, id *uint) domain.Category {
	var category domain.Category

	err := db.First(&category, id).Error
	helper.PanicIfError(err)
	return category
}

func (repository *CategoryRepositoryImpl) Create(db *gorm.DB, category *domain.Category) (*domain.Category, error) {
	err := db.Create(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (repository *CategoryRepositoryImpl) Update(db *gorm.DB, category *domain.Category) *domain.Category {
	err := db.Updates(&category).First(&category).Error
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(db *gorm.DB, id, deletedByID uint) {
	err := db.First(&domain.Category{}, id).Error
	helper.PanicIfError(err)

	// soft delete
	err = db.Updates(&domain.Category{
		Model:       gorm.Model{ID: uint(id)},
		DeletedByID: &deletedByID,
	}).Delete(&domain.Category{}, id).Error

	helper.PanicIfError(err)
}
