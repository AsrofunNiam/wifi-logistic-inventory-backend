package service

import (
	"time"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/exception"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type StockInServiceImpl struct {
	StockInRepository repository.StockInRepository
	ProductRepository repository.ProductRepository
	DB                *gorm.DB
	Validate          *validator.Validate
}

func NewStockInService(
	stockInRepository repository.StockInRepository,
	productRepository repository.ProductRepository,
	db *gorm.DB,
	validate *validator.Validate,
) StockInService {
	return &StockInServiceImpl{
		StockInRepository: stockInRepository,
		ProductRepository: productRepository,
		DB:                db,
		Validate:          validate,
	}
}

func (service *StockInServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.StockInResponse {
	stockIns := service.StockInRepository.FindAll(service.DB, filters)
	return stockIns.ToStockInResponses()
}

func (service *StockInServiceImpl) FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.StockInResponse {
	stockIn := service.StockInRepository.FindByID(service.DB, id)
	return stockIn.ToStockInResponse()
}

func (service *StockInServiceImpl) Create(auth *auth.AccessDetails, request *web.StockInCreateRequest, c *gin.Context) web.StockInResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Start transaction
	tx := service.DB.Begin()

	// Parse date
	date, err := time.Parse("2006-01-02", request.Date)
	if err != nil {
		tx.Rollback()
		panic(exception.NewBadRequestError("Invalid date format"))
	}

	stockIn := domain.StockIn{
		Code:        request.Code,
		Date:        date,
		ProductID:   request.ProductID,
		SupplierID:  request.SupplierID,
		Quantity:    request.Quantity,
		Notes:       request.Notes,
		CreatedByID: auth.ID,
	}

	createdStockIn, err := service.StockInRepository.Create(tx, &stockIn)
	if err != nil {
		tx.Rollback()
		panic(exception.NewBadRequestError(err.Error()))
	}

	// Update product stock
	product := service.ProductRepository.FindByID(tx, &request.ProductID)
	product.Stock += request.Quantity
	service.ProductRepository.Update(tx, &product)

	tx.Commit()

	return createdStockIn.ToStockInResponse()
}

func (service *StockInServiceImpl) Update(auth *auth.AccessDetails, id uint, request *web.StockInUpdateRequest, c *gin.Context) web.StockInResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Start transaction
	tx := service.DB.Begin()

	// Get existing stock in
	existingStockIn := service.StockInRepository.FindByID(tx, &id)

	// Parse date
	date, err := time.Parse("2006-01-02", request.Date)
	if err != nil {
		tx.Rollback()
		panic(exception.NewBadRequestError("Invalid date format"))
	}

	// Calculate stock difference
	stockDiff := request.Quantity - existingStockIn.Quantity

	stockIn := domain.StockIn{
		Model:       gorm.Model{ID: id},
		Code:        request.Code,
		Date:        date,
		ProductID:   request.ProductID,
		SupplierID:  request.SupplierID,
		Quantity:    request.Quantity,
		Notes:       request.Notes,
		UpdatedByID: auth.ID,
	}

	updatedStockIn := service.StockInRepository.Update(tx, &stockIn)

	// Update product stock
	product := service.ProductRepository.FindByID(tx, &request.ProductID)
	product.Stock += stockDiff
	service.ProductRepository.Update(tx, &product)

	tx.Commit()

	return updatedStockIn.ToStockInResponse()
}

func (service *StockInServiceImpl) Delete(auth *auth.AccessDetails, id uint, c *gin.Context) {
	// Start transaction
	tx := service.DB.Begin()

	// Get existing stock in
	existingStockIn := service.StockInRepository.FindByID(tx, &id)

	// Update product stock (subtract)
	product := service.ProductRepository.FindByID(tx, &existingStockIn.ProductID)
	product.Stock -= existingStockIn.Quantity
	service.ProductRepository.Update(tx, &product)

	service.StockInRepository.Delete(tx, id, auth.ID)

	tx.Commit()
}
