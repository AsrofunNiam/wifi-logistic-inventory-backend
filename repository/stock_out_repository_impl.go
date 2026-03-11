package repository

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"gorm.io/gorm"
)

type StockOutRepositoryImpl struct {
}

func NewStockOutRepository() StockOutRepository {
	return &StockOutRepositoryImpl{}
}

func (repository *StockOutRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.StockOuts {
	stockOuts := domain.StockOuts{}
	tx := db.Model(&domain.StockOut{})

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	err = tx.Preload("Product").Preload("User").Find(&stockOuts).Error
	helper.PanicIfError(err)

	return stockOuts
}

func (repository *StockOutRepositoryImpl) FindByID(db *gorm.DB, id *uint) domain.StockOut {
	var stockOut domain.StockOut

	err := db.Preload("Product").Preload("User").First(&stockOut, id).Error
	helper.PanicIfError(err)
	return stockOut
}

func (repository *StockOutRepositoryImpl) Create(db *gorm.DB, stockOut *domain.StockOut) (*domain.StockOut, error) {
	err := db.Create(&stockOut).Error
	if err != nil {
		return nil, err
	}

	// Preload relations after create
	err = db.Preload("Product").Preload("User").First(&stockOut, stockOut.ID).Error
	if err != nil {
		return nil, err
	}

	return stockOut, nil
}

func (repository *StockOutRepositoryImpl) Update(db *gorm.DB, stockOut *domain.StockOut) *domain.StockOut {
	err := db.Updates(&stockOut).Error
	helper.PanicIfError(err)

	err = db.Preload("Product").Preload("User").First(&stockOut, stockOut.ID).Error
	helper.PanicIfError(err)

	return stockOut
}

func (repository *StockOutRepositoryImpl) Delete(db *gorm.DB, id, deletedByID uint) {
	err := db.First(&domain.StockOut{}, id).Error
	helper.PanicIfError(err)

	// soft delete
	err = db.Updates(&domain.StockOut{
		Model:       gorm.Model{ID: uint(id)},
		DeletedByID: &deletedByID,
	}).Delete(&domain.StockOut{}, id).Error

	helper.PanicIfError(err)
}
