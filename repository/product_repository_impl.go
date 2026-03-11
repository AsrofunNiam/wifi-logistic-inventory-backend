package repository

import (
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
	tx := db.Model(&domain.Product{})

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	err = tx.Preload("Category").Preload("Supplier").Find(&products).Error
	helper.PanicIfError(err)

	return products
}

func (repository *ProductRepositoryImpl) FindByID(db *gorm.DB, id *uint) domain.Product {
	var product domain.Product

	err := db.Preload("Category").Preload("Supplier").First(&product, id).Error
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

func (repository *ProductRepositoryImpl) Delete(db *gorm.DB, id uint, deletedByID uint) {
	err := db.First(&domain.Product{}, id).Error
	helper.PanicIfError(err)

	deleteByIDPtr := &deletedByID
	// soft delete
	err = db.Updates(&domain.Product{
		Model:       gorm.Model{ID: uint(id)},
		DeletedByID: deleteByIDPtr,
	}).Delete(&domain.Product{}, id).Error

	helper.PanicIfError(err)
}
