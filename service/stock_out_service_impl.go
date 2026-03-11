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

type StockOutServiceImpl struct {
	StockOutRepository repository.StockOutRepository
	ProductRepository  repository.ProductRepository
	DB                 *gorm.DB
	RedisClient        *redis.Client
	Validate           *validator.Validate
}

func NewStockOutService(
	stockOutRepository repository.StockOutRepository,
	productRepository repository.ProductRepository,
	db *gorm.DB,
	redisClient *redis.Client,
	validate *validator.Validate,
) StockOutService {
	return &StockOutServiceImpl{
		StockOutRepository: stockOutRepository,
		ProductRepository:  productRepository,
		DB:                 db,
		RedisClient:        redisClient,
		Validate:           validate,
	}
}

func (service *StockOutServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.StockOutResponse {
	ctx := context.Background()
	key := "stock_outs:all"

	// Check cache in Redis
	data, err := service.RedisClient.Get(ctx, key).Result()
	if err == nil {
		var cachedStockOuts []web.StockOutResponse
		if err := json.Unmarshal([]byte(data), &cachedStockOuts); err == nil {
			return cachedStockOuts
		}
	}

	// If cache not found
	stockOuts := service.StockOutRepository.FindAll(service.DB, filters)
	stockOutResponses := stockOuts.ToStockOutResponses()

	// Save to Redis
	jsonData, err := json.Marshal(stockOutResponses)
	if err == nil {
		_ = service.RedisClient.Set(ctx, key, jsonData, 30*time.Minute).Err()
	}

	return stockOutResponses
}

func (service *StockOutServiceImpl) FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.StockOutResponse {
	ctx := context.Background()
	key := fmt.Sprintf("stock_out:%d", *id)

	// Check cache in Redis
	data, err := service.RedisClient.Get(ctx, key).Result()
	if err == nil {
		var cachedStockOut web.StockOutResponse
		if err := json.Unmarshal([]byte(data), &cachedStockOut); err == nil {
			return cachedStockOut
		}
	}

	// If cache not found, get from database
	stockOut := service.StockOutRepository.FindByID(service.DB, id)
	stockOutResponse := stockOut.ToStockOutResponse()

	// Save to Redis
	jsonData, _ := json.Marshal(stockOutResponse)
	_ = service.RedisClient.Set(ctx, key, jsonData, 30*time.Minute).Err()

	return stockOutResponse
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
		panic(exception.NewBadRequestError("Insufficient stock"))
	}

	// Parse date
	date, err := time.Parse("2006-01-02", request.Date)
	if err != nil {
		tx.Rollback()
		panic(exception.NewBadRequestError("Invalid date format"))
	}

	stockOut := domain.StockOut{
		Code:        request.Code,
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
	service.ProductRepository.Update(tx, &product)

	tx.Commit()

	// Invalidate cache
	ctx := context.Background()
	_ = service.RedisClient.Del(ctx, "stock_outs:all").Err()
	_ = service.RedisClient.Del(ctx, "products:all").Err()
	_ = service.RedisClient.Del(ctx, fmt.Sprintf("product:%d", request.ProductID)).Err()

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
		panic(exception.NewBadRequestError("Insufficient stock"))
	}

	// Parse date
	date, err := time.Parse("2006-01-02", request.Date)
	if err != nil {
		tx.Rollback()
		panic(exception.NewBadRequestError("Invalid date format"))
	}

	stockOut := domain.StockOut{
		Model:       gorm.Model{ID: id},
		Code:        request.Code,
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

	// Invalidate cache
	ctx := context.Background()
	_ = service.RedisClient.Del(ctx, "stock_outs:all").Err()
	_ = service.RedisClient.Del(ctx, fmt.Sprintf("stock_out:%d", id)).Err()
	_ = service.RedisClient.Del(ctx, "products:all").Err()
	_ = service.RedisClient.Del(ctx, fmt.Sprintf("product:%d", request.ProductID)).Err()

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

	// Invalidate cache
	ctx := context.Background()
	_ = service.RedisClient.Del(ctx, "stock_outs:all").Err()
	_ = service.RedisClient.Del(ctx, fmt.Sprintf("stock_out:%d", id)).Err()
	_ = service.RedisClient.Del(ctx, "products:all").Err()
	_ = service.RedisClient.Del(ctx, fmt.Sprintf("product:%d", existingStockOut.ProductID)).Err()
}
