package service

import (
	"fmt"
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

type StockOutServiceImpl struct {
	StockOutRepository repository.StockOutRepository
	ProductRepository  repository.ProductRepository
	DB                 *gorm.DB
	Validate           *validator.Validate
}

func NewStockOutService(
	stockOutRepository repository.StockOutRepository,
	productRepository repository.ProductRepository,
	db *gorm.DB,
	validate *validator.Validate,
) StockOutService {
	return &StockOutServiceImpl{
		StockOutRepository: stockOutRepository,
		ProductRepository:  productRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *StockOutServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.StockOutResponse {
	stockOuts := service.StockOutRepository.FindAll(service.DB, filters)
	return stockOuts.ToStockOutResponses()
}

func (service *StockOutServiceImpl) FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.StockOutResponse {
	stockOut := service.StockOutRepository.FindByID(service.DB, id)
	return stockOut.ToStockOutResponse()
}

func (service *StockOutServiceImpl) Create(auth *auth.AccessDetails, request *web.StockOutCreateRequest, c *gin.Context) web.StockOutResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Start transaction
	tx := service.DB.Begin()

	// Check product stock
	product := service.ProductRepository.FindByID(tx, &request.ProductID)
	if product.Stock < request.Quantity {
		tx.Rollback()
		panic(exception.NewBadRequestError(fmt.Sprintf("Insufficient stock for product %s. Available stock: %d, requested: %d", product.Name, product.Stock, request.Quantity)))
	}

	// Parse date
	date, err := time.Parse("2006-01-02", request.Date)
	if err != nil {
		tx.Rollback()
		panic(exception.NewBadRequestError("Invalid date format"))
	}

	generatedCode, err := helper.GenerateTransactionCode(tx, &domain.StockOut{}, "SO", date)
	if err != nil {
		tx.Rollback()
		panic(exception.NewBadRequestError("Failed to generate stock out code"))
	}

	stockOut := domain.StockOut{
		Code:        generatedCode,
		Date:        date,
		ProductID:   request.ProductID,
		Destination: request.Destination,
		Quantity:    request.Quantity,
		Notes:       request.Notes,
		CreatedByID: auth.ID,
	}

	createdStockOut, err := service.StockOutRepository.Create(tx, &stockOut)
	if err != nil {
		tx.Rollback()
		panic(exception.NewBadRequestError(err.Error()))
	}

	// Update product stock (subtract)
	product.Stock -= request.Quantity

	if product.Stock < 0 {
		err := &exception.ErrorSendToResponse{Err: "Insufficient stock after update"}
		helper.PanicIfError(err)
	}

	service.ProductRepository.Update(tx, &product)

	tx.Commit()

	return createdStockOut.ToStockOutResponse()
}

func (service *StockOutServiceImpl) Update(auth *auth.AccessDetails, id uint, request *web.StockOutUpdateRequest, c *gin.Context) web.StockOutResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Start transaction
	tx := service.DB.Begin()

	// Get existing stock out
	existingStockOut := service.StockOutRepository.FindByID(tx, &id)

	// Calculate stock difference
	stockDiff := request.Quantity - existingStockOut.Quantity

	// Check if there's enough stock
	product := service.ProductRepository.FindByID(tx, &request.ProductID)
	if product.Stock < stockDiff {
		tx.Rollback()
		panic(exception.NewBadRequestError(fmt.Sprintf("Insufficient stock for product %s. Available stock: %d, additional requested: %d", product.Name, product.Stock, stockDiff)))
	}

	// Parse date
	date, err := time.Parse("2006-01-02", request.Date)
	if err != nil {
		tx.Rollback()
		panic(exception.NewBadRequestError("Invalid date format"))
	}

	stockOut := domain.StockOut{
		Model:       gorm.Model{ID: id},
		Code:        existingStockOut.Code,
		Date:        date,
		ProductID:   request.ProductID,
		Destination: request.Destination,
		Quantity:    request.Quantity,
		Notes:       request.Notes,
		UpdatedByID: auth.ID,
	}

	updatedStockOut := service.StockOutRepository.Update(tx, &stockOut)

	// Update product stock
	product.Stock -= stockDiff
	service.ProductRepository.Update(tx, &product)

	tx.Commit()

	return updatedStockOut.ToStockOutResponse()
}

func (service *StockOutServiceImpl) Delete(auth *auth.AccessDetails, id uint, c *gin.Context) {
	// Start transaction
	tx := service.DB.Begin()

	// Get existing stock out
	existingStockOut := service.StockOutRepository.FindByID(tx, &id)

	// Update product stock (add back)
	product := service.ProductRepository.FindByID(tx, &existingStockOut.ProductID)
	product.Stock += existingStockOut.Quantity
	service.ProductRepository.Update(tx, &product)

	service.StockOutRepository.Delete(tx, id, auth.ID)

	tx.Commit()
}
