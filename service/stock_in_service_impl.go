package service

import (
	"context"
	"encoding/json"
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
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type StockInServiceImpl struct {
	StockInRepository repository.StockInRepository
	ProductRepository repository.ProductRepository
	DB                *gorm.DB
	RedisClient       *redis.Client
	Validate          *validator.Validate
}

func NewStockInService(
	stockInRepository repository.StockInRepository,
	productRepository repository.ProductRepository,
	db *gorm.DB,
	redisClient *redis.Client,
	validate *validator.Validate,
) StockInService {
	return &StockInServiceImpl{
		StockInRepository: stockInRepository,
		ProductRepository: productRepository,
		DB:                db,
		RedisClient:       redisClient,
		Validate:          validate,
	}
}

func (service *StockInServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.StockInResponse {
	ctx := context.Background()
	key := "stock_ins:all"

	// Check cache in Redis
	data, err := service.RedisClient.Get(ctx, key).Result()
	if err == nil {
		var cachedStockIns []web.StockInResponse
		if err := json.Unmarshal([]byte(data), &cachedStockIns); err == nil {
			return cachedStockIns
		}
	}

	// If cache not found
	stockIns := service.StockInRepository.FindAll(service.DB, filters)
	stockInResponses := stockIns.ToStockInResponses()

	// Save to Redis
	jsonData, err := json.Marshal(stockInResponses)
	if err == nil {
		_ = service.RedisClient.Set(ctx, key, jsonData, 30*time.Minute).Err()
	}

	return stockInResponses
}

func (service *StockInServiceImpl) FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.StockInResponse {
	ctx := context.Background()
	key := fmt.Sprintf("stock_in:%d", *id)

	// Check cache in Redis
	data, err := service.RedisClient.Get(ctx, key).Result()
	if err == nil {
		var cachedStockIn web.StockInResponse
		if err := json.Unmarshal([]byte(data), &cachedStockIn); err == nil {
			return cachedStockIn
		}
	}

	// If cache not found, get from database
	stockIn := service.StockInRepository.FindByID(service.DB, id)
	stockInResponse := stockIn.ToStockInResponse()

	// Save to Redis
	jsonData, _ := json.Marshal(stockInResponse)
	_ = service.RedisClient.Set(ctx, key, jsonData, 30*time.Minute).Err()

	return stockInResponse
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

	// Invalidate cache
	ctx := context.Background()
	_ = service.RedisClient.Del(ctx, "stock_ins:all").Err()
	_ = service.RedisClient.Del(ctx, "products:all").Err()
	_ = service.RedisClient.Del(ctx, fmt.Sprintf("product:%d", request.ProductID)).Err()

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

	// Invalidate cache
	ctx := context.Background()
	_ = service.RedisClient.Del(ctx, "stock_ins:all").Err()
	_ = service.RedisClient.Del(ctx, fmt.Sprintf("stock_in:%d", id)).Err()
	_ = service.RedisClient.Del(ctx, "products:all").Err()
	_ = service.RedisClient.Del(ctx, fmt.Sprintf("product:%d", request.ProductID)).Err()

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

	// Invalidate cache
	ctx := context.Background()
	_ = service.RedisClient.Del(ctx, "stock_ins:all").Err()
	_ = service.RedisClient.Del(ctx, fmt.Sprintf("stock_in:%d", id)).Err()
	_ = service.RedisClient.Del(ctx, "products:all").Err()
	_ = service.RedisClient.Del(ctx, fmt.Sprintf("product:%d", existingStockIn.ProductID)).Err()
}
