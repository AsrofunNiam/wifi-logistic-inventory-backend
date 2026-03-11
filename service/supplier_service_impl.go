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

type SupplierServiceImpl struct {
	SupplierRepository repository.SupplierRepository
	DB                 *gorm.DB
	RedisClient        *redis.Client
	Validate           *validator.Validate
}

func NewSupplierService(
	supplierRepository repository.SupplierRepository,
	db *gorm.DB,
	redisClient *redis.Client,
	validate *validator.Validate,
) SupplierService {
	return &SupplierServiceImpl{
		SupplierRepository: supplierRepository,
		DB:                 db,
		RedisClient:        redisClient,
		Validate:           validate,
	}
}

func (service *SupplierServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.SupplierResponse {
	ctx := context.Background()
	key := "suppliers:all"

	// Check cache in Redis
	data, err := service.RedisClient.Get(ctx, key).Result()
	if err == nil {
		var cachedSuppliers []web.SupplierResponse
		if err := json.Unmarshal([]byte(data), &cachedSuppliers); err == nil {
			return cachedSuppliers
		}
	}

	// If cache not found
	suppliers := service.SupplierRepository.FindAll(service.DB, filters)
	supplierResponses := suppliers.ToSupplierResponses()

	// Save suppliers to Redis
	jsonData, err := json.Marshal(supplierResponses)
	if err == nil {
		_ = service.RedisClient.Set(ctx, key, jsonData, 100*time.Minute).Err()
	}

	return supplierResponses
}

func (service *SupplierServiceImpl) FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.SupplierResponse {
	ctx := context.Background()
	key := fmt.Sprintf("supplier:%d", *id)

	// Check cache in Redis
	data, err := service.RedisClient.Get(ctx, key).Result()
	if err == nil {
		var cachedSupplier web.SupplierResponse
		if err := json.Unmarshal([]byte(data), &cachedSupplier); err == nil {
			return cachedSupplier
		}
	}

	// If cache not found, get from database
	supplier := service.SupplierRepository.FindByID(service.DB, id)
	supplierResponse := supplier.ToSupplierResponse()

	// Save to Redis
	jsonData, _ := json.Marshal(supplierResponse)
	_ = service.RedisClient.Set(ctx, key, jsonData, 100*time.Minute).Err()

	return supplierResponse
}

func (service *SupplierServiceImpl) Create(auth *auth.AccessDetails, request *web.SupplierCreateRequest, c *gin.Context) web.SupplierResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	supplier := domain.Supplier{
		Code:        request.Code,
		Name:        request.Name,
		Contact:     request.Contact,
		Phone:       request.Phone,
		Email:       request.Email,
		Address:     request.Address,
		Status:      request.Status,
		CreatedByID: auth.ID,
	}

	createdSupplier, err := service.SupplierRepository.Create(service.DB, &supplier)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	// Invalidate cache
	ctx := context.Background()
	_ = service.RedisClient.Del(ctx, "suppliers:all").Err()

	return createdSupplier.ToSupplierResponse()
}

func (service *SupplierServiceImpl) Update(auth *auth.AccessDetails, id uint, request *web.SupplierUpdateRequest, c *gin.Context) web.SupplierResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	supplier := domain.Supplier{
		Model:       gorm.Model{ID: id},
		Code:        request.Code,
		Name:        request.Name,
		Contact:     request.Contact,
		Phone:       request.Phone,
		Email:       request.Email,
		Address:     request.Address,
		Status:      request.Status,
		UpdatedByID: auth.ID,
	}

	updatedSupplier := service.SupplierRepository.Update(service.DB, &supplier)

	// Invalidate cache
	ctx := context.Background()
	_ = service.RedisClient.Del(ctx, "suppliers:all").Err()
	_ = service.RedisClient.Del(ctx, fmt.Sprintf("supplier:%d", id)).Err()

	return updatedSupplier.ToSupplierResponse()
}

func (service *SupplierServiceImpl) Delete(auth *auth.AccessDetails, id uint, c *gin.Context) {
	service.SupplierRepository.Delete(service.DB, id, auth.ID)

	// Invalidate cache
	ctx := context.Background()
	_ = service.RedisClient.Del(ctx, "suppliers:all").Err()
	_ = service.RedisClient.Del(ctx, fmt.Sprintf("supplier:%d", id)).Err()
}
