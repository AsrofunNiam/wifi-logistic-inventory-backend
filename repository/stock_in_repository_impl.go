package repository

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"gorm.io/gorm"
)

type StockInRepositoryImpl struct {
}

func NewStockInRepository() StockInRepository {
	return &StockInRepositoryImpl{}
}

func (repository *StockInRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.StockIns {
	stockIns := domain.StockIns{}
	tx := db.Preload("Product").Preload("Supplier").Preload("User")

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	err = tx.Find(&stockIns).Error
	helper.PanicIfError(err)

	return stockIns
}

func (repository *StockInRepositoryImpl) FindByID(db *gorm.DB, id *uint) domain.StockIn {
	var stockIn domain.StockIn

	err := db.Preload("Product").Preload("Supplier").Preload("User").First(&stockIn, id).Error
	helper.PanicIfError(err)
	return stockIn
}

func (repository *StockInRepositoryImpl) Create(db *gorm.DB, stockIn *domain.StockIn) (*domain.StockIn, error) {
	err := db.Create(&stockIn).Error
	if err != nil {
		return nil, err
	}

	// Preload relations after create
	err = db.Preload("Product").Preload("Supplier").Preload("User").First(&stockIn, stockIn.ID).Error
	if err != nil {
		return nil, err
	}

	return stockIn, nil
}

func (repository *StockInRepositoryImpl) Update(db *gorm.DB, stockIn *domain.StockIn) *domain.StockIn {
	err := db.Updates(&stockIn).Error
	helper.PanicIfError(err)

	err = db.Preload("Product").Preload("Supplier").Preload("User").First(&stockIn, stockIn.ID).Error
	helper.PanicIfError(err)

	return stockIn
}

func (repository *StockInRepositoryImpl) Delete(db *gorm.DB, id, deletedByID uint) {
	err := db.First(&domain.StockIn{}, id).Error
	helper.PanicIfError(err)

	// soft delete
	err = db.Updates(&domain.StockIn{
		Model:       gorm.Model{ID: uint(id)},
		DeletedByID: &deletedByID,
	}).Delete(&domain.StockIn{}, id).Error

	helper.PanicIfError(err)
}
