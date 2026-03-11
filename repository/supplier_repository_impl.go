package repository

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"gorm.io/gorm"
)

type SupplierRepositoryImpl struct {
}

func NewSupplierRepository() SupplierRepository {
	return &SupplierRepositoryImpl{}
}

func (repository *SupplierRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.Suppliers {
	suppliers := domain.Suppliers{}
	tx := db.Model(&domain.Supplier{})

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	err = tx.Find(&suppliers).Error
	helper.PanicIfError(err)

	return suppliers
}

func (repository *SupplierRepositoryImpl) FindByID(db *gorm.DB, id *uint) domain.Supplier {
	var supplier domain.Supplier

	err := db.First(&supplier, id).Error
	helper.PanicIfError(err)
	return supplier
}

func (repository *SupplierRepositoryImpl) Create(db *gorm.DB, supplier *domain.Supplier) (*domain.Supplier, error) {
	err := db.Create(&supplier).Error
	if err != nil {
		return nil, err
	}
	return supplier, nil
}

func (repository *SupplierRepositoryImpl) Update(db *gorm.DB, supplier *domain.Supplier) *domain.Supplier {
	err := db.Updates(&supplier).First(&supplier).Error
	helper.PanicIfError(err)

	return supplier
}

func (repository *SupplierRepositoryImpl) Delete(db *gorm.DB, id, deletedByID uint) {
	err := db.First(&domain.Supplier{}, id).Error
	helper.PanicIfError(err)

	// soft delete
	err = db.Updates(&domain.Supplier{
		Model:       gorm.Model{ID: uint(id)},
		DeletedByID: &deletedByID,
	}).Delete(&domain.Supplier{}, id).Error

	helper.PanicIfError(err)
}
