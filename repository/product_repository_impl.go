package repository

import (
	"time"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.Products {
	products := domain.Products{}
	currentDate := time.Now().Format("2006-01-02")
	tx := db.Model(&domain.Product{})

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	err = tx.Preload("ProductPrice", "start_date <= ? AND end_date >= ?", currentDate, currentDate).Preload("Company").Find(&products).Error
	helper.PanicIfError(err)

	return products
}

func (repository *ProductRepositoryImpl) FindByID(db *gorm.DB, id *uint) domain.Product {
	var product domain.Product
	currentDate := time.Now().Format("2006-01-02")

	err := db.Preload("ProductPrice", "start_date <= ? AND end_date >= ?", currentDate, currentDate).
		First(&product, id).Error
	helper.PanicIfError(err)
	return product
}

func (repository *ProductRepositoryImpl) Create(db *gorm.DB, product *domain.Product) (*domain.Product, error) {
	err := db.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (repository *ProductRepositoryImpl) Update(db *gorm.DB, product *domain.Product) *domain.Product {
	err := db.Updates(&product).First(&product).Error
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(db *gorm.DB, id, deletedByID uint) {
	err := db.First(&domain.Product{}, id).Error
	helper.PanicIfError(err)

	// soft delete
	err = db.Updates(&domain.Product{
		Model:       gorm.Model{ID: uint(id)},
		DeletedByID: deletedByID,
	}).Delete(&domain.Company{}, id).Error

	helper.PanicIfError(err)
}
